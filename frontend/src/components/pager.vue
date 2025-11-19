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

<script setup>
import { h, ref, inject, computed, watch } from 'vue';
import zhCN from 'ant-design-vue/es/locale/zh_CN';

const locale = zhCN;
const props = defineProps(['loading', 'curr', 'size', 'total']);
const emit = defineEmits(['update:curr', 'update:size', 'page-change'])

const showTotal = (total) => {
  return h('span', ['共', h('strong', {style: 'font-size: 1.1rem; padding: 0 8px;'}, total), '条']);
}
const curr = computed({get: () => props.curr, set: (value) => emit('update:curr', value)})
const size = computed({get: () => props.size, set: (value) => emit('update:size', value)})
watch(() => [curr.value, size.value], () => emit('page-change'))

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
