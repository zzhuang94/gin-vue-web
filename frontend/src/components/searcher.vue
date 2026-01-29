<template>
  <div class="searcher">
    <a-row align="middle">
      <a-col :span="2" class="label">筛选条件：</a-col>
      <a-col :span="22">
        <a-space :size="4">
          <label v-for="v in rules" :key="v.key" @click="toggle(v.key)"
            class="btn btn-sm" :class="visibleKeys.has(v.key) ? 'btn-success' : 'btn-default'">
            {{ v.name }}
          </label>
        </a-space>
      </a-col>
    </a-row>

    <a-row align="middle" v-for="v in rules" :key="v.key" v-show="visibleKeys.has(v.key)">
      <a-col :span="2" class="label">{{ v.name }}：</a-col>
      <a-col :span="22">

          <a-select v-if="v.limit && v.limit.length"
            v-model:value="formData[v.key]"
            show-search
            allow-clear
            :search-value="searchText"
            @search="handleSearch"
            :filterOption="lib.filterByLabel"
            :placeholder="`请选择${v.name}`"
            :mode="(typeof v.search === 'number' && v.search === 3) ? 'multiple' : undefined"
            class="search-input"
            >
          <a-select-option v-for="lv in v.limit" :key="lv.key" :value="lv.key" :label="lv.label">
            {{ lv.label }}
          </a-select-option>
        </a-select>

        <AjaxSelect
          v-else-if="v.trans && v.trans.ajax"
          v-model:value="formData[v.key]"
          :placeholder="`请选择${v.name}`"
          :translate="v.trans"
          class="search-input"
          />

        <a-input v-else
          v-model:value="formData[v.key]"
          :placeholder="`请输入${v.name}`"
          class="search-input"
          />

      </a-col>
    </a-row>

    <a-row>
      <a-col :span="2"></a-col>
      <a-col :span="22">
        <a-space :size="4">
          <button class="btn btn-sm btn-primary" @click="emit('search')">
            <i class="fa fa-search"></i> 搜 索
          </button>
          <button class="btn btn-sm btn-secondary" @click="clear">
            <i class="fa fa-refresh"></i> 重 置
          </button>
        </a-space>
      </a-col>
    </a-row>

  </div>
</template>

<script setup lang="ts">
import { watch, computed, onMounted, reactive, ref } from 'vue'
import lib from '@libs/lib.ts'
import AjaxSelect from '@components/ajax-select.vue';

interface Rule {
  key: string
  name: string
  search?: boolean
  limit?: any[]
  trans?: any
  [key: string]: any
}

interface Props {
  rules: Rule[]
  arg: Record<string, any>
  bind?: any
}

const props = defineProps<Props>()
const emit = defineEmits<{
  'update:arg': [arg: Record<string, any>]
  'search': []
  'clear': []
}>()

const searchText = ref('')
const handleSearch = (value: string) => {
  searchText.value = value
}

const visibleKeys = ref<Set<string>>(new Set())
const formData = reactive<Record<string, any>>(props.arg)

const rules = computed(() => {
  return props.rules.filter(i => i.search)
})

const initVisibleKeys = () => {
  visibleKeys.value.clear()
  // 先找出有值的字段
  for (const item of rules.value) {
    if (props.arg && props.arg[item.key] !== undefined) {
      visibleKeys.value.add(item.key)
    }
  }
  // 如果没有有值的字段，默认显示第一个
  if (visibleKeys.value.size === 0 && rules.value.length > 0 && rules.value[0]) {
    visibleKeys.value.add(rules.value[0].key)
  }
}

const toggle = (key: string) => {
  if (visibleKeys.value.has(key)) {
    visibleKeys.value.delete(key)
    delete formData[key]
  } else {
    visibleKeys.value.add(key)
  }
}

watch(formData, (newFormData) => { emit('update:arg', newFormData) })

const clear = () => {
  Object.keys(formData).forEach(key => {
    delete formData[key]
  })
  initVisibleKeys()
  emit('clear')
}

onMounted(initVisibleKeys)
</script>

<style scoped>
.searcher {
  margin-left: 5px;
  margin-bottom: 5px;
}
.label {
  font-weight: bold;
}
.ant-row {
  padding: 2px 0;
}
.search-input {
  width: 480px;
}
.ant-space:not(.auto-hide-l1) {
  display: flex !important;
  flex-wrap: wrap !important;
}
</style>

