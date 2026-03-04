<template>
  <button class="btn" :class="`btn-${menu.color}`" @click="handleClick">
    <i :class="`fa fa-${menu.icon}`"></i>
    {{ menu.title }}
  </button>
</template>

<script setup lang="ts">
import { inject } from 'vue'
import type { Menu } from '@libs/frm.ts'
interface Props {
  menu: Menu
}

const props = defineProps<Props>()
const menuClick = inject<(menu: Menu) => void>('menuClick')
const handleClick = () => {
  if (menuClick) {
    menuClick(props.menu)
  }
}
</script>
