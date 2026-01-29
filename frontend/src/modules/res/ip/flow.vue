<template>
  <a-modal width="80%" maskClosable="false" v-model:open="open" :footer="null">
    <template #title>
      <i class="fa fa-chart-line"></i> <code>{{ ip }}</code> 流量图
    </template>
    <hr />
    <Chart :option="option" height="600px" />
  </a-modal>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Chart from '@components/chart.vue'

interface Props {
  ip?: string
  option?: Record<string, any>
}

defineProps<Props>()
const open = ref(true)
</script>
