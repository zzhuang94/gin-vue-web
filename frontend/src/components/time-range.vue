<template>
  <a-range-picker
    show-time
    :format="format"
    v-model:value="timeRange"
    :placeholder="['开始时间', '结束时间']"
    @change="handleChange"
  />
</template>

<script setup>
import { ref, watch } from 'vue'
import dayjs from 'dayjs'

const emit = defineEmits(['change'])
const props = defineProps({
  start: {
    type: String,
    default: '',
  },
  end: {
    type: String,
    default: '',
  },
  format: {
    type: String,
    default: 'YYYY-MM-DD HH:mm:00',
  },
})

const timeRange = ref(initTimeRange())

const handleChange = (times, timeStrings) => {
  if (times && times.length === 2) {
    emit('change', timeStrings)
  }
}

function initTimeRange() {
  const ans = []
  if (dayjs(props.start, props.format, true).isValid()) {
    ans.push(dayjs(props.start, props.format, true))
  } else {
    ans.push(null)
  }
  if (dayjs(props.end, props.format, true).isValid()) {
    ans.push(dayjs(props.end, props.format, true))
  } else {
    ans.push(null)
  }
  return ans
}
</script>
