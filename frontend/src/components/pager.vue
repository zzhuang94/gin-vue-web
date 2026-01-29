<template>
  <div>
    <a-config-provider :locale="locale">
      <a-pagination
        v-model:current="curr"
        v-model:pageSize="size"
        :total="total"
        :disabled="loading"
        show-quick-jumper
        show-size-changer
        :showTotal="showTotal"
        :pageSizeOptions="[5, 10, 20, 50, 100, 200, 500]"
        class="my-a-pagination"
        />
    </a-config-provider>
  </div>
</template>

<script setup lang="ts">
import { h, computed, watch } from 'vue';
import zhCN from 'ant-design-vue/es/locale/zh_CN';
import lib from '@libs/lib.ts';

const locale = zhCN;

interface Props {
  loading?: boolean
  curr?: number
  size?: number
  total?: number
}

const props = defineProps<Props>()
const emit = defineEmits<{
  'update:curr': [value: number]
  'update:size': [value: number]
  'page-change': []
}>()

const showTotal = (total: number) => {
  return h('span', ['共', h('strong', {style: 'font-size: 1.1rem; padding: 0 8px;'}, total), '条']);
}

const curr = computed({get: () => props.curr ?? 1, set: (value: number) => emit('update:curr', value)})
const size = computed({get: () => props.size ?? 10, set: (value: number) => emit('update:size', value)})
watch(() => [curr.value, size.value], () => emit('page-change'))

watch(() => size.value, (newSize: number) => {
  lib.curl(`/base/user/set?key=page_size&val=${newSize}`)
})

</script>

<style scoped>
.my-a-pagination {
  margin: 8px 0 0 5px;
}
</style>
<style>
.my-a-pagination .ant-pagination-options-size-changer {
  min-width: 100px;
}
.my-a-pagination .ant-pagination-options-quick-jumper input {
  min-width: 80px;
}
</style>
