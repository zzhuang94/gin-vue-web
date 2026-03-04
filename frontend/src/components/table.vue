<template>
  <div class="table-responsive">
    <a-spin :spinning="loading">
      <table class="table table-hover" :class="small ? 'table-sm' : ''" :id="id ? id : undefined" :style="{ width, height, margin }">
        <thead>
          <tr>
            <th v-if="batchSelect" style="width: 15px;">
              <input type="checkbox" v-model="allSelected" @change="toggleAll" />
            </th>
            <th v-for="r in rules" :key="r.key" :class="calcClass(r)" style="white-space: nowrap; cursor: pointer" @click="sortChange(r.key, r)">
              <a-tooltip v-if="r.describe" :title="r.describe">
                <i class="fa fa-info-circle fa-sm text-info" style="margin-right: 0.25rem"></i>
              </a-tooltip>
              <span v-html="r.name"></span> <i v-if="! noSort && ! r.no_sort" :class="sortIcon(r.key)"></i>
            </th>
            <th v-if="table_menus.length" class="table-op-col">操作</th>
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
            <td class="table-op-col">
              <span v-for="(tm, i) in aloneMenus(d)" :key="'alone-' + i" class="alone-menu">
                <button class="btn btn-sm" :class="`btn-${tm.color}`" @click="runMenu(tm, d)">
                  <i :class="`fa fa-${tm.icon}`"></i> {{ tm.title }}
                </button>
              </span>
              <a-dropdown v-if="dropdownMenus(d).length > 1" placement="bottomRight" :trigger="['click']">
                <button class="btn btn-default btn-sm table-op-more-btn">
                  <span class="fa fa-gear fa-sm" style="margin-right: 0.35rem"></span>
                  <span v-if="aloneMenus(d).length>0">更多操作</span>
                  <span class="fa fa-angle-down" style="margin-left: 0.25rem"></span>
                </button>
                <template #overlay>
                  <a-menu>
                    <a-menu-item v-for="(tm, i) in dropdownMenus(d)" :key="i" @click="runMenu(tm, d)">
                      <i style="width: 1.5rem" :class="`fa fa-${tm.icon}`"></i> {{ tm.title }}
                    </a-menu-item>
                  </a-menu>
                </template>
              </a-dropdown>
              <button v-if="onlyOneMenu(d)" class="btn btn-sm btn-info" @click="runMenu(onlyOneMenu(d)!, d)">
                <i :class="`fa fa-${onlyOneMenu(d)!.icon}`"></i>
                {{ onlyOneMenu(d)!.title }}
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </a-spin>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch, toRaw, onMounted, onUnmounted } from 'vue'
import { isEmpty } from 'lodash'
import oplib from '@libs/oplib.ts'
import Td from '@components/td.vue'
import type { Rule, TableMenu, Menu } from '@libs/frm.ts'

interface DataRow {
  id: string | number
  table_menus: TableMenu[]
  [key: string]: any
}

interface Props {
  id?: string
  rules: Rule[]
  data: DataRow[]
  table_menus?: TableMenu[]
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
  table_menus: () => [],
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
  'menu-click': [m: Menu]
}>()

const selectedRows = ref<(string | number)[]>([])
const allSelected = ref(false)

const NARROW_BREAKPOINT = 768
const isNarrowView = ref(false)
const checkNarrow = () => {
  isNarrowView.value = window.innerWidth < NARROW_BREAKPOINT
}
onMounted(() => {
  checkNarrow()
  window.addEventListener('resize', checkNarrow)
})
onUnmounted(() => {
  window.removeEventListener('resize', checkNarrow)
})

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
    d.table_menus = oplib.filterMenus(props.table_menus, d)
    ans.push(d)
  }
  return ans
}

const runMenu = (tm: TableMenu, d: DataRow) => {
  emit('menu-click', oplib.calcMenu(tm, d))
}

const aloneMenus = (d: DataRow): TableMenu[] => {
  if (isNarrowView.value) {
    return []
  }
  return d.table_menus.filter((tm) => tm.alone)
}

const dropdownMenus = (d: DataRow): TableMenu[] => {
  if (isNarrowView.value) {
    return d.table_menus
  }
  return d.table_menus.filter((tm) => !tm.alone)
}

const onlyOneMenu = (d: DataRow): TableMenu | undefined => {
  const tms = dropdownMenus(d)
  return tms.length === 1 ? tms[0] : undefined
}
</script>

<style scoped>
.alone-menu {
  margin-right: 0.35rem;
}
</style>
