<template>
  <a-spin :spinning="loading" tip="加载中..." size="large">
    <div ref="chartRef" :style="{ width, height, margin }" />
  </a-spin>
</template>

<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount, watch } from 'vue'
import * as echarts from 'echarts'
import china from '@libs/china.json'
import { isEmpty } from 'lodash'

interface Props {
  option?: Record<string, any>
  url?: string
  width?: string
  height?: string
  margin?: string
  china?: boolean
  loading?: boolean
}

const props = defineProps<Props>()

let chart: echarts.ECharts | null = null
const chartRef = ref<HTMLElement | null>(null)

const getOption = async (): Promise<Record<string, any>> => {
  if (props.option) {
    return props.option
  }
  return await loadOption(props.url ?? '')
}

const loadOption = async (url: string): Promise<Record<string, any>> => {
  const response = await fetch(url)
  return await response.json()
}

const initChart = async () => {
  if (props.china) {
    echarts.registerMap('china', china as any)
  }
  if (chartRef.value) {
    chart = echarts.init(chartRef.value)
    chart.setOption(await getOption());
    chart.on("click", function(e: any){
      if (!isEmpty(e.data?.url)) {
        window.open(e.data.url)
      }
    })
    window.addEventListener('resize', () => chart?.resize())
  }
}

watch(
  () => props.option,
  (newOption) => {
    if (chart && newOption) {
      chart.setOption(newOption)
    }
  },
  { deep: true }
)

watch(
  () => props.url,
  async (newUrl) => {
    if (chart && newUrl) {
      chart.setOption(await loadOption(newUrl))
    }
  },
  { deep: true }
)

onMounted(initChart)
onBeforeUnmount(() => {
  if (chart) {
    chart.dispose()
    chart = null
  }
})
</script>
