<template>
  <a-switch :checked="isChecked" @change="handleChange" :disabled />
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface Props {
  modelValue: string | number | boolean
  trueValue?: string | number | boolean
  falseValue?: string | number | boolean
  disabled?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  trueValue: '1',
  falseValue: '0',
  disabled: false
})

const emit = defineEmits<{
  'update:modelValue': [value: string | number | boolean]
  'change': [value: string | number | boolean]
}>()

const isChecked = computed(() => {
  return props.modelValue == props.trueValue
})

const handleChange = (checked: boolean) => {
  const newValue = checked ? props.trueValue : props.falseValue
  emit('update:modelValue', newValue)
  emit('change', newValue)
}
</script>

<style scoped>
.ant-switch-checked {
  background-color: #34bfa3 !important;
}
</style>
