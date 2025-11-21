<template>
  <table style="width: 100%; border: 0;">
    <tr><td colspan="2" class="table-op-col"><b>{{ op.name }}</b></td></tr>
    <tr v-for="r, i in rules" :key="i">
      <td style="padding: 0 5px 0 0; text-align: right; width: 100px; border: 0;">
        {{ r.name }}
      </td>
      <td style="padding: 0 0 0 5px; text-align: left; width: 100px; border: 0;">
        {{ r.key }}
      </td>
    </tr>
  </table>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'

import lib from '@libs/lib.ts'

const props = defineProps(['id', 'ops'])
const op = ref({})
const log = ref({})
const rules = ref([])
const dnew = ref({})
const dold = ref({})

const keys = computed(() => {
  
})

onMounted(async () => {
  const r = await lib.curl(`log?id=${props.id}`)
  log.value = r.log
  rules.value = r.rules
  op.value = props.ops[r.log.data_table]
  dold.value = JSON.parse(r.log.data_old)
  dnew.value = JSON.parse(r.log.data_new)
})
</script>