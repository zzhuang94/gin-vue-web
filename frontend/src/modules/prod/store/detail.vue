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
      <component :is="mc" v-bind="mp" @submit="reload" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { useModal } from '@/libs/modal'

interface Store {
  id: string
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

const props = defineProps<Props>()

const { mc, mp, loadModal } = useModal()

const reload = () => {
  window.location.reload()
}

const reject = () => {
  loadModal('reject?id=' + props.store.id)
}

const plus = () => {
  loadModal('plus?id=' + props.store.id)
}

const minus = () => {
  loadModal('minus?id=' + props.store.id)
}

const edit = () => {
  loadModal('edit?id=' + props.store.id)
}

const history = () => {
  loadModal('history?id=' + props.store.id)
}

</script>

<style scoped>
tr td:last-child {
  font-size: 1.2rem;
  font-weight: bold;
}
</style>