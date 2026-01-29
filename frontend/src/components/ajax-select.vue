<template>
  <a-select
    v-model:value="val"
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
    <a-select-option v-for="item in options" :key="item.value" :value="item.value">
      {{ item.label }}
    </a-select-option>
  </a-select>
</template>

<script setup lang="ts">
import { ref, watch, onMounted } from 'vue'
import lib from '@libs/lib.ts'

interface Option {
  value: string | number
  label: string
}

interface Translate {
  ajax?: boolean
  custom?: boolean
  [key: string]: any
}

interface Props {
  value?: string
  placeholder?: string
  disabled?: boolean
  class?: string
  style?: Record<string, any>
  translate?: Translate
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
  'change': [value: string]
}>()

const val = ref<string>(props.value ?? '')
const options = ref<Option[]>([])
const loading = ref(false)
const lastFetchId = ref(0)

watch(() => props.value, (newVal) => { val.value = newVal ?? '' })
watch(val, (newVal) => {
  emit('update:value', newVal)
  emit('change', newVal)
})

const handleSearch = (value: string) => {
  fetchData(value)
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
  const r = await lib.curl('/base/trans/init?val=' + props.value, props.translate)
  if (r.code == 1) {
    ans.push({ value: r.data.key, label: r.data.val })
  } else if (props.translate?.custom) {
    ans.push({ value: props.value, label: props.value })
  }
  return ans
}

async function loadData(term: string): Promise<any[]> {
  let r: any = {}
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
  const data = await loadData(term)
  let ans = data.map((d: any) => ({ value: d.key, label: d.val }))
  if (! (props.translate?.custom ?? 0) || term == '' || ans.length > 0 && ans[0]?.label == term) {
    return ans
  }
  ans.unshift({ value: term, label: term })
  return ans
}

onMounted(async () => {
  options.value = await initOptions()
})
</script>
