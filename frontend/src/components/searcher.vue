<template>
  <div class="searcher">
    <a-row align="middle">
      <a-col :xs="8" :md="4" class="label">筛选条件：</a-col>
      <a-col :xs="16" :md="20">
        <a-space :size="4" class="filter-tags">
          <label v-for="r in rules" :key="r.key" @click="toggle(r.key)"
            class="btn btn-sm" :class="keys.has(r.key) ? 'btn-success' : 'btn-default'">
            {{ r.name }}
          </label>
        </a-space>
      </a-col>
    </a-row>

    <a-row align="middle" v-for="r in rules" :key="r.key" v-show="keys.has(r.key)">
      <a-col :xs="8" :md="4" class="label">{{ r.name }}：</a-col>
      <a-col :xs="16" :md="20">

          <a-select v-if="r.limit"
            v-model:value="arg[r.key]"
            show-search
            allow-clear
            :filterOption="lib.filterByLabel"
            :placeholder="`请选择${r.name}`"
            :mode="r.search === 3 ? 'multiple' : undefined"
            class="search-input"
            >
          <a-select-option v-for="lv in r.limit" :key="lv.key" :label="lv.label">
            {{ lv.label }}
          </a-select-option>
        </a-select>

        <AjaxSelect
          v-else-if="r.trans && r.trans.ajax"
          v-model:value="arg[r.key] as string"
          :placeholder="`请选择${r.name}`"
          :translate="r.trans"
          class="search-input"
          />

        <a-input v-else
          v-model:value="arg[r.key] as string"
          :placeholder="`请输入${r.name}`"
          class="search-input"
          />

      </a-col>
    </a-row>

    <a-row>
      <a-col :xs="8" :md="4"></a-col>
      <a-col :xs="16" :md="20">
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
import { computed, onMounted, ref } from 'vue'
import type { Rule, Arg } from '@libs/frm'
import lib from '@libs/lib'

import AjaxSelect from '@components/ajax-select.vue'

interface Props {
  rules: Rule[]
  arg: Arg
}

const props = defineProps<Props>()
const emit = defineEmits<{
  'update:arg': [arg: Arg]
  'search': []
  'clear': []
}>()

const rules = computed(() => {
  return props.rules.filter(i => i.search)
})

const arg = computed({
  get: () => props.arg,
  set: (newArg: Arg) => {
    emit('update:arg', newArg)
  }
})

const keys = ref<Set<string>>(new Set())

const initKeys = () => {
  keys.value.clear()
  // 先找出有值的字段
  for (const r of rules.value) {
    if (arg.value && arg.value[r.key] !== undefined) {
      keys.value.add(r.key)
    }
  }
  // 如果没有有值的字段，默认显示第一个
  if (keys.value.size === 0 && rules.value.length > 0) {
    keys.value.add(rules.value[0]!.key)
  }
}

const toggle = (key: string) => {
  if (keys.value.has(key)) {
    keys.value.delete(key)
    delete arg.value[key]
  } else {
    keys.value.add(key)
  }
}

const clear = () => {
  arg.value = {}
  initKeys()
  emit('clear')
}

onMounted(initKeys)
</script>

<style scoped>
.searcher {
  margin-left: 5px;
  margin-bottom: 5px;
  width: 100%;
  max-width: 480px;
}
.label {
  font-weight: bold;
}
.ant-row {
  padding: 2px 0;
}
.search-input {
  width: 100%;
  max-width: 480px;
}
@media (max-width: 768px) {
  .search-input {
    max-width: none;
  }
}
.ant-space:not(.auto-hide-l1) {
  display: flex !important;
  flex-wrap: wrap !important;
}
/* 大屏时筛选条件标签不折行 */
.searcher .filter-tags {
  flex-wrap: nowrap !important;
}
@media (max-width: 768px) {
  .searcher .filter-tags {
    flex-wrap: wrap !important;
  }
}
</style>

