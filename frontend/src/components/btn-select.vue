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

<script setup>
import { ref, watch } from 'vue'

const emit = defineEmits(['update'])
const props = defineProps({
  name: {
    type: String,
    required: true
  },
  items: {
    type: Array,
    required: true,
    validator: (value) => value.length > 0
  },
  default: {
    type: String,
    default: ''
  },
  required: {
    type: Boolean,
    default: false
  }
})

const selected = ref(initSelected())

watch(() => props.default, (newVal) => { selected.value = newVal })
watch(selected, (newVal) => { emit('update', newVal) })

const isSelected = (key) => {
  return selected.value === key
}

const handleClick = (key) => {
  if (props.required) {
    if (! isSelected(key)) {
      selected.value = key
    }
  } else {
    selected.value = isSelected(key) ? '' : key
  }
}

function initSelected() {
  if (props.default != '') {
    return props.default
  }
  if (props.required) {
    return props.items[0].key
  }
  return ''
}

defineExpose({
  getSelected: () => selected.value,
  setSelected: (key) => { selected.value = key }
})
</script>
