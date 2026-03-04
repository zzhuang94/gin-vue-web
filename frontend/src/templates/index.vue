<template>
  <div>
    <component :is="headerHintComponent" v-if="headerHintComponent" />
    <component :is="modalCurr" v-bind="modalProps" @submit="fetchData(true)" @reload="reloadModal()" />

    <div class="portlet">
      <div v-if="! isEmpty(top_menus) || dump" class="portlet-head">
        <a-space :size="6">
          <Button v-for="menu, i in top_menus" :key="i" :menu />

          <button v-if="dump" :disabled="dumping" class="btn btn-accent" @click="dumpExcel">
            <i class="fa fa-file-excel"></i> 导出Excel
          </button>
        </a-space>
      </div>

      <div class="portlet-body">
        <Searcher v-model:arg="arg" :rules="props.rules ?? []" @search="fetchData" @clear="fetchData(true)" />

        <Table ref="tableRef"
           :loading :rules :data :table_menus
           :batch-select="batch" :id="tableId"
           v-model:sort-key="sort.key" v-model:sort-order="sort.order"
           @sort-change="fetchData(true)"
           @menu-click="menuClick"
           />

        <Pager :loading :total="page.total" v-model:curr="page.curr" v-model:size="page.size" @page-change="fetchData" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, provide, ref, shallowRef, nextTick } from 'vue'
import type { Component } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { isEmpty } from 'lodash'
import type { Menu, TableMenu, Rule } from '@libs/frm.ts'

import lib from '@libs/lib.ts'
import swal from '@libs/swal.ts'
import excel from '@libs/excel.ts'

import Button from '@components/button.vue'
import Searcher from '@components/searcher.vue'
import Table from '@components/table.vue'
import Pager from '@components/pager.vue'

interface Props {
  headerHint?: string
  rules: Rule[]
  top_menus: Menu[]
  table_menus: TableMenu[]
  arg?: Record<string, any>
  page_size?: number
  sort?: Record<string, any>
  batch?: boolean
  dump?: boolean
}

const props = defineProps<Props>()

const route = useRoute()
const router = useRouter()

const loading = ref(true)
const data = ref<any[]>([])
const arg = ref<Record<string, any>>(props.arg ?? {})
const sort = ref<Record<string, any>>(props.sort ?? {})
const page = ref<{ curr: number; size: number; total?: number }>({ curr: 1, size: props.page_size ?? 10 })

const tableId = 'index-table-id'
const dumping = ref(false)
const tableRef = ref<InstanceType<typeof Table> | null>(null)

const headerHintComponent = computed(() => {
  if (props.headerHint) {
    return lib.loadComponent(props.headerHint)
  }
  return null
})

const modalCurr = shallowRef<Component | null>(null)
const modalProps = ref<Record<string, any>>({})
const modalUrl = ref('')
const reloadModal = () => {
  lib.reloadModal(modalUrl.value, modalProps)
}

const menuClick = async (m: Menu) => {
  if (m.type == 'modal') {
    modalUrl.value = m.url
    lib.loadModal(modalUrl.value, modalCurr, modalProps)
  } else if (m.type == 'link') {
    lib.redirect(m.url)
  } else if (m.type == 'async') {
    const ok = await lib.confirmCurl(m.url)
    if (ok) {
      fetchData()
    }
  } else if (m.type.startsWith('batch-')) {
    const ids = tableRef.value?.getSelectedIds() ?? []
    if (ids.length == 0) {
      swal.warn("注意！", "请至少选择一条数据");
      return
    }
    if (m.type == 'batch-edit') {
      const url = 'batch-edit?count=' + ids.length + '&ids=' + ids.join(',')
      lib.loadModal(url, modalCurr, modalProps)
    } else if (m.type == 'batch-modal') {
      const url = (m.url) + '?count=' + ids.length + '&ids=' + ids.join(',')
      lib.loadModal(url, modalCurr, modalProps)
    } else if (m.type == 'batch-delete') {
      const url = 'batch-delete?ids=' + ids.join(',')
      const ok = await lib.confirmCurl(url, '您将删除 ' + ids.length + ' 条数据')
      if (ok) {
        fetchData()
      }
    }
  }
}

const fetchData = async (page1?: boolean) => {
  loading.value = true
  if (page1) {
    page.value.curr = 1
  }
  try {
    const params = {arg: arg.value, sort: sort.value, page: page.value}
    const resp = await lib.curl('fetch', params)
    if (resp) {
      data.value = resp.data
      page.value = resp.page
    }
  } catch (error) {
    console.error(error)
  } finally {
    updateUrl()
    loading.value = false
  }
}

const updateUrl = () => {
  router.replace({
    path: route.path,
    query: lib.mapToUriParams(arg.value),
  })
}

const dumpExcel = async () => {
  dumping.value = true
  await nextTick()
  excel.exportTableToExcel(tableId)
  dumping.value = false
}

provide('menuClick', menuClick)

onMounted(fetchData)
</script>
