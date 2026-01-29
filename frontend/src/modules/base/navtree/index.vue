<template>
  <div class="portlet">
    <div class="portlet-body">
      <Tree :data :ops @right-click="rightClick" @left-click="leftClick" />
      <component :is="modalCurr" v-bind="modalProps" @submit="loadData" @reload="loadData" />
    </div>
  </div>
</template>


<script setup lang="ts">
import { ref, shallowRef } from 'vue';
import type { Component } from 'vue';
import lib from '@libs/lib.ts'
import Tree from '@components/tree.vue'

interface Props {
  data?: Record<string, any>
}

const props = defineProps<Props>()
const data = ref<Record<string, any>>(props.data ?? {})
const ops = [
  {
    title: '编辑节点',
    icon: 'edit',
    op: 'edit',
  },
  {
    title: '删除节点',
    icon: 'trash',
    op: 'del',
  },
  {
    title: '新增子节点',
    icon: 'plus',
    op: 'add',
  },
  {
    title: '清空子节点',
    icon: 'remove',
    op: 'clear',
  },
]

const modalCurr = shallowRef<Component | null>(null)
const modalProps = ref<Record<string, any>>({})

const loadData = async () => {
  const ans = await lib.curl('fetch')
  if (ans) {
    data.value = ans.data
  }
}

const leftClick = async (id: string) => {
  lib.loadModal('edit?id=' + id, modalCurr, modalProps)
}

const rightClick = async (id: string, op: string) => {
  if (op == 'add') {
    lib.loadModal('add?id=' + id, modalCurr, modalProps)
  } else if (op == 'edit') {
    lib.loadModal('edit?id=' + id, modalCurr, modalProps)
  } else if (op == 'del') {
    const ok = await lib.confirmCurl('/base/navtree/delete?id=' + id, '若有子节点也将连带删除')
    if (ok) {
      loadData()
    }
  } else if (op == 'clear') {
    const ok = await lib.confirmCurl('/base/navtree/delete-sub?id=' + id, '仅清空子节点，不删除节点本身')
    if (ok) {
      loadData()
    }
  }
}
</script>
