<template>
  <div class="portlet">
    <div class="portlet-body">
      <a-row :gutter="[16, 16]" style="margin-bottom: 8px">
        <a-col :span="8">
          <a-input v-model:value="q" allow-clear placeholder="搜索图标，如 'truck' 或 'calendar'" />
        </a-col>
      </a-row>
      <a-row :gutter="[12, 12]" align="top">
        <a-col v-for="(name, i) in displayed" :key="i" :xs="8" :sm="6" :md="4" :lg="3" :xl="2" :xxl="2">
          <div class="icon-tile">
            <i :class="`fas fa-${name}`"></i>
            <div class="icon-name">{{ name }}</div>
          </div>
        </a-col>
      </a-row>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch, onBeforeUnmount } from 'vue'
import iconsJson from './fa-icon.json'

const q = ref('')
const raw = ref([])
const displayed = ref([])

const filtered = computed(() => {
  const keyword = q.value.trim().toLowerCase()
  const uniq = Array.from(new Set(raw.value))
  if (!keyword) return uniq
  return uniq.filter(n => n && n.toLowerCase().includes(keyword))
})

onMounted(() => {
  raw.value = iconsJson
  return
})

// Progressive rendering
let rafId = 0
const step = 60 // items per frame
const startProgressive = () => {
  cancelAnimationFrame(rafId)
  displayed.value = []
  const src = filtered.value
  let idx = 0
  const tick = () => {
    const end = Math.min(idx + step, src.length)
    if (end > idx) {
      displayed.value = src.slice(0, end)
      idx = end
    }
    if (idx < src.length) {
      rafId = requestAnimationFrame(tick)
    }
  }
  tick()
}

watch(filtered, startProgressive, { immediate: true })
onBeforeUnmount(() => cancelAnimationFrame(rafId))
</script>

<style scoped>
.icon-tile {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 110px; /* 正方形更大一些 */
  padding: 8px;
  border: 1px solid #f0f0f0;
  border-radius: 8px;
  background: #fff;
  transition: box-shadow 0.2s ease, transform 0.1s ease;
}
.icon-tile:hover {
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.08);
  transform: translateY(-1px);
}
.icon-tile i {
  font-size: 28px; /* 图标稍大 */
  line-height: 1;
  margin-bottom: 8px;
}
.icon-name {
  text-align: center;
  font-size: 12px;
  color: #666;
  word-break: break-all;
}
</style>
