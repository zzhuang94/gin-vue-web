<template>
  <table style="width: 100%; border: 0;">
    <tr v-if="name">
      <td style="padding: 0 5px 0 0; text-align: right; width: 100px; border: 0;">
        <code><b>改动对象</b></code>
      </td>
      <td style="padding: 0; border: 0; max-width: 500px; word-break: break-all;">
        <code><b>{{ name }}</b></code>
      </td>
    </tr>
    <tr v-for="d, i in diffs" :key="i">
      <td style="padding: 0 5px 0 0; text-align: right; width: 100px; border: 0;">
        <code><b>{{ d.rule.name }}</b></code>
      </td>
      <td style="padding: 0; border: 0; max-width: 500px; word-break: break-all;">
        <mark v-if="d.old !== d.new" style="display: inline-flex; align-items: center; flex-wrap: nowrap;">
          <Td :r="d.rule" :v="d.old" />
          <i class="fa fa-arrow-right" style="margin: 0 5px; flex-shrink: 0;"></i>
          <Td :r="d.rule" :v="d.new" />
        </mark>
        <span v-else><Td :r="d.rule" :v="d.old" /></span>
      </td>
    </tr>
  </table>
</template>

<script setup lang="ts">
import Td from '@components/td.vue'

interface Diff {
  rule: any
  old: any
  new: any
}

interface Props {
  name?: string
  diffs?: Diff[]
}

defineProps<Props>()
</script>