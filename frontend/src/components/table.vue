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
                <td v-else-if="d.option.length == 1" class="table-op-col">
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

<script setup>
import { ref, computed, watch, toRaw } from 'vue'
import { isEmpty } from 'lodash'
import oplib from '@libs/oplib.ts'
import Td from '@components/td.vue'

const emit = defineEmits(['update:sort-key', 'update:sort-order', 'sort-change', 'op-click'])
const props = defineProps({
  id: { type: String, default: '' },
  rules: Array,
  data: Array,
  option: { type: Array, default: [] },
  loading: Boolean,
  sortKey: String,
  sortOrder: String,
  noSort: { type: Boolean, default: false },
  small: { type: Boolean, default: false },
  batchSelect: {
    type: Boolean,
    default: false,
  },
  width: { type: String, default: '' },
  height: { type: String, default: '' },
  margin: { type: String, default: '0' },
})

const selectedRows = ref([])
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

const calcClass = (r) => {
  if (isEmpty(r.auto_hide)) {
    return ''
  }
  return 'auto-hide-' + r.auto_hide
}

const getSelectedIds = () => {
  return toRaw(selectedRows.value)
}

defineExpose({ getSelectedIds })

const sortKey = computed({get: () => props.sortKey, set: (value) => emit('update:sort-key', value)})
const sortOrder = computed({get: () => props.sortOrder, set: (value) => emit('update:sort-order', value)})

const sortIcon = (k) => {
  if (k == sortKey.value) {
    return sortOrder.value == 'ASC' ? 'fas fa-caret-up' : 'fas fa-caret-down'
  }
  return 'fas fa-unsorted fa-xs'
}

const sortChange = (k, v) => {
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

const processedData = () => {
  const ans = []
  for (const d of props.data) {
    d.option = oplib.filterOption(props.option, d)
    ans.push(d)
  }
  return ans
}

const runOp = (op, d) => {
  emit('op-click', oplib.calcOp(op, d))
}
</script>
