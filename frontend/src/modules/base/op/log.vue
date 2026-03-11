<template>
  <LogDiffs v-if="! detail" :name="name" :diffs="diffs" />
  <tr v-else>
    <td>{{ log.id }}</td>
    <td>{{ name }}</td>
    <td><Td :r="opRule" :v="log.op" /></td>
    <td>{{ log.data_table }}</td>
    <td>{{ log.data_id }}</td>
    <td><LogDiffs :diffs="diffs" /></td>
    <td>{{ time }}</td>
  </tr>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import type { Rule } from '@libs/frm'
import lib from '@libs/lib'

import Td from '@components/td.vue'
import LogDiffs, { type Diff } from './log-diffs.vue'

interface Props {
  id: string
  detail?: boolean
  opRule: Rule
}

const props = withDefaults(defineProps<Props>(), {
  detail: false,
})

const name = ref('')
const log = ref<Record<string, any>>({})
const diffs = ref<Diff[]>([])
const time = ref('')

const fetchLog = async () => {
  const r = await lib.curl(`log?id=${props.id}`)
  name.value = r.name
  log.value = r.log
  diffs.value = r.diffs
  time.value = r.time
}

onMounted(fetchLog)
watch(() => props.id, fetchLog)
</script>
