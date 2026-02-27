<template>
  <a-modal width="50%" :footer="null" v-model:open="open">
    <template #title>
      <i class="fa fa-history"></i> 变更历史
    </template>
    <hr />
    <table class="table table-hover" style="width: 100%; border: 0;">
      <thead>
        <tr>
          <th>操作人</th>
          <th>操作</th>
          <th>数量</th>
          <th>备注</th>
          <th>操作时间</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="op in ops" :key="op.id">
          <td>{{ op.user }}</td>
          <td>
            <span v-if="op.op === 'PLUS'" class="badge badge-primary">良品入库</span>
            <span v-if="op.op === 'MINUS'" class="badge badge-success">良品出库</span>
            <span v-if="op.op === 'REJECT'" class="badge badge-danger">劣品上报</span>
            <span v-if="op.op === 'EDIT'" class="badge badge-warning">库存编辑</span>
          </td>
          <td>{{ op.count }}</td>
          <td>{{ op.remark }}</td>
          <td>{{ strlib.formatTime(op.created) }}</td>
        </tr>
      </tbody>
    </table>
  </a-modal>
</template>


<script setup lang="ts">
import { ref } from 'vue';
import strlib from '@libs/strlib.ts'

interface StoreOp {
  id: string | number
  op: string
  count: number
  remark: string
  user: string
  created: string
}

interface Props {
  ops?: StoreOp[]
}

defineProps<Props>()
const open = ref(true)
</script>
