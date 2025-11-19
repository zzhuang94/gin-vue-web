<template>
  <a-spin :spinning="loading" tip="加载中..." size="large">
    <div ref="chartRef" :style="{ width, height, margin }" />
  </a-spin>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount, watch } from 'vue'
import * as echarts from 'echarts'
import china from '@libs/china.json'

const props = defineProps(['option', 'url', 'width', 'height', 'margin', 'china', 'loading'])

let chart = null
const chartRef = ref(null)

const getOption = async () => {
  if (props.option) {
    return props.option
  }
  return await loadOption(props.url)
}

const loadOption = async (url) => {
  const response = await fetch(url)
  return await response.json()
}

const initChart = async () => {
  if (props.china) {
    echarts.registerMap('china', china)
  }
  chart = echarts.init(chartRef.value)
  chart.setOption(await getOption());
  chart.on("click", function(e){
    if (! _.isEmpty(e.data.url)) {
      window.open(e.data.url)
    }
  })
  window.addEventListener('resize', () => chart.resize())
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
