<template>
  <a-modal :width="width" cancel-text="取消" ok-text="保存" :maskClosable="false" @ok="submit" v-model:open="open" :confirm-loading="submitting">
    <template #title>
      <div v-html="title"></div>
      <div v-if="subtitle" v-html="subtitle" style="font-size: 0.9rem; color: navy;"></div>
    </template>

    <hr />

    <a-form :model="formData" :label-col="{span: 4}" :wrapper-col="{span: 19}">

      <a-form-item v-if="data.id" style="display: none">
        <a-input v-model:value="formData.id" type="hidden" />
      </a-form-item>

      <template v-for="v in rules" :key="v.key">
        <a-form-item v-if="! check || ! v.readonly" :required="v.required && ! check">
          <template #label>
            <a-checkbox v-if="check" v-model:checked="enabledKeys[v.key]"></a-checkbox>&nbsp;&nbsp;
            <Tooltip v-if="v.describe" placement="left" color="blue" :icon="false" :msg="v.describe">
              <i class="fa fa-info-circle text-info"></i>&nbsp;{{ v.name }}
            </Tooltip>
            <span v-else>{{ v.name }}</span>
          </template>

          <a-select
            v-if="v.limit"
            v-model:value="formData[v.key]"
            show-search allow-clear
            :filterOption="lib.filterByLabel"
            :placeholder="`请选择${v.name}`"
            :disabled="! enabledKeys[v.key] || (data.id || check) && v.readonly"
            :mode="v.split_sep ? 'multiple' : 'undefined'"
            >
            <a-select-option v-for="lv in v.limit" :key="lv.key || lv" :value="(lv.key !== undefined ? lv.key : lv)" :label="lv.label || lv">
              {{ lv.label || lv }}
            </a-select-option>
          </a-select>

          <AjaxSelect
            v-else-if="v.trans && v.trans.ajax"
            v-model:value="formData[v.key]"
            :placeholder="`请选择${v.name}`"
            :disabled="! enabledKeys[v.key] || (data.id || check) && v.readonly"
            :translate="v.trans"
            />

          <a-textarea
            v-else-if="v.textarea"
            v-show="! check || ! v.readonly"
            v-model:value="formData[v.key]"
            :placeholder="`请输入${v.name}，支持换行`"
            :readonly="(data.id || check) && v.readonly"
            :disabled="! enabledKeys[v.key] || (data.id || check) && v.readonly"
            :auto-size="{minRows: 2}"
            />

          <a-input
            v-else
            v-show="! check || ! v.readonly"
            v-model:value="formData[v.key]"
            :placeholder="`请输入${v.name}`"
            :readonly="(data.id || check) && v.readonly"
            :disabled="! enabledKeys[v.key] || (data.id || check) && v.readonly"
            />

        </a-form-item>
      </template>

    </a-form>

    <hr />
  </a-modal>
</template>

<script setup>
import { ref, reactive, toRaw, toRefs, watch } from 'vue'
import { isArray } from 'lodash'
import lib from '@libs/lib.ts'
import strlib from '@libs/strlib.ts'
import AjaxSelect from '@components/ajax-select.vue'
import Tooltip from '@components/tooltip.vue'

const props = defineProps({
  width: { type: String, default: '50%' },
  title: String,
  subtitle: { type: String, default: '' },
  data: Object,
  rules: { type: Array, default: () => [] },
  check: { type: Boolean, default: false },
  action: String,
})
const { data, rules } = toRefs(props)
const emit = defineEmits(['submit', 'reload'])
const open = ref(true)
const submitting = ref(false)

const enabledKeys = reactive(initEnableKeys())

watch(
  enabledKeys,
  (newVal) => {
    Object.keys(newVal).forEach((key) => {
      if (! newVal[key]) {
        delete formData[key]
      }
    })
  },
  { deep: true }
)

const formData = reactive(initFormData())

const submit = async () => {
  submitting.value = true
  const ok = await lib.ajax(props.action, buildSubmitData())
  if (ok) {
      open.value = false
      emit('submit')
  }
  submitting.value = false
}

function initEnableKeys() {
  const ans = {}
  for (const item of rules.value) {
    ans[item.key] = ! props.check
  }
  return ans
}

function initFormData() {
  const ans = {}
  for (const r of rules.value) {
    if (data.value[r.key] !== undefined) {
      if (r.split_sep && r.textarea) {
        ans[r.key] = data.value[r.key].split(r.split_sep).join('\n')
      } else if (r.textarea && r.json) {
        ans[r.key] = strlib.formatJson(data.value[r.key])
      } else if (r.split_sep && r.limit) {
        ans[r.key] = data.value[r.key].split(r.split_sep)
      } else {
        ans[r.key] = data.value[r.key]
      }
    } else {
      if (r.split_sep && (r.limit || r.limit_list)) {
        ans[r.key] = []
      } else {
        ans[r.key] = r.default || ''
      }
    }
  }
  return ans
}

function buildSubmitData() {
  const payload = {}
  if (formData.id !== undefined) payload.id = String(formData.id)
  for (const r of rules.value) {
    const val = formData[r.key]
    if (val === undefined) continue
    if (isArray(val)) {
      if (r.split_sep) {
        payload[r.key] = val.map((x) => String(x)).join(r.split_sep)
      } else {
        payload[r.key] = val.map((x) => String(x))
      }
    } else {
      payload[r.key] = String(val)
    }
  }
  return payload
}
</script>
