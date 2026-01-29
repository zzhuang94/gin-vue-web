<template>
  <button class="btn" :class="`btn-${color}`" @click="handleClick">
    <i :class="`fa fa-${icon}`"></i>
    {{ title }}
  </button>
</template>

<script setup lang="ts">
import { inject } from 'vue'

interface Props {
  title?: string
  icon?: string
  url?: string
  type?: string
  color?: string
}

const props = defineProps<Props>()
const toolClick = inject<(props: Props) => void>('toolClick')
const handleClick = () => {
  if (toolClick) {
    toolClick(props)
  }
}
</script>
