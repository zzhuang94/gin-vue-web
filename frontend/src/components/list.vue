<template>
  <a-modal :width="width" :maskClosable="false" @ok="submit" v-model:open="open" @cancel="emit('submit')">
    <template #title>
      <span v-html="title"></span>
      <span v-if="subtitle" v-html="subtitle" style="font-size: 0.9rem; color: gray; margin-left: 10px"></span>
    </template>

    <template #footer>
      <a-button v-if="action.sort" type="primary" @click="saveSort" :loading="sortLoading">保存排序</a-button>
    </template>

    <hr />

    <table class="table table-hover table-sm" width="100%">
      <thead>
        <tr>
          <th v-for="rule in rules" :key="rule.key">
            {{ rule.name }}
          </th>
          <th class="table-op-col">操作</th>
        </tr>
      </thead>

      <draggable :list="localRows" tag="tbody" item-key="id" handle=".drag-handle">
        <template #item="{ element }">
          <tr @mouseenter="activeRow = element.id">
            <td v-for="rule in rules" :key="rule.key" :style="rule.width ? { width: rule.width } : {}">
              <AjaxSelect
                v-if="rule.trans && rule.trans.ajax"
                v-model:value="element[rule.key]"
                :placeholder="`请选择${rule.name}`"
                :disabled="!action.edit || rule.readonly"
                :translate="rule.trans"
                :style="{ width: '100%' }"
              />

              <span
                v-else-if="activeRow !== element.id || !action.edit || rule.is_readonly"
                v-html="lib.displayDK(rule, element[rule.key], true)"
              ></span>

              <BadgeSelect
                v-else-if="rule.limit && lib.isLimitBadge(rule.limit) && !rule.split_sep"
                :limit="rule.limit" :limit_map="rule.limit_map" v-model:value="element[rule.key]"
              />

              <a-select
                v-else-if="rule.limit"
                :filterOption="lib.filterByLabel" v-model:value="element[rule.key]"
                style="width: 100%" show-search
              >
                <a-select-option
                  v-for="option in rule.limit"
                  :key="option.key || option"
                  :value="option.key !== undefined ? option.key : option"
                  :label="option.label || option"
                >
                  {{ option.label || option }}
                </a-select-option>
              </a-select>

              <a-textarea v-else-if="rule.textarea" v-model:value="element[rule.key]" :auto-size="{ minRows: 1 }"></a-textarea>

              <a-input v-else v-model:value="element[rule.key]"></a-input>
            </td>

            <td class="table-op-col">
              <a-space :size="5">
                <button v-if="action.edit" class="btn btn-sm btn-outline-success" @click="autoCurl('save?id=' + element.id, element)" title="保存">
                  <i class="fa fa-save"></i>
                </button>

                <button v-if="action.del" class="btn btn-sm btn-outline-danger" @click="autoCurl('delete?id=' + element.id)" title="删除">
                  <i class="fa fa-trash"></i>
                </button>

                <button v-if="action.sort" class="btn btn-sm btn-outline-warning drag-handle" title="拖拽排序">
                  <i class="fa fa-unsorted"></i>
                </button>
              </a-space>
            </td>
          </tr>
        </template>
      </draggable>

      <tr v-if="action.add" style="background-color: rgb(171 248 196)">
        <td v-for="rule in rules" :key="rule.key" :style="rule.width ? { width: rule.width } : {}">
          <BadgeSelect
            v-if="rule.limit && lib.isLimitBadge(rule.limit) && !rule.split_sep"
            :limit="rule.limit" :limit_map="rule.limit_map" v-model:value="newRow[rule.key]"
          />

          <a-select v-else-if="rule.limit" v-model:value="newRow[rule.key]" :filterOption="lib.filterByLabel" style="width: 100%" show-search>
            <a-select-option v-for="option in rule.limit" :key="option.key" :value="option.key" :label="option.label || option">
              {{ option.label || option }}
            </a-select-option>
          </a-select>

          <AjaxSelect
            v-else-if="rule.trans && rule.trans.ajax" v-model:value="newRow[rule.key]"
            :placeholder="`请选择${rule.name}`" :translate="rule.trans" :style="{ width: '100%' }"
          />

          <a-textarea
            v-else-if="rule.textarea"
            v-model:value="newRow[rule.key]" :auto-size="{ minRows: 1 }" :placeholder="`请输入${rule.name}，支持多行`"
          ></a-textarea>

          <a-input v-else v-model:value="newRow[rule.key]" :placeholder="`请输入${rule.name}`"></a-input>
        </td>

        <td class="table-op-col">
          <button class="btn btn-ty btn-outline-primary" @click="autoCurl('save', { ...args, ...newRow })" title="保存">
            <i class="fa fa-plus"></i>
          </button>
        </td>
      </tr>
    </table>

    <hr v-if="action.sort" />
  </a-modal>
</template>

<script setup>
import { ref, reactive, watch } from 'vue'
import draggable from 'vuedraggable'
import lib from '@libs/lib.ts'
import AjaxSelect from '@components/ajax-select.vue'
import BadgeSelect from '@components/badge-select.vue'

const emit = defineEmits(['submit', 'reload'])
const props = defineProps(['width', 'title', 'subtitle', 'args', 'data', 'rules', 'action'])

const open = ref(true)
const localRows = ref([...props.data])
const sortLoading = ref(false)
const activeRow = ref(-1)

let newRow = reactive(initNewRow())

watch(
  () => props.data,
  (newData) => {
    localRows.value = [...newData]
  },
  { deep: true }
)

function initNewRow() {
  const result = {}
  for (const rule of props.rules) {
    result[rule.key] = rule.default || ''
  }
  return result
}

const saveSort = async () => {
  sortLoading.value = true
  const ids = localRows.value.map((item) => item.id)
  await autoCurl('sort', { ids })
  sortLoading.value = false
}

const autoCurl = async (action, params) => {
  const url = props.action.prefix + action
  const ok = await lib.ajax(url, params)

  if (ok) {
    emit('reload')
    newRow = reactive(initNewRow())
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
