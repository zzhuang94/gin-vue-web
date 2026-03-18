<template>
  <a-modal :width="width" cancel-text="取消" ok-text="保存" :maskClosable="false"
    @ok="submit" v-model:open="open" :confirm-loading="submitting">
    <template #title>
      <div v-html="title"></div>
      <div v-if="subtitle" v-html="subtitle" class="subtitle"></div>
    </template>
    <hr />
    <a-form :model="form" :label-col="{span: 4}" :wrapper-col="{span: 19}">
      <template v-for="r in rules" :key="r.key">
        <a-form-item v-if="! (pick && r.readonly)" :required="r.required && ! pick">
          <template #label>
            <a-checkbox v-if="pick" @click="toggleKey(r.key)"></a-checkbox>
            <Tooltip v-if="r.describe" :placement="pick ? 'top' : 'left'"
              color="blue" :icon="false" :msg="r.describe">
              <i class="fa fa-info-circle text-info"></i>&nbsp;{{ r.name }}
            </Tooltip>
            <span v-else style="cursor: pointer;">{{ r.name }}</span>
          </template>

          <a-select
            v-if="r.limit"
            v-model:value="form[r.key]"
            show-search allow-clear
            :filterOption="lib.filterByLabel"
            :placeholder="`请选择${r.name}`"
            :disabled="isDisabled(r)"
            :mode="r.split_sep ? 'multiple' : 'undefined'"
            >
            <a-select-option v-for="l in r.limit" :key="l.key" :label="l.label">
              {{ l.label }}
            </a-select-option>
          </a-select>

          <AjaxSelect
            v-else-if="r.trans && r.trans.ajax"
            v-model:value="form[r.key]"
            :placeholder="`请选择${r.name}`"
            :disabled="isDisabled(r)"
            :translate="r.trans"
            />

          <a-textarea
            v-else-if="r.textarea"
            v-model:value="form[r.key]"
            :placeholder="`请输入${r.name}，支持换行`"
            :disabled="isDisabled(r)"
            :auto-size="{minRows: 2}"
            />

          <a-date-picker
            v-else-if="r.date"
            v-model:value="form[r.key]"
            :placeholder="`请选择${r.name}`"
            :disabled="isDisabled(r)"
            style="width: 100%;"
            />

          <a-date-picker
            v-else-if="r.datetime"
            show-time
            v-model:value="form[r.key]"
            :placeholder="`请选择${r.name}`"
            :disabled="isDisabled(r)"
            style="width: 100%;"
            />

          <a-input
            v-else
            v-model:value="form[r.key]"
            :placeholder="`请输入${r.name}`"
            :disabled="isDisabled(r)"
            />

        </a-form-item>
      </template>
    </a-form>
    <hr />
  </a-modal>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { isArray, isEmpty } from 'lodash'
import dayjs from 'dayjs'

import type { Rule, Data } from '@libs/frm'
import lib from '@libs/lib'
import swal from '@libs/swal'
import strlib from '@libs/strlib'
import AjaxSelect from '@components/ajax-select.vue'
import Tooltip from '@components/tooltip.vue'

interface Props {
  title: string
  data?: Data,
  rules: Rule[]
  action: string
  width?: string
  subtitle?: string
  pick?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  data: () => ({id: ''}) as Data,
  width: '50%',
  subtitle: '',
  pick: false
})
const emit = defineEmits<{
  'submit': []
  'reload': []
}>()

const open = ref(true)
const submitting = ref(false)
const form = reactive(initForm())
const keys = reactive(initKeys())

const isDisabled = (r: Rule): boolean => {
  if (props.pick) {
    return ! keys[r.key]
  }
  return !!props.data.id && r.readonly
}

const toggleKey = (key: string) => {
  keys[key] = ! keys[key]
  if (! keys[key]) {
    delete form[key]
  }
}

const submit = async () => {
  const data = buildSubmitData()
  if (isEmpty(data)) {
    swal.error('请至少选择一条要修改的属性')
    return
  }
  submitting.value = true
  const ok = await lib.ajax(props.action, data)
  if (ok) {
    open.value = false
    emit('submit')
  }
  submitting.value = false
}

function initKeys(): Record<string, boolean> {
  const ans: Record<string, boolean> = {}
  for (const r of props.rules) {
    ans[r.key] = ! props.pick
  }
  return ans
}

function initForm(): Record<string, any> {
  const ans: Record<string, any> = {}
  for (const r of props.rules) {
    if (props.data[r.key] !== undefined) {
      if (r.split_sep && r.textarea) {
        ans[r.key] = String(props.data[r.key]).split(r.split_sep).join('\n')
      } else if (r.textarea && r.json) {
        ans[r.key] = strlib.formatJson(props.data[r.key])
      } else if (r.split_sep && r.limit) {
        if (props.data[r.key] === '') {
          ans[r.key] = []
        } else {
          ans[r.key] = String(props.data[r.key]).split(r.split_sep)
        }
      } else if (r.date && props.data[r.key]) {
        const d = dayjs(props.data[r.key])
        ans[r.key] = d.isValid() ? d : props.data[r.key]
      } else if (r.datetime && props.data[r.key]) {
        const d = dayjs(props.data[r.key])
        ans[r.key] = d.isValid() ? d : props.data[r.key]
      } else {
        ans[r.key] = props.data[r.key]
      }
    } else {
      if (r.split_sep && r.limit) {
        ans[r.key] = []
      } else {
        ans[r.key] = r.default || ''
      }
    }
  }
  return ans
}

function buildSubmitData(): Record<string, any> {
  const payload: Record<string, any> = {}
  for (const r of props.rules) {
    if (! keys[r.key] || form[r.key] === undefined) {
      continue
    }
    const val = form[r.key]
    if (isArray(val)) {
      if (r.split_sep) {
        payload[r.key] = val.map((x) => String(x)).join(r.split_sep)
      } else {
        payload[r.key] = val.map((x) => String(x))
      }
    } else if (r.date && val) {
      payload[r.key] = dayjs(val).isValid() ? dayjs(val).format('YYYY-MM-DD') : String(val)
    } else if (r.datetime && val) {
      payload[r.key] = dayjs(val).isValid() ? dayjs(val).format('YYYY-MM-DD HH:mm:ss') : String(val)
    } else {
      payload[r.key] = String(val)
    }
  }
  return payload
}
</script>

<style scoped>
.subtitle {
  font-size: 1.1rem;
  margin-top: 0.5rem;
  color: navy;
}
</style>