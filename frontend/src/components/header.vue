<template>
  <a-space :size="6">
    <button v-for="m, i in topMenus" :key="i":class="`btn btn-${m.color}`"
      @click="emit('menu-click', m)">
      <i :class="`fa fa-${m.icon}`"></i>
      {{ m.title }}
    </button>
    <button v-if="dump" :disabled="dumping" class="btn btn-accent" @click="dumpExcel">
      <i class="fa fa-file-excel"></i> 导出Excel
    </button>
  </a-space>
</template>

<script setup lang="ts">
import { ref, nextTick } from 'vue'
import type { Menu } from '@libs/frm'
import excel from '@libs/excel'

interface Props {
  topMenus: Menu[]
  dump: boolean
}

defineProps<Props>()
const emit = defineEmits<{
  'menu-click': [m: Menu]
}>()

const dumping = ref(false)

const dumpExcel = async () => {
  dumping.value = true
  await nextTick()
  excel.exportExcel()
  dumping.value = false
}
</script>