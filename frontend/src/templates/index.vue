<template>
  <div>
    <component :is="headerHintComponent" v-if="headerHintComponent" />
    <component :is="modalCurr" v-bind="modalProps" @submit="fetchData(true)" @reload="reloadModal()" />

    <div class="portlet">
      <div v-if="! isEmpty(tool) || ! isEmpty(lock) || dump" class="portlet-head">
        <a-space :size="6">
          <Button v-for="btn, i in tool" :key="i" v-bind="btn" />

          <button v-if="dump" :disabled="dumping" class="btn btn-accent" @click="dumpExcel">
            <i class="fa fa-file-excel"></i> 导出Excel
          </button>

          <Lock v-for="lck, i in lock" :key="i" v-bind="lck" />
        </a-space>
      </div>

      <div class="portlet-body">
        <Searcher v-model:arg="arg" :rules="props.rules ?? []" @search="fetchData" @clear="fetchData(true)" />

        <Table ref="tableRef"
           :loading :rules="props.rules ?? []" :data :option
           :batch-select="batch" :id="tableId"
           v-model:sort-key="sort.key" v-model:sort-order="sort.order"
           @sort-change="fetchData(true)"
           @op-click="handleOpClick"
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

import lib from '@libs/lib.ts'
import swal from '@libs/swal.ts'
import excel from '@libs/excel.ts'

import Lock from '@components/lock.vue'
import Button from '@components/button.vue'
import Searcher from '@components/searcher.vue'
import Table from '@components/table.vue'
import Pager from '@components/pager.vue'

interface ToolObj {
  type: string
  url?: string
  [key: string]: any
}

interface Props {
  headerHint?: string
  rules?: any[]
  lock?: any[]
  tool?: ToolObj[]
  option?: any[]
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

const handleOpClick = async (op: any) => {
  const obj = op as ToolObj
  await toolClick(obj)
}

const toolClick = async (obj: ToolObj) => {
  if (obj.type == 'modal') {
    modalUrl.value = obj.url ?? ''
    lib.loadModal(modalUrl.value, modalCurr, modalProps)
  } else if (obj.type == 'link') {
    lib.redirect(obj.url ?? '')
  } else if (obj.type == 'async') {
    const ok = await lib.confirmCurl(obj.url ?? '')
    if (ok) {
      fetchData()
    }
  } else if (obj.type.startsWith('batch-')) {
    const ids = tableRef.value?.getSelectedIds() ?? []
    if (ids.length == 0) {
      swal.warn("注意！", "请至少选择一条数据");
      return
    }
    if (obj.type == 'batch-edit') {
      const url = 'batch-edit?count=' + ids.length + '&ids=' + ids.join(',')
      lib.loadModal(url, modalCurr, modalProps)
    } else if (obj.type == 'batch-modal') {
      const url = (obj.url ?? '') + '?count=' + ids.length + '&ids=' + ids.join(',')
      lib.loadModal(url, modalCurr, modalProps)
    } else if (obj.type == 'batch-delete') {
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

provide('toolClick', toolClick)

onMounted(fetchData)
</script>
