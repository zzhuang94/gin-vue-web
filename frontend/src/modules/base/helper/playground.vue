<template>
  <div>
    Name: {{ name }}
    Age: {{ age }}
    <hr />
    <SubComp v-model:name="name" v-model:age="age" />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import SubComp from './sub-comp.vue'

const name = ref<string>('John')
const age = ref<number>(20)
</script>
