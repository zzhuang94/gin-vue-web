<template>
  <a-switch :checked="isChecked" @change="handleChange" :disabled />
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  modelValue: { type: [String, Number, Boolean], required: true },
  trueValue: { type: [String, Number, Boolean], default: '1' },
  falseValue: { type: [String, Number, Boolean], default: '0' },
  disabled: { type: Boolean, default: false }
})

const emit = defineEmits(['update:modelValue'])

const isChecked = computed(() => {
  return props.modelValue == props.trueValue
})

const handleChange = (checked) => {
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
