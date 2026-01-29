<template>
  <table class="table table-hover table-bordered" :id="id ? id : undefined">
    <tbody>
      <tr v-for="r in rules">
        <td><b>{{ r.name }}</b></td>
        <td><Td :r :v="data[r.key]" /></td>
      </tr>
    </tbody>
  </table>
</template>

<script setup lang="ts">
import Td from '@components/td.vue'

interface Rule {
  name: string
  key: string
  [key: string]: any
}

interface Props {
  id?: string
  rules: Rule[]
  data: Record<string, any>
}

withDefaults(defineProps<Props>(), {
  id: ''
})
</script>
