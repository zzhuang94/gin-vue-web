<template>
  <a-modal :width="width" :maskClosable="false" v-model:open="open" @cancel="emit('submit')">
    <template #title>
      <span v-html="title"></span>
    </template>

    <template #footer>
      <a-button v-if="action.sort" @click="saveSort" type="primary" :loading="sortLoading">保存排序</a-button>
    </template>

    <hr />

    <table class="table table-hover table-sm" width="100%">
      <thead>
        <tr>
          <th v-for="r in rules" :key="r.key">{{ r.name }}</th>
          <th class="table-op-col">操作</th>
        </tr>
      </thead>

      <draggable :list="rows" tag="tbody" item-key="id" handle=".drag-handle">
        <template #item="{ element }">
          <tr @mouseenter="activeRow = element.id">
            <td v-for="r in rules" :key="r.key" :style="r.width ? { width: r.width } : {}">
              <AjaxSelect
                v-if="r.trans && r.trans.ajax"
                v-model:value="element[r.key]"
                :placeholder="`请选择${r.name}`"
                :disabled="!action.edit || r.readonly"
                :translate="r.trans"
                :style="{ width: '100%' }"
              />

              <span
                v-else-if="activeRow !== element.id || !action.edit || r.readonly"
                v-html="lib.displayDK(r, element[r.key], true)"
              ></span>

              <BadgeSelect v-else-if="isBadgeLimit(r)" v-model:value="element[r.key]" :r />

              <a-select
                v-else-if="r.limit"
                :filterOption="lib.filterByLabel" v-model:value="element[r.key]"
                style="width: 100%" show-search
              >
                <a-select-option
                  v-for="option in r.limit"
                  :key="option.key || option"
                  :value="option.key !== undefined ? option.key : option"
                  :label="option.label || option"
                >
                  {{ option.label || option }}
                </a-select-option>
              </a-select>

              <a-textarea v-else-if="r.textarea" v-model:value="element[r.key]"
                :auto-size="{ minRows: 1 }">
              </a-textarea>

              <a-input v-else v-model:value="element[r.key]"></a-input>
            </td>

            <td class="table-op-col">
              <a-space :size="5">
                <button v-if="action.edit" class="btn btn-sm btn-outline-success"
                  @click="autoCurl('save?id=' + element.id, element)" title="保存">
                  <i class="fa fa-save"></i>
                </button>

                <button v-if="action.del" class="btn btn-sm btn-outline-danger"
                  @click="autoCurl('delete?id=' + element.id)" title="删除">
                  <i class="fa fa-trash"></i>
                </button>

                <button v-if="action.sort" class="btn btn-sm btn-outline-warning drag-handle"
                  title="拖拽排序">
                  <i class="fa fa-unsorted"></i>
                </button>
              </a-space>
            </td>
          </tr>
        </template>
      </draggable>

      <tr v-if="action.add" style="background-color: rgb(171 248 196)">
        <td v-for="r in rules" :key="r.key" :style="r.width ? { width: r.width } : {}">
          <BadgeSelect v-if="isBadgeLimit(r)" v-model:value="newRow[r.key]" :r />

          <a-select
            v-else-if="r.limit"
            v-model:value="newRow[r.key]"
            :filterOption="lib.filterByLabel"
            style="width: 100%"
            show-search>
            <a-select-option v-for="l in r.limit" :key="l.key" :label="l.label">
              {{ l.label }}
            </a-select-option>
          </a-select>

          <AjaxSelect
            v-else-if="r.trans && r.trans.ajax"
            v-model:value="newRow[r.key]"
            :placeholder="`请选择${r.name}`"
            :translate="r.trans"
            :style="{ width: '100%' }"
          />

          <a-textarea
            v-else-if="r.textarea"
            v-model:value="newRow[r.key]"
            :auto-size="{ minRows: 1 }"
            :placeholder="`请输入${r.name}，支持多行`"
          />

          <a-input
            v-else
            v-model:value="newRow[r.key]"
            :placeholder="`请输入${r.name}`"
          />
        </td>

        <td class="table-op-col">
          <button class="btn btn-ty btn-outline-primary"
            @click="autoCurl('save', { ...args, ...newRow })" title="保存">
            <i class="fa fa-plus"></i>
          </button>
        </td>
      </tr>
    </table>

    <hr v-if="action?.sort" />
  </a-modal>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import draggable from 'vuedraggable'

import { type Rule, type Arg, isBadgeLimit } from '@libs/frm'
import lib from '@libs/lib'

import AjaxSelect from '@components/ajax-select.vue'
import BadgeSelect from '@components/badge-select.vue'

interface Data {
  [key: string]: string
}

interface Action {
  prefix: string
  edit?: boolean
  del?: boolean
  sort?: boolean
  add?: boolean
}

interface Props {
  title: string
  width: string
  args: Arg
  data: Data[]
  rules: Rule[]
  action: Action
}

const props = defineProps<Props>()
const emit = defineEmits<{
  'submit': []
  'reload': []
}>()

const open = ref(true)
const rows = ref([...props.data])
const newRow = ref(initNewRow())
const sortLoading = ref(false)
const activeRow = ref('')

watch(
  () => props.data,
  (newData) => {
    rows.value = [...(newData ?? [])]
  },
  { deep: true }
)

function initNewRow(): Data {
  const ans: Data = {}
  for (const r of props.rules) {
    ans[r.key] = r.default || ''
  }
  return ans
}

const saveSort = async () => {
  sortLoading.value = true
  const ids = rows.value.map((d: Data) => d.id)
  await autoCurl('sort', { ids })
  sortLoading.value = false
}

const autoCurl = async (action: string, params?: Record<string, any>) => {
  const url = (props.action.prefix ?? '') + action
  const ok = await lib.ajax(url, params)

  if (ok) {
    emit('reload')
    newRow.value = initNewRow()
  }
}
</script>

<style scoped>
.table th, .table td {
  padding: 0.33rem;
}
.table-op-col {
  width: 10%;
}
</style>
