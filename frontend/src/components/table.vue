<template>
  <div>
    <a-spin :spinning="loading">
      <table class="table table-hover" :class="small ? 'table-sm' : ''" :id="id ? id : undefined" :style="{ width, height, margin }">
        <thead>
          <tr>
            <th v-if="batchSelect" style="width: 15px;">
              <input type="checkbox" v-model="allSelected" @change="toggleAll" />
            </th>
            <th v-for="r in rules" :key="r.key" :class="calcClass(r)" style="white-space: nowrap; cursor: pointer" @click="sortChange(r.key, r)">
              <span v-html="r.name"></span> <i v-if="! noSort && ! r.no_sort" :class="sortIcon(r.key)"></i>
            </th>
            <th v-if="option.length" class="table-op-col">操作</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="d in processedData()">
            <td v-if="batchSelect">
              <input type="checkbox" v-model="selectedRows" :value="d.id" />
            </td>
            <td v-for="r in rules" :key="r.key" :class="calcClass(r)" :style="r.width ? { 'word-break': 'break-all', 'max-width': r.width } : {}">
              <Td :r :v="d[r.key]" />
              </td>
              <template v-if="option.length">
                <td v-if="d.option.length > 1" class="table-op-col">
                  <a-dropdown placement="bottomRight" :trigger="['click']">
                    <button class="btn btn-default btn-sm">
                      <span class="fa fa-gear fa-sm" style="margin-right: 0.5rem"></span>
                      <span class="fa fa-angle-down"></span>
                    </button>
                    <template #overlay>
                      <a-menu>
                        <a-menu-item v-for="op, i in d.option" :key="i" @click="runOp(op, d)">
                          <i style="width: 1.5rem" :class="`fa fa-${op.icon}`"></i> {{ op.title }}
                        </a-menu-item>
                      </a-menu>
                    </template>
                  </a-dropdown>
                </td>
                <td v-else-if="d.option.length == 1 && d.option[0]" class="table-op-col">
                  <button class="btn btn-sm btn-info" @click="runOp(d.option[0], d)">
                    <i :class="`fa fa-${d.option[0].icon}`"></i>
                    {{ d.option[0].title }}
                  </button>
                </td>
                <td v-else class="table-op-col">
                  -
                </td>
              </template>
          </tr>
        </tbody>
      </table>
    </a-spin>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch, toRaw } from 'vue'
import { isEmpty } from 'lodash'
import oplib from '@libs/oplib.ts'
import Td from '@components/td.vue'

interface Rule {
  key: string
  name: string
  auto_hide?: string
  width?: string
  no_sort?: boolean
  [key: string]: any
}

interface Option {
  icon: string
  title: string
  [key: string]: any
}

interface DataRow {
  id: string | number
  option: Option[]
  [key: string]: any
}

interface Props {
  id?: string
  rules: Rule[]
  data: DataRow[]
  option?: Option[]
  loading?: boolean
  sortKey?: string
  sortOrder?: 'ASC' | 'DESC'
  noSort?: boolean
  small?: boolean
  batchSelect?: boolean
  width?: string
  height?: string
  margin?: string
}

const props = withDefaults(defineProps<Props>(), {
  id: '',
  option: () => [],
  loading: false,
  noSort: false,
  small: false,
  batchSelect: false,
  width: '',
  height: '',
  margin: '0'
})

const emit = defineEmits<{
  'update:sort-key': [value: string]
  'update:sort-order': [value: 'ASC' | 'DESC']
  'sort-change': []
  'op-click': [op: Option]
}>()

const selectedRows = ref<(string | number)[]>([])
const allSelected = ref(false)

const toggleAll = () => {
  if (allSelected.value) {
    selectedRows.value = props.data.map(row => row.id)
  } else {
    selectedRows.value = []
  }
}

watch(
  () => props.data,
  () => {
    selectedRows.value = []
    allSelected.value = false
  },
)
watch(selectedRows, (newVal) => {
  allSelected.value = newVal.length === props.data.length
})

const calcClass = (r: Rule): string => {
  if (isEmpty(r.auto_hide)) {
    return ''
  }
  return 'auto-hide-' + r.auto_hide
}

const getSelectedIds = () => {
  return toRaw(selectedRows.value)
}

defineExpose({ getSelectedIds })

const sortKey = computed({
  get: () => props.sortKey ?? '',
  set: (value: string) => emit('update:sort-key', value)
})
const sortOrder = computed({
  get: () => props.sortOrder ?? 'ASC',
  set: (value: 'ASC' | 'DESC') => emit('update:sort-order', value)
})

const sortIcon = (k: string): string => {
  if (k == sortKey.value) {
    return sortOrder.value == 'ASC' ? 'fas fa-caret-up' : 'fas fa-caret-down'
  }
  return 'fas fa-unsorted fa-xs'
}

const sortChange = (k: string, v: Rule) => {
  if (props.noSort || v.no_sort) {
    return
  }
  if (k == sortKey.value) {
    sortOrder.value = sortOrder.value == 'ASC' ? 'DESC' : 'ASC'
  } else {
    sortKey.value = k
  }
  emit('sort-change')
}

const processedData = (): DataRow[] => {
  const ans: DataRow[] = []
  for (const d of props.data) {
    d.option = oplib.filterOption(props.option, d)
    ans.push(d)
  }
  return ans
}

const runOp = (op: Option, d: DataRow) => {
  emit('op-click', oplib.calcOp(op, d))
}
</script>
