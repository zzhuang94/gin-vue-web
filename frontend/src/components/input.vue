<template>
  <div class="form-input-container">
    <div v-if="label != ''" class="custom-label" :style="{ width: labelWidth }">{{ label }}</div>

    <a-input
      v-if="type === 'input'"
      v-model:value="val"
      :placeholder="placeholder"
      :disabled="disabled"
      :class="label != '' ? 'has-label' : ''"
    />

    <a-textarea
      v-else-if="type === 'textarea'"
      v-model:value="val"
      :placeholder="placeholder"
      :auto-size="{ minRows: minRows }"
      :disabled="disabled"
      :class="label != '' ? 'has-label' : ''"
    />

    <a-select
      v-else-if="type === 'select'"
      v-model:value="val"
      :placeholder="placeholder"
      :options="options"
      :mode="multiple ? 'multiple' : undefined"
      :filterOption="lib.filterByLabel"
      :disabled="disabled"
      style="width: 100%"
      :search-value="searchText"
      @search="handleSearch"
      :class="['input-select', label != '' ? 'has-label' : '']"
      show-search
      allow-clear
    />
    <div v-else-if="type === 'checkbox'" :class="['input-checkbox', label != '' ? 'has-label' : '']">
      <a-checkbox v-model:checked="val" :disabled="disabled" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import lib from '@libs/lib.ts'

interface Option {
  label: string
  value: string | number
}

interface Props {
  label?: string
  type?: 'input' | 'textarea' | 'select' | 'checkbox'
  value?: string | string[]
  placeholder?: string
  labelWidth?: string
  minRows?: string
  options?: Option[]
  multiple?: boolean
  disabled?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  label: '',
  type: 'input',
  value: '',
  placeholder: '请输入内容',
  labelWidth: '100px',
  minRows: '1',
  options: () => [],
  multiple: false,
  disabled: false
})

const emit = defineEmits<{
  'update:value': [value: string | string[]]
  'change': [value: string | string[]]
}>()

const val = ref<string | string[]>(props.value ?? '')

const searchText = ref('')
const handleSearch = (value: string) => {
  searchText.value = value
}

watch(() => props.value, (newVal) => { val.value = newVal ?? '' })
watch(val, (newVal) => {
  emit('update:value', newVal)
  emit('change', newVal)
})
</script>

<style scoped>
.form-input-container {
  display: flex;
  align-items: center;
}

.custom-label {
  border: 1px solid #d9d9d9;
  border-right: none;
  border-radius: 6px 0 0 6px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  background-color: #f5f5f5;
  font-weight: bold;
  height: 32px;
  line-height: 32px;
}

.input-checkbox {
  border: 1px solid #d9d9d9;
  border-radius: 0 6px 6px 0;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 32px;
  line-height: 32px;
}

.ant-input, .input-checkbox {
  min-height: 32px !important;
}
.ant-input.has-label, .input-checkbox.has-label {
  border-radius: 0 6px 6px 0 !important;
}
</style>

<style>
.input-select>.ant-select-selector {
  min-height: 32px;
}
.input-select.has-label>.ant-select-selector {
  border-radius: 0 6px 6px 0 !important;
}
</style>
