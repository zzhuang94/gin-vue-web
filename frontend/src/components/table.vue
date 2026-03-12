<template>
  <div class="table-responsive">
    <a-spin :spinning="loading">
      <table class="table table-hover" :class="classes" :id="id" :style="styles">
        <thead>
          <tr>
            <th v-if="batch" style="width: 15px;">
              <input type="checkbox" v-model="pickAll" @change="toggleAll" />
            </th>
            <th v-for="r in rules" :key="r.key" :class="calcClass(r)"
              style="white-space: nowrap; cursor: pointer" @click="sortChange(r)">
              <a-tooltip v-if="r.describe" :title="r.describe">
                <i class="fa fa-info-circle fa-sm text-info" style="margin-right: 0.25rem"></i>
              </a-tooltip>
              <span v-html="r.name" style="margin-right: 0.2rem;"></span>
              <i v-if="! noSort && ! r.no_sort" :class="sortIcon(r.key)"></i>
            </th>
            <th v-if="tableMenus.length" class="table-op-col">操作</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="d in data">
            <td v-if="batch">
              <input type="checkbox" v-model="ids" :value="d.id" />
            </td>
            <td v-for="r in rules" :key="r.key" :class="calcClass(r)" :style="tdStyle(r)">
              <Td :r :v="d[r.key]" />
            </td>
            <td class="table-op-col">
              <span v-for="tm, i in aloneMenus(d)" :key="'alone-' + i" class="alone-menu">
                <button class="btn btn-sm" :class="`btn-${tm.color}`" @click="runMenu(tm, d)">
                  <i :class="`fa fa-${tm.icon}`"></i> {{ tm.title }}
                </button>
              </span>
              <a-dropdown v-if="dropdownMenus(d).length > 1" :trigger="['click']">
                <button class="btn btn-default btn-sm table-op-more-btn">
                  <span class="fa fa-gear fa-sm" style="margin-right: 0.35rem"></span>
                  <span v-if="aloneMenus(d).length>0">更多操作</span>
                  <span class="fa fa-angle-down" style="margin-left: 0.25rem"></span>
                </button>
                <template #overlay>
                  <a-menu>
                    <a-menu-item v-for="tm, i in dropdownMenus(d)" :key="i" @click="runMenu(tm, d)">
                      <i style="width: 1.5rem" :class="`fa fa-${tm.icon}`"></i> {{ tm.title }}
                    </a-menu-item>
                  </a-menu>
                </template>
              </a-dropdown>
              <button v-if="onlyOneMenu(d)" class="btn btn-sm btn-info"
                @click="runMenu(onlyOneMenu(d)!, d)">
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
import { ref, computed, watch, onMounted, onUnmounted } from 'vue'
import { isEmpty } from 'lodash'
import oplib from '@libs/oplib.ts'
import Td from '@components/td.vue'
import type { Rule, TableMenu, Menu, Data } from '@libs/frm.ts'

interface DataRow extends Data {
  tableMenus: TableMenu[]
}

interface Props {
  id?: string
  rules: Rule[]
  data: Data[]
  tableMenus?: TableMenu[]
  loading?: boolean
  noSort?: boolean
  sortKey?: string
  sortOrder?: 'ASC' | 'DESC'
  small?: boolean
  batch?: boolean
  ids?: string[]
  width?: string
  height?: string
  margin?: string
}

const props = withDefaults(defineProps<Props>(), {
  id: '',
  tableMenus: () => [],
  loading: false,
  noSort: false,
  sortKey: '',
  sortOrder: 'ASC',
  small: false,
  batch: false,
  ids: () => [],
  width: '',
  height: '',
  margin: '0'
})

const emit = defineEmits<{
  'update:ids': [value: string[]]
  'update:sort-key': [value: string]
  'update:sort-order': [value: 'ASC' | 'DESC']
  'sort-change': []
  'menu-click': [m: Menu]
}>()

const classes = computed(() => {
  return {
    'table-sm': props.small,
  }
})

const styles = computed(() => {
  return {
    width: props.width,
    height: props.height,
    margin: props.margin,
  }
})

const ids = computed({
  get: () => props.ids,
  set: (value: string[]) => {
    emit('update:ids', value)
    pickAll.value = value.length === props.data.length
  }
})

const pickAll = ref(false)

const toggleAll = () => {
  if (pickAll.value) {
    ids.value = props.data.map(row => row.id)
  } else {
    ids.value = []
  }
}

const data = computed(() =>{
    const ans: DataRow[] = []
    for (const d of props.data) {
      ans.push({ ...d, tableMenus: oplib.filterMenus(props.tableMenus, d) })
    }
    return ans
})

watch(
  () => props.data,
  () => {
    ids.value = []
    pickAll.value = false
  },
)

const calcClass = (r: Rule): string => {
  if (isEmpty(r.auto_hide)) {
    return ''
  }
  return 'auto-hide-' + r.auto_hide
}

const tdStyle = (r: Rule): string => {
  if (! r.width) {
    return ''
  }
  return 'word-break: break-all; max-width: ' + r.width + ';'
}

const sortIcon = (k: string): string => {
  if (k == sortKey.value) {
    return sortOrder.value == 'ASC' ? 'fas fa-caret-up' : 'fas fa-caret-down'
  }
  return 'fas fa-unsorted fa-xs'
}

const sortKey = computed({
  get: () => props.sortKey ?? '',
  set: (value: string) => emit('update:sort-key', value)
})
const sortOrder = computed({
  get: () => props.sortOrder ?? 'ASC',
  set: (value: 'ASC' | 'DESC') => emit('update:sort-order', value)
})

const sortChange = (r: Rule) => {
  if (props.noSort || r.no_sort) {
    return
  }
  if (r.key == sortKey.value) {
    sortOrder.value = sortOrder.value == 'ASC' ? 'DESC' : 'ASC'
  } else {
    sortKey.value = r.key
  }
  emit('sort-change')
}

const runMenu = (tm: TableMenu, d: DataRow) => {
  emit('menu-click', oplib.calcMenu(tm, d))
}

const aloneMenus = (d: DataRow): TableMenu[] => {
  if (isNarrowView.value) {
    return []
  }
  return d.tableMenus.filter((tm) => tm.alone)
}

const dropdownMenus = (d: DataRow): TableMenu[] => {
  if (isNarrowView.value) {
    return d.tableMenus
  }
  return d.tableMenus.filter((tm) => !tm.alone)
}

const onlyOneMenu = (d: DataRow): TableMenu | undefined => {
  const tms = dropdownMenus(d)
  return tms.length === 1 ? tms[0] : undefined
}

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
</script>

<style scoped>
.alone-menu {
  margin-right: 0.35rem;
}
</style>
