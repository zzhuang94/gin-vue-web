<template>
  <table class="table table-hover table-bordered" :id="id ? id : undefined">
    <tbody>
      <tr v-for="v, k in config">
        <td><b>{{ v.name }}</b></td>
        <td><Td :d="data" :k :v /></td>
      </tr>
    </tbody>
  </table>
</template>

<script setup>
import Td from '@components/td.vue'
const props = defineProps({
  id: { type: String, default: '' },
  config: Object,
  data: Array,
})
</script>
