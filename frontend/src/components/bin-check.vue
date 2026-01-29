<template>
  <a-checkbox :checked="isChecked" @change="handleChange">
    <slot />
  </a-checkbox>
</template>

<script setup lang="ts">
import { computed } from 'vue';

interface Props {
  modelValue: string | number | boolean
  trueValue?: string | number | boolean
  falseValue?: string | number | boolean
}

const props = withDefaults(defineProps<Props>(), {
  trueValue: '1',
  falseValue: '0'
});

const emit = defineEmits<{
  'update:modelValue': [value: string | number | boolean]
}>();

// 计算当前是否选中
const isChecked = computed(() => {
  return props.modelValue == props.trueValue
})

// 处理复选框变化
const handleChange = (e: { target: { checked: boolean } }) => {
  const newValue = e.target.checked ? props.trueValue : props.falseValue;
  emit('update:modelValue', newValue);
};
</script>
