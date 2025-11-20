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
        <Searcher v-model:arg="arg" :rules @search="fetchData" @clear="fetchData(true)" />

        <Table ref="tableRef"
           :loading :rules :data :option
           :batch-select="batch" :id="tableId"
           v-model:sort-key="sort.key" v-model:sort-order="sort.order"
           @sort-change="fetchData(true)"
           @op-click="toolClick"
           />

        <Pager :loading :total="page.total" v-model:curr="page.curr" v-model:size="page.size" @page-change="fetchData" />
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, onMounted, provide, ref, shallowRef, nextTick } from 'vue'
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

const props = defineProps(['headerHint', 'rules', 'lock', 'tool', 'option', 'arg', 'page_size', 'sort', 'batch', 'dump'])

const route = useRoute()
const router = useRouter()

const loading = ref(true)
const data = ref([])
const arg = ref(props.arg)
const sort = ref(props.sort)
const page = ref({ curr: 1, size: props.page_size })

const tableId = 'index-table-id'
const dumping = ref(false)
const tableRef = ref(null)

const headerHintComponent = computed(() => {
  if (props.headerHint) {
    return lib.loadComponent(props.headerHint)
  }
  return null
})

const modalCurr = shallowRef(null)
const modalProps = ref({})
const modalUrl = ref('')
const reloadModal = () => {
  lib.reloadModal(modalUrl.value, modalProps)
}

const toolClick = async (obj) => {
  if (obj.type == 'modal') {
    modalUrl.value = obj.url
    lib.loadModal(modalUrl.value, modalCurr, modalProps)
  } else if (obj.type == 'link') {
    lib.redirect(obj.url)
  } else if (obj.type == 'async') {
    const ok = await lib.confirmCurl(obj.url)
    if (ok) {
      fetchData()
    }
  } else if (obj.type.startsWith('batch-')) {
    const ids = tableRef.value.getSelectedIds()
    if (ids.length == 0) {
      swal.warn("注意！", "请至少选择一条数据");
      return
    }
    if (obj.type == 'batch-edit') {
      const url = 'batch-edit?count=' + ids.length + '&ids=' + ids.join(',')
      lib.loadModal(url, modalCurr, modalProps)
    } else if (obj.type == 'batch-modal') {
      const url = obj.url + '?count=' + ids.length + '&ids=' + ids.join(',')
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

const fetchData = async (page1) => {
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
