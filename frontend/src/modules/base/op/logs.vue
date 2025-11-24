<template>
  <a-modal width="90%" :maskClosable="false" :open="open" title="改动详情" @cancel="cancel" :footer="null">
    <hr />
    <table class="table table-hover" style="width: 100%; border: 0;">
      <thead>
        <tr>
          <th>ID</th>
          <th>操作对象</th>
          <th>操作类型</th>
          <th>数据库表</th>
          <th>数据ID</th>
          <th style="padding-left: 100px">改动详情</th>
          <th>操作时间</th>
        </tr>
      </thead>
      <tbody>
        <Log v-for="lid in lids" :key="lid" :id="lid" :detail="true" :op_rule />
      </tbody>
    </table>
  </a-modal>
</template>

<script setup>
import { ref, computed } from 'vue'
import lib from '@libs/lib.ts'
import Log from './log.vue'

const props = defineProps(['open', 'lids', 'log_rules'])
const emit = defineEmits(['update:open'])

const lids = computed(() => {
  return props.lids.split(',')
})

const op_rule = computed(() => {
  return props.log_rules.find(r => r.key === 'op')
})

const cancel = () => {
  emit('update:open', false);
}
</script>