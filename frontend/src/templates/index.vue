<template>
  <div>
    <component v-if="header" :is="header" />

    <component :is="mc" v-bind="mp" @submit="reFetch" @reload="reloadModal" />

    <div class="portlet">
      <div v-if="topMenus.length > 0 || dump" class="portlet-head">
        <a-space :size="6">
          <Button v-for="menu, i in topMenus" :key="i" :menu />

          <button v-if="dump" :disabled="dumping" class="btn btn-accent" @click="dumpExcel">
            <i class="fa fa-file-excel"></i> 导出Excel
          </button>
        </a-space>
      </div>

      <div class="portlet-body">
        <Searcher v-model:arg="arg" :rules @search="fetch" @clear="reFetch" />

        <Table ref="tableRef"
           :loading :rules :data :tableMenus :batch :id="tableId"
           v-model:ids="ids" v-model:sort-key="sort.key" v-model:sort-order="sort.order"
           @sort-change="reFetch"
           @menu-click="menuClick"
           />

        <Pager :loading v-model:page="page" @update:page="fetch" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, provide, ref, nextTick, computed } from 'vue'
import type { Menu, TableMenu, Rule, Sort, Arg } from '@libs/frm'
import { useFetch } from '@/libs/fetch'
import { useModal } from '@/libs/modal'

import lib from '@libs/lib'
import swal from '@libs/swal'
import excel from '@libs/excel'

import Button from '@components/button.vue'
import Searcher from '@components/searcher.vue'
import Table from '@components/table.vue'
import Pager from '@components/pager.vue'

interface Props {
  header: string
  rules: Rule[]
  topMenus: Menu[]
  tableMenus: TableMenu[]
  arg: Arg
  pageSize: number
  sort: Sort
  batch: boolean
  dump: boolean
}

const props = defineProps<Props>()

const header = computed(() => {
  if (props.header) {
    return lib.loadComponent(props.header)
  }
  return null
})

const { loading, data, arg, page, sort, fetch, reFetch } = useFetch({
  arg: props.arg,
  pageSize: props.pageSize,
  sort: props.sort,
})

const { mc, mp, loadModal, reloadModal } = useModal()

const tableId = 'index-table-id'
const ids = ref<string[]>([])
const dumping = ref(false)

const menuClick = async (m: Menu) => {
  if (m.type == 'modal') {
    loadModal(m.url)
  } else if (m.type == 'link') {
    lib.redirect(m.url)
  } else if (m.type == 'async') {
    const ok = await lib.confirmCurl(m.url)
    if (ok) {
      fetch()
    }
  } else if (m.type.startsWith('batch-')) {
    if (ids.value.length == 0) {
      swal.warn("注意！", "请至少选择一条数据");
      return
    }
    if (m.type == 'batch-edit') {
      const url = 'batch-edit?count=' + ids.value.length + '&ids=' + ids.value.join(',')
      loadModal(url)
    } else if (m.type == 'batch-modal') {
      const url = (m.url) + '?count=' + ids.value.length + '&ids=' + ids.value.join(',')
      loadModal(url)
    } else if (m.type == 'batch-delete') {
      const url = 'batch-delete?ids=' + ids.value.join(',')
      const ok = await lib.confirmCurl(url, '您将删除 ' + ids.value.length + ' 条数据')
      if (ok) {
        fetch()
      }
    }
  }
}

const dumpExcel = async () => {
  dumping.value = true
  await nextTick()
  excel.exportTableToExcel(tableId)
  dumping.value = false
}

provide('menuClick', menuClick)

onMounted(fetch)
</script>
