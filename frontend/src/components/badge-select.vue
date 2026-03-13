<template>
  <a-dropdown :trigger="['click']">
    <button class="btn btn-default" style="padding: 0 5px">
      <span v-if="curr" :class="`badge badge-${curr.badge}`">{{ curr.label }}</span>
      <span v-else>{{ value || '请选择' }}</span>
      <span class="fa fa-angle-down" style="font-size: 1.2rem"></span>
    </button>
    <template #overlay>
      <a-menu @click="handleMenuClick">
        <a-menu-item v-for="l in r.limit" :key="l.key">
          <span :class="`badge badge-${l.badge}`">{{ l.label }}</span>
        </a-menu-item>
      </a-menu>
    </template>
  </a-dropdown>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import type { Rule } from '@libs/frm'

interface Props {
  value: string | undefined
  r: Rule
}

const props = defineProps<Props>()
const emit = defineEmits<{
  'update:value': [value: string]
  'change': []
}>();

const curr = computed(() => {
  if (!props.value || !props.r.limit_map) {
    return null
  }
  return props.r.limit_map[props.value]
})

const handleMenuClick = ({ key }: { key: string }) => {
  emit('update:value', key)
  emit('change')
}
</script>

<style scoped>
span {
  margin-right: 5px;
}
</style>