<template>
  <a-collapse v-model:activeKey="activeKey" accordion>
    <template v-for="d, k in data">
      <a-collapse-panel v-if="! isEmpty(d.data)" :key="k">
        <template #header><code>{{ d.title }}</code></template>
        <Table :data="d.data" :config="d.config" noSort />
      </a-collapse-panel>
    </template>
  </a-collapse>
</template>

<script setup>
import { ref } from 'vue'
import { isEmpty } from 'lodash'
import Table from '@components/table.vue'

const props = defineProps(['data'])
const activeKey = ref(initActiveKey())

function initActiveKey() {
  for (let k in props.data) {
    if (! isEmpty(props.data[k].data)) {
      return [k]
    }
  }
  return []
}
</script>
