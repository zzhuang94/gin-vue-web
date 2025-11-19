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

<script setup>
import { ref, watch, onMounted } from 'vue'
import lib from '@libs/lib.ts'

const emit = defineEmits(['update:value', 'change'])
const props = defineProps({
  value: { type: String, default: '' },
  placeholder: { type: String, default: '请选择' },
  disabled: { type: Boolean, default: false },
  class: { type: String, default: '' },
  style: { type: Object },
  translate: { type: Object },
  url: { type: String, default: '' },
})

const val = ref(props.value)
const options = ref([])
const loading = ref(false)
const lastFetchId = ref(0)

watch(() => props.value, (newVal) => { val.value = newVal })
watch(val, (newVal) => {
  emit('update:value', newVal)
  emit('change', newVal)
})

const handleSearch = (value) => {
  fetchData(value)
}

const handleDropdownVisibleChange = (open) => {
  if (open && !options.value.length) {
    fetchData('')
  }
}

const fetchData = async (term) => {
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

async function initOptions() {
  if (! props.value) {
    return []
  }
  const ans = []
  const r = await lib.curl('/base/trans/init?val=' + props.value, props.translate)
  if (r.code == 1) {
    ans.push({ value: r.data.key, label: r.data.val })
  } else if (props.translate.custom) {
    ans.push({ value: props.value, label: props.value })
  }
  return ans
}

async function loadData(term) {
  let r = {}
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

async function buildOptions(term) {
  const data = await loadData(term)
  let ans = data.map(d => ({ value: d.key, label: d.val }))
  if (! (props.translate?.custom ?? 0) || term == '' || ans.length > 0 && ans[0].label == term) {
    return ans
  }
  ans.unshift({ value: term, label: term })
  return ans
}

onMounted(async () => {
  options.value = await initOptions()
})
</script>
