<template>
  <a-range-picker
    show-time
    :format="format"
    v-model:value="timeRange"
    :placeholder="['开始时间', '结束时间']"
    @change="handleChange"
  />
</template>

<script setup lang="ts">
import { ref } from 'vue'
import dayjs, { Dayjs } from 'dayjs'

interface Props {
  start?: string
  end?: string
  format?: string
}

const props = withDefaults(defineProps<Props>(), {
  start: '',
  end: '',
  format: 'YYYY-MM-DD HH:mm:00'
})

const emit = defineEmits<{
  'change': [timeStrings: string[]]
}>()

const timeRange = ref<(Dayjs | null)[]>(initTimeRange())

const handleChange = (times: (Dayjs | null)[] | null, timeStrings: string[]) => {
  if (times && times.length === 2) {
    emit('change', timeStrings)
  }
}

function initTimeRange(): (Dayjs | null)[] {
  const ans: (Dayjs | null)[] = []
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
