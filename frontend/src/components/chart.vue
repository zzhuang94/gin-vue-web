<template>
  <span v-show="! init">图表加载中...</span>
  <Chart :option :url :width :height :margin :china :loading />
</template>

<script setup>
import { ref, defineAsyncComponent } from 'vue'

const props = defineProps({
  option: { type: Object, default: null },
  url: { type: String, default: '' },
  width: { type: String, default: '100%' },
  height: { type: String, default: '100%' },
  margin: { type: String, default: '0' },
  china: { type: Boolean, default: false },
  loading: { type: Boolean, default: false },
})

const init = ref(false)

const Chart = defineAsyncComponent(() => 
  import('@components/_chart.vue').finally(() => {
    init.value = true
  })
)
</script>
