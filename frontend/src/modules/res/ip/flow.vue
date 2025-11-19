<template>
  <a-modal width="80%" maskClosable="false" v-model:open="open" :footer="null">
    <template #title>
      <i class="fa fa-chart-line"></i> <code>{{ ip }}</code> 流量图
    </template>
    <hr />
    <Chart :option="option" height="600px" />
  </a-modal>
</template>

<script setup>
import { ref } from 'vue'
import Chart from '@components/chart.vue'
defineProps(['ip', 'option'])
const open = ref(true)
</script>
