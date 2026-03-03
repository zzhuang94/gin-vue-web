<template>
  <a-dropdown :trigger="['click']">
    <button class="btn btn-default" style="padding: 0 5px">
      <span v-if="currentItem" :class="`badge badge-${currentItem.badge || 'default'}`" style="margin-right: 5px">{{ currentItem.label || value }}</span>
      <span v-else style="margin-right: 5px">{{ value || '请选择' }}</span>
      <span class="fa fa-angle-down" style="font-size: 1.2rem"></span>
    </button>
    <template #overlay>
      <a-menu @click="handleMenuClick">
        <a-menu-item v-for="v in limit" :key="v.key || v">
          <span :class="`badge badge-${v.badge || 'default'}`">{{ v.label || v }}</span>
        </a-menu-item>
      </a-menu>
    </template>
  </a-dropdown>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import type { RuleLimit } from '@libs/frm.ts'

interface Props {
  value: string
  limit: RuleLimit[] | null
  limit_map: Record<string, RuleLimit> | null
}

const props = defineProps<Props>()
const emit = defineEmits<{
  'change': []
  'update:value': [value: string]
}>();

const currentItem = computed(() => {
  if (!props.value || !props.limit_map) {
    return null
  }
  return props.limit_map[props.value] || null
})

const handleMenuClick = ({ key }: { key: string }) => {
  emit('update:value', key)
  emit('change')
}
</script>
