<template>
  <a-collapse v-model:activeKey="activeKey" accordion>
    <template v-for="d, k in data">
      <a-collapse-panel v-if="! isEmpty(d.data)" :key="k">
        <template #header><code>{{ d.title }}</code></template>
        <Table :data="d.data" :rules="d.config" noSort />
      </a-collapse-panel>
    </template>
  </a-collapse>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { isEmpty } from 'lodash'
import Table from '@components/table.vue'

interface DataItem {
  title: string
  data: any[]
  config: any
}

interface Props {
  data: Record<string, DataItem>
}

const props = defineProps<Props>()
const activeKey = ref<string[]>(initActiveKey())

function initActiveKey(): string[] {
  for (let k in props.data) {
    if (props.data[k] && ! isEmpty(props.data[k].data)) {
      return [k]
    }
  }
  return []
}
</script>
