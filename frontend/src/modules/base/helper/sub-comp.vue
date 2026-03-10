<template>
  <div>
    Sub-Name: {{ name }}
    Sub-Age: {{ age }}
    <input v-model="name" />
    <input type="number" v-model="age" />
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface Props {
  name: string
  age: number
}

const props = defineProps<Props>()
const emit = defineEmits<{
  'update:name': [value: string]
  'update:age': [value: number]
}>()

const name = computed({
  get: () => props.name,
  set: (v) => emit('update:name', v),
})

const age = computed({
  get: () => props.age,
  set: (v) => emit('update:age', v),
})
</script>