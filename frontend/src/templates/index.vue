<template>
  <div>
    <component v-if="header" :is="header" />

    <component :is="mc" v-bind="mp" @submit="reFetch" @reload="reloadModal" />

    <div class="portlet">
      <div v-if="topMenus.length > 0 || dump" class="portlet-head">
        <Header :topMenus :dump @menu-click="menuClick" />
      </div>

      <div class="portlet-body">
        <Searcher v-model:arg="arg" :rules @search="fetch" @clear="reFetch" />

        <Table ref="tableRef"
           :loading :rules :data :tableMenus :batch :id="excel.DEFAULT_ID"
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
import { onMounted, computed } from 'vue'
import type { Menu, TableMenu, Rule, Sort, Arg } from '@libs/frm'
import { useFetch } from '@/libs/fetch'
import { useModal } from '@/libs/modal'
import { useMenu } from '@/libs/menu'
import { useBatch } from '@/libs/batch'

import lib from '@libs/lib'
import excel from '@libs/excel'
import Header from '@components/header.vue'
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
const { runMenu } = useMenu(loadModal, fetch)
const { ids, runBatch } = useBatch(loadModal, fetch)

const menuClick = async (m: Menu) => {
  if (m.type.startsWith('batch-')) {
    runBatch(m)
  } else {
    runMenu(m)
  }
}

onMounted(fetch)
</script>
