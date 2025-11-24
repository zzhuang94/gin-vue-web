<template>
  <div class="portlet">

    <Logs v-model:open="logsOpen" :lids :log_rules />

    <div class="portlet-head">
      <button @click="tryRollback" v-show="! rollbacking" :disabled="loading" class="btn btn-primary">
        <i class="fa fa-undo"></i> 回滚操作
      </button>
      <a-space style="margin-bottom: 10px">
        <button @click="rollback(confirmIds)" v-show="rollbacking" class="btn btn-warning"><i class="fa fa-undo"></i> 确认回滚</button>
        <button @click="cancel" v-show="rollbacking" class="btn btn-success"><i class="fa fa-arrow-left"></i> 取消操作</button>
      </a-space>
      <a-alert v-show="rollbacking" type="warning" message="由于操作之间的存在依赖关系，您必须同时回滚下面的全部操作"></a-alert>
    </div>
    <div class="portlet-body">
      <Searcher v-show="! rollbacking" v-model:arg="arg" :rules @search="fetchData" @clear="fetchData(true)" />

      <a-spin :spinning="loading">
        <table class="table table-hover">
          <thead>
            <tr>
              <th style="width: 10px"><input v-show="! rollbacking" type="checkbox" v-model="allSelected" @change="toggleAll" /></th>
              <th>ID</th>
              <th>用户</th>
              <th>路由</th>
              <th>详情</th>
              <th style="padding-left: 110px">改动示例</th>
              <th>操作时间</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="d, i in data" :key="d.id">
              <td><input v-show="! rollbacking" type="checkbox" v-model="ids" :value="d.id" /></td>
              <td>{{ d.id }}</td>
              <td>{{ d.user }}</td>
              <td>{{ d.path }}</td>
              <td>
                <label class="badge badge-info" style="cursor:pointer" @click="logs(d.lids)">
                  <i class="fa fa-search"></i> <b style="font-size: 1.02em">{{ d.lids.split(',').length }}</b>
                </label>
              </td>
              <td><Log :id="d.lids.split(',')[0]" /></td>
              <td>{{ d.created }}</td>
            </tr>
          </tbody>
        </table>
      </a-spin>

      <Pager v-show="! rollbacking" :loading :total="page.total" v-model:curr="page.curr" v-model:size="page.size" @page-change="fetchData" />

      <component :is="modalCurr" v-bind="modalProps" />
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref, shallowRef, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import axios from 'axios'

import lib from '@libs/lib.ts'
import swal from '@libs/swal.ts'

import Searcher from '@components/searcher.vue'
import Pager from '@components/pager.vue'
import Log from './log.vue'
import Logs from './logs.vue'

const props = defineProps(['rules', 'log_rules', 'arg', 'page_size'])
const route = useRoute()
const router = useRouter()
const modalCurr = shallowRef(null)
const modalProps = ref({})
const loading = ref(true)
const rollbacking = ref(false)
const data = ref([])
const arg = ref(props.arg)
const page = ref({ curr: 1, size: props.page_size })
const logsOpen = ref(false)
const lids = ref('')

const ids = ref([])
const confirmIds = ref([])
const allSelected = ref(false)
const toggleAll = () => {
  ids.value = allSelected.value ? data.value.map(r => r.id) : []
}
watch(ids, (newVal) => { allSelected.value = newVal.length > 0 && newVal.length === data.value.length })

const fetchData = async (page1) => {
  loading.value = true
  if (page1) {
    page.value.curr = 1
  }
  try {
    const params = {arg: arg.value, page: page.value}
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

const logs = (ids) => {
  console.log(ids, "IDS")
  logsOpen.value = true
  lids.value = ids
}

const tryRollback = async () => {
  if (ids.value.length == 0) {
    swal.warn("请至少勾选一条数据")
    return
  }
  loading.value = true
  const r = await lib.curlPost('confirm', { ids: ids.value })
  if (r.code) {
    rollback(ids.value)
  } else {
    rollbacking.value = true
    loading.value = false
    confirmIds.value = r.data.allids
    data.value = r.data.data
  }
}

const rollback = async (ids) => {
  loading.value = true
  const cf = await swal.confirm("确认操作？")
  if (! cf) {
    loading.value = false
    return
  }
  await lib.ajaxPost('rollback', { ids })
  cancel()
}

const cancel = () => {
  rollbacking.value = false
  confirmIds.value = []
  ids.value = []
  fetchData()
}

onMounted(fetchData)
</script>
