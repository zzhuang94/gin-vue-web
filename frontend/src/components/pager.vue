<template>
  <div>
    <a-config-provider :locale="zhCN">
      <a-pagination
        v-model:current="curr"
        v-model:pageSize="size"
        :total="page.total"
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
import { h, computed, watch } from 'vue'
import type { Page } from '@libs/frm.ts'
import zhCN from 'ant-design-vue/es/locale/zh_CN'
import lib from '@libs/lib.ts'

interface Props {
  loading: boolean
  page: Page
}

const props = defineProps<Props>()
const emit = defineEmits<{'update:page': [value: Page]}>()

const showTotal = (total: number) => {
  return h('span', ['共', h(
    'strong',
    { style: 'font-size: 1.1rem; padding: 0 8px;' },
    total.toLocaleString()
  ), '条'])
}

const curr = computed({
  get: () => props.page.curr,
  set: (value: number) => {
    emit('update:page', { ...props.page, curr: value })
  },
})

const size = computed({
  get: () => props.page.size,
  set: (value: number) => {
    emit('update:page', { ...props.page, size: value })
  },
})

watch(
  () => size.value,
  (newSize: number) => {
    lib.curl(`/base/user/set?key=page_size&val=${newSize}`)
  },
)

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
