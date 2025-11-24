<template>
  <LogDiffs v-if="! detail" :name="name" :diffs="diffs" />
  <tr v-else>
    <td>{{ log.id }}</td>
    <td>{{ name }}</td>
    <td><Td :r="op_rule" :v="log.op" /></td>
    <td>{{ log.data_table }}</td>
    <td>{{ log.data_id }}</td>
    <td><LogDiffs :diffs="diffs" /></td>
    <td>{{ log.created }}</td>
  </tr>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'

import lib from '@libs/lib.ts'
import Td from '@components/td.vue'
import LogDiffs from './log-diffs.vue'

const props = defineProps({
  id: { type: String, default: '' },
  detail: { type: Boolean, default: false },
  op_rule: { type: Object, default: null },
})
const name = ref('')
const log = ref({})
const diffs = ref([])

const fetchLog = async () => {
  if (!props.id) return
  const r = await lib.curl(`log?id=${props.id}`)
  name.value = r.name
  log.value = r.log
  diffs.value = r.diffs
}

onMounted(fetchLog)
watch(() => props.id, fetchLog)
</script>