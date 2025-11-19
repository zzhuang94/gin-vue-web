<template>
  <a-checkbox :checked="isChecked" @change="handleChange">
    <slot />
  </a-checkbox>
</template>

<script setup>
import { computed } from 'vue';

const props = defineProps({
  modelValue: { type: [String, Number, Boolean], required: true },
  // 自定义真假值映射（默认支持 1/0）
  trueValue: { type: [String, Number, Boolean], default: '1' },
  falseValue: { type: [String, Number, Boolean], default: '0' },
});

const emit = defineEmits(['update:modelValue']);

// 计算当前是否选中
const isChecked = computed(() => {
  return props.modelValue == props.trueValue
})

// 处理复选框变化
const handleChange = (e) => {
  const newValue = e.target.checked ? props.trueValue : props.falseValue;
  emit('update:modelValue', newValue);
};
</script>
