<template>
  <span v-show="! init">图表加载中...</span>
  <Chart :option :url :width :height :margin :china :loading />
</template>

<script setup lang="ts">
import { ref, defineAsyncComponent } from 'vue'

interface Props {
  option?: Record<string, any>
  url?: string
  width?: string
  height?: string
  margin?: string
  china?: boolean
  loading?: boolean
}

withDefaults(defineProps<Props>(), {
  url: '',
  width: '100%',
  height: '100%',
  margin: '0',
  china: false,
  loading: false
})

const init = ref(false)

const Chart = defineAsyncComponent(() => 
  import('@components/_chart.vue').finally(() => {
    init.value = true
  })
)
</script>
