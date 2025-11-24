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

<script setup>
import Td from '@components/td.vue'
const props = defineProps({
  id: { type: String, default: '' },
  rules: Array,
  data: Object,
})
</script>
