<template>
  <div class="portlet">
    <div class="portlet-body">
      <table class="table table-hover" style="width: 100%; border: 0;">
        <tbody>
          <tr>
            <td>类型</td>
            <td>{{ store.category }}</td>
          </tr>
          <tr>
            <td>材质</td>
            <td>{{ store.material }}</td>
          </tr>
          <tr>
            <td>颜色</td>
            <td>{{ store.color }}</td>
          </tr>
          <tr>
            <td>良品库存</td>
            <td>{{ store.goods }}</td>
          </tr>
          <tr>
            <td>劣品库存</td>
            <td>{{ store.bads }}</td>
          </tr>
          <tr>
            <td>备注</td>
            <td>{{ store.remark }}</td>
          </tr>
        </tbody>
      </table>
      <hr />
      <div style="text-align: center;">
       <a-space direction="vertical">
          <button class="btn btn-lg btn-primary" @click="plus">
            <i class="fa fa-download"></i> 良品入库
          </button>
          <button class="btn btn-lg btn-success" @click="minus">
            <i class="fa fa-upload"></i> 良品出库
          </button>
          <button class="btn btn-lg btn-danger" @click="reject">
            <i class="fa fa-warning"></i> 劣品上报
          </button>
          <button class="btn btn-lg btn-warning" @click="edit">
            <i class="fa fa-edit"></i> 库存编辑
          </button>
          <button class="btn btn-lg btn-info" @click="history">
            <i class="fa fa-history"></i> 变更历史
          </button>
        </a-space>
        <hr/>
        <a class="btn btn-lg btn-focus" href="/prod/store/index">
        <i class="fa fa-list"></i> 全部库存列表
      </a>
      </div>
      <component :is="modalCurr" v-bind="modalProps" @submit="reload" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, shallowRef } from 'vue';
import type { Component } from 'vue';
import lib from '@libs/lib.ts'

interface Store {
  id: string | number
  category: string
  material: string
  color: string
  goods: number
  bads: number
  remark: string
}

interface Props {
  store: Store
}

const modalCurr = shallowRef<Component | null>(null)
const modalProps = ref<Record<string, any>>({})

const props = defineProps<Props>()

const reload = () => {
  window.location.reload()
}

const reject = () => {
  lib.loadModal('reject?id=' + props.store.id, modalCurr, modalProps)
}

const plus = () => {
  lib.loadModal('plus?id=' + props.store.id, modalCurr, modalProps)
}

const minus = () => {
  lib.loadModal('minus?id=' + props.store.id, modalCurr, modalProps)
}

const edit = () => {
  lib.loadModal('edit?id=' + props.store.id, modalCurr, modalProps)
}

const history = () => {
  lib.loadModal('history?id=' + props.store.id, modalCurr, modalProps)
}

</script>

<style scoped>
tr td:last-child {
  font-size: 1.2rem;
  font-weight: bold;
}
</style>