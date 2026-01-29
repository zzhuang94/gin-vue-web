<template>
  <div style="margin: 3px 0">
    <button class="btn btn-sm"><b>{{ name }}：</b></button>
    <button
      v-for="i in items"
      :key="i.key"
      class="btn btn-sm"
      :class="{'btn-success': isSelected(i.key), 'btn-secondary': !isSelected(i.key)}"
      @click="handleClick(i.key)"
    >
      {{ i.label }}
    </button>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'

interface Item {
  key: string | number
  label: string
}

interface Props {
  name: string
  items: Item[]
  default?: string
  required?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  default: '',
  required: false
})

const emit = defineEmits<{
  'update': [value: string]
}>()

const selected = ref<string>(initSelected())

watch(() => props.default, (newVal) => { selected.value = newVal })
watch(selected, (newVal) => { emit('update', newVal) })

const isSelected = (key: string | number) => {
  return selected.value === key
}

const handleClick = (key: string | number) => {
  if (props.required) {
    if (! isSelected(key)) {
      selected.value = String(key)
    }
  } else {
    selected.value = isSelected(key) ? '' : String(key)
  }
}

function initSelected(): string {
  if (props.default != '') {
    return props.default
  }
  if (props.required && props.items.length > 0 && props.items[0]) {
    return String(props.items[0].key)
  }
  return ''
}

defineExpose({
  getSelected: () => selected.value,
  setSelected: (key: string | number) => { selected.value = String(key) }
})
</script>
