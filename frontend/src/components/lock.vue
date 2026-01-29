<template>
  <div :style="title ? { border: '1px solid #C0C0C0', borderRadius: '5px', padding: '5px 10px' } : {}">
    <b v-if="title != ''" style="margin-right: 0.5rem">{{ title }}</b>
    <a-switch v-model:checked="checked" @change="switchLock" class="warning-switch">
      <template #checkedChildren><i class="fa fa-lock lock-icon"></i></template>
      <template #unCheckedChildren><i class="fa fa-lock-open lock-icon"></i></template>
    </a-switch>
    <span v-for="u, i in users" :key="i" class="badge badge-warning">
      <i class="fa fa-lock"></i> {{ u }}
    </span>
  </div>
</template>
<script setup lang="ts">
import { ref, watch } from 'vue'
import lib from '@libs/lib.ts'

interface Props {
  user?: string
  type?: string
  title?: string
  locked?: boolean
  lockers?: string[]
}

const props = withDefaults(defineProps<Props>(), {
  lockers: () => []
})
const checked = ref(props.locked ?? false)
const users = ref([...(props.lockers ?? [])])

// 监听所有 props 变化，更新 checked 和 users
watch(
  () => ({ locked: props.locked, lockers: props.lockers }), // 只监听 locked 和 lockers
  (newProps) => {
    checked.value = newProps.locked
    users.value.splice(0, users.value.length, ...newProps.lockers) // 替换整个数组
  },
  { deep: true } // 深度监听（确保数组变化能被检测到）
)

const switchLock = async () => {
  const ok = await lib.ajax('/base/lock/lock', { name: props.type ?? '' })
  if (! ok) {
    checked.value = !checked.value
    return
  }
  if (checked.value && props.user) {
    if (! users.value.includes(props.user)) {
      users.value.push(props.user)
    }
  } else if (props.user) {
    const index = users.value.indexOf(props.user)
    if (index > -1) {
      users.value.splice(index, 1)
    }
  }
}
</script>

<style scoped>
.warning-switch.ant-switch-checked {
  background-color: #ffb822 !important;
}
.lock-icon {
  padding: 3px;
}
.badge {
  margin-left: 5px;
}
</style>
