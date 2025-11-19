<template>
  <button class="btn" :class="`btn-${color}`" @click="toolClick(props)">
    <i :class="`fa fa-${icon}`"></i>
    {{ title }}
  </button>
</template>

<script setup>
import { inject } from 'vue'

const props = defineProps(['title', 'icon', 'url', 'type', 'color'])
const toolClick = inject('toolClick')
</script>
