<template>
  <a-select
    v-model:value="value"
    :placeholder :disabled :class :style
    show-search allow-clear
    :filterOption="false"
    :notFoundContent="loading ? undefined : '无匹配结果'"
    @search="handleSearch"
    @dropdownVisibleChange="handleDropdownVisibleChange"
    >
    <template v-if="loading" #notFoundContent>
      <a-spin size="small" />
    </template>
    <a-select-option v-for="item in options" :key="item.key">
      {{ item.val }}
    </a-select-option>
  </a-select>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import type { RuleTrans } from '@libs/frm.ts'
import lib from '@libs/lib.ts'

interface Option {
  key: string
  val: string
}

interface Props {
  value?: string
  translate: RuleTrans
  placeholder?: string
  disabled?: boolean
  class?: string
  style?: Record<string, any>
  url?: string
}

const props = withDefaults(defineProps<Props>(), {
  value: '',
  placeholder: '请选择',
  disabled: false,
  class: '',
  url: ''
})

const emit = defineEmits<{
  'update:value': [value: string]
}>()

const options = ref<Option[]>([])
const loading = ref(false)
const lastFetchId = ref(0)

const value = computed({
  get: () => props.value,
  set: (newVal: string) => {
    emit('update:value', newVal)
  }
})

const handleSearch = (v: string) => {
  fetchData(v)
}

const handleDropdownVisibleChange = (open: boolean) => {
  if (open && !options.value.length) {
    fetchData('')
  }
}

const fetchData = async (term: string) => {
  loading.value = true
  lastFetchId.value += 1
  const fetchId = lastFetchId.value

  try {
    const tmp = await buildOptions(term)
    if (fetchId !== lastFetchId.value) {
      return
    }
    options.value = tmp
  } catch (error) {
    console.error('获取数据失败:', error)
  } finally {
    loading.value = false
  }
}

async function initOptions(): Promise<Option[]> {
  if (! props.value) {
    return []
  }
  const ans: Option[] = []
  let r : { code: number, data: Option }
  r = await lib.curl('/base/trans/init?val=' + props.value, props.translate)
  if (r.code == 1) {
    ans.push(r.data)
  } else if (props.translate.custom) {
    ans.push({ key: props.value, val: props.value })
  }
  return ans
}

async function loadData(term: string): Promise<Option[]> {
  let r: { data: Option[] } = { data: [] }
  if (props.url != '') {
    r = await lib.curl(props.url, { term })
  } else {
    r = await lib.curl('/base/trans/load?term=' + term, props.translate)
  }
  if (r.data) {
    return r.data
  }
  return []
}

async function buildOptions(term: string): Promise<Option[]> {
  const ans = await loadData(term)
  if (! props.translate.custom || term == '' || ans.length > 0 && ans[0]?.val == term) {
    return ans
  }
  ans.unshift({ key: term, val: term })
  return ans
}

onMounted(async () => {
  options.value = await initOptions()
})
</script>
