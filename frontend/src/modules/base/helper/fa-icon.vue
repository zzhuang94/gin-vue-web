<template>
  <div class="portlet">
    <div class="portlet-body">
      <a-row :gutter="[16, 16]" style="margin-bottom: 8px">
        <a-col :span="8">
          <a-input v-model:value="q" allow-clear placeholder="搜索图标，如 'truck' 或 'calendar'" />
        </a-col>
      </a-row>

      <!-- New Icons Section -->
      <div v-if="displayedNew.length > 0 || filteredNew.length > 0" class="icon-section">
        <h3 class="section-title">
          <span class="badge new-badge">NEW</span>
          Font Awesome 7 新图标 ({{ filteredNew.length }})
        </h3>
        <a-row :gutter="[12, 12]" align="top">
          <a-col v-for="(name, i) in displayedNew" :key="`new-${i}`" :xs="8" :sm="6" :md="4" :lg="3" :xl="2" :xxl="2">
            <div class="icon-tile new-icon">
              <i :class="`fas fa-${name}`"></i>
              <div class="icon-name">{{ name }}</div>
            </div>
          </a-col>
        </a-row>
      </div>

      <!-- Brands Icons Section -->
      <div v-if="displayedBrands.length > 0 || filteredBrands.length > 0" class="icon-section" style="margin-top: 32px">
        <h3 class="section-title">
          <span class="badge brands-badge">BRANDS</span>
          Font Awesome 7 Brands 品牌图标 ({{ filteredBrands.length }})
        </h3>
        <a-row :gutter="[12, 12]" align="top">
          <a-col v-for="(name, i) in displayedBrands" :key="`brands-${i}`" :xs="8" :sm="6" :md="4" :lg="3" :xl="2" :xxl="2">
            <div class="icon-tile brands-icon">
              <i :class="`fab fa-${name}`"></i>
              <div class="icon-name">{{ name }}</div>
            </div>
          </a-col>
        </a-row>
      </div>

      <!-- Old Icons Section -->
      <div v-if="displayedOld.length > 0 || filteredOld.length > 0" class="icon-section" style="margin-top: 32px">
        <h3 class="section-title">
          <span class="badge old-badge">OLD</span>
          Font Awesome 4 旧图标 ({{ filteredOld.length }})
        </h3>
        <a-row :gutter="[12, 12]" align="top">
          <a-col v-for="(name, i) in displayedOld" :key="`old-${i}`" :xs="8" :sm="6" :md="4" :lg="3" :xl="2" :xxl="2">
            <div class="icon-tile old-icon">
              <i :class="`fas fa-${name}`"></i>
              <div class="icon-name">{{ name }}</div>
            </div>
          </a-col>
        </a-row>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch, onBeforeUnmount } from 'vue'
import iconsJson from './fa-icon.json'

const q = ref('')
const oldIcons = ref([])
const newIcons = ref([])
const brandsIcons = ref([])
const displayedNew = ref([])
const displayedOld = ref([])
const displayedBrands = ref([])

const filteredNew = computed(() => {
  const keyword = q.value.trim().toLowerCase()
  const uniq = Array.from(new Set(newIcons.value))
  if (!keyword) return uniq
  return uniq.filter(n => n && n.toLowerCase().includes(keyword))
})

const filteredBrands = computed(() => {
  const keyword = q.value.trim().toLowerCase()
  const uniq = Array.from(new Set(brandsIcons.value))
  if (!keyword) return uniq
  return uniq.filter(n => n && n.toLowerCase().includes(keyword))
})

const filteredOld = computed(() => {
  const keyword = q.value.trim().toLowerCase()
  const uniq = Array.from(new Set(oldIcons.value))
  if (!keyword) return uniq
  return uniq.filter(n => n && n.toLowerCase().includes(keyword))
})

onMounted(() => {
  oldIcons.value = iconsJson.old || []
  newIcons.value = iconsJson.new || []
  brandsIcons.value = iconsJson.brands || []
})

// Progressive rendering for new icons
let rafIdNew = 0
const step = 60 // items per frame
const startProgressiveNew = () => {
  cancelAnimationFrame(rafIdNew)
  displayedNew.value = []
  const src = filteredNew.value
  let idx = 0
  const tick = () => {
    const end = Math.min(idx + step, src.length)
    if (end > idx) {
      displayedNew.value = src.slice(0, end)
      idx = end
    }
    if (idx < src.length) {
      rafIdNew = requestAnimationFrame(tick)
    }
  }
  tick()
}

// Progressive rendering for brands icons
let rafIdBrands = 0
const startProgressiveBrands = () => {
  cancelAnimationFrame(rafIdBrands)
  displayedBrands.value = []
  const src = filteredBrands.value
  let idx = 0
  const tick = () => {
    const end = Math.min(idx + step, src.length)
    if (end > idx) {
      displayedBrands.value = src.slice(0, end)
      idx = end
    }
    if (idx < src.length) {
      rafIdBrands = requestAnimationFrame(tick)
    }
  }
  tick()
}

// Progressive rendering for old icons
let rafIdOld = 0
const startProgressiveOld = () => {
  cancelAnimationFrame(rafIdOld)
  displayedOld.value = []
  const src = filteredOld.value
  let idx = 0
  const tick = () => {
    const end = Math.min(idx + step, src.length)
    if (end > idx) {
      displayedOld.value = src.slice(0, end)
      idx = end
    }
    if (idx < src.length) {
      rafIdOld = requestAnimationFrame(tick)
    }
  }
  tick()
}

watch(filteredNew, startProgressiveNew, { immediate: true })
watch(filteredBrands, startProgressiveBrands, { immediate: true })
watch(filteredOld, startProgressiveOld, { immediate: true })
onBeforeUnmount(() => {
  cancelAnimationFrame(rafIdNew)
  cancelAnimationFrame(rafIdBrands)
  cancelAnimationFrame(rafIdOld)
})
</script>

<style scoped>
.icon-section {
  margin-bottom: 24px;
}

.section-title {
  font-size: 18px;
  font-weight: 600;
  margin-bottom: 16px;
  color: #333;
  display: flex;
  align-items: center;
  gap: 8px;
}

.badge {
  display: inline-block;
  padding: 2px 8px;
  border-radius: 4px;
  font-size: 12px;
  font-weight: 700;
  letter-spacing: 0.5px;
}

.new-badge {
  background: #52c41a;
  color: #fff;
}

.old-badge {
  background: #faad14;
  color: #fff;
}

.brands-badge {
  background: #1890ff;
  color: #fff;
}

.icon-tile {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 110px;
  padding: 8px;
  border: 1px solid #f0f0f0;
  border-radius: 8px;
  background: #fff;
  transition: box-shadow 0.2s ease, transform 0.1s ease;
}

.icon-tile.new-icon {
  border-color: #b7eb8f;
}

.icon-tile.brands-icon {
  border-color: #91d5ff;
}

.icon-tile.old-icon {
  border-color: #ffe58f;
}

.icon-tile:hover {
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.08);
  transform: translateY(-1px);
}

.icon-tile.new-icon:hover {
  box-shadow: 0 2px 10px rgba(82, 196, 26, 0.15);
}

.icon-tile.brands-icon:hover {
  box-shadow: 0 2px 10px rgba(24, 144, 255, 0.15);
}

.icon-tile.old-icon:hover {
  box-shadow: 0 2px 10px rgba(250, 173, 20, 0.15);
}

.icon-tile i {
  font-size: 28px;
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
