<template>
  <div class="board-page">
    <div class="board-toolbar">
      <a-input
        v-model:value="q"
        allow-clear
        placeholder="搜索机器名称…"
        class="board-search"
      >
        <template #prefix><i class="fas fa-search"></i></template>
      </a-input>
      <a-select
        v-model:value="sortMode"
        class="board-sort"
        :options="sortOptions"
      />
      <span class="stat-total">
        共 <b>{{ stats.total }}</b> 台
        <template v-if="q.trim()">，显示 <b>{{ filtered.length }}</b> 台</template>
      </span>
      <div class="board-stats">
        <span class="stat-chip stat-ok" title="正常">
          <i class="fas fa-circle"></i>正常 {{ stats.ok }}
        </span>
        <span class="stat-chip stat-fault" title="故障">
          <i class="fas fa-circle"></i>故障 {{ stats.fault }}
        </span>
        <span class="stat-chip stat-run" title="生产中">
          <i class="fas fa-circle"></i>生产中 {{ stats.running }}
        </span>
      </div>
      <div class="board-actions">
        <button type="button" class="btn-tiny btn-tiny-primary" @click="loadData" :disabled="loading">
          <i class="fa" :class="loading ? 'fa-spinner fa-spin' : 'fa-refresh'"></i> 刷新
        </button>
        <a class="btn-tiny btn-tiny-default" href="/prod/machine/index">
          <i class="fa fa-list"></i> 列表
        </a>
      </div>
    </div>

    <a-spin :spinning="loading">
      <a-empty v-if="!loading && filtered.length === 0" description="暂无机器" class="board-empty" />
      <a-row v-else :gutter="gutter" class="board-grid">
        <a-col
          v-for="m in displayed"
          :key="m.id"
          :xs="8"
          :sm="6"
          :md="4"
          :lg="3"
          :xl="2"
          class="board-col"
        >
          <div
            class="machine-card"
            :class="cardClass(m)"
            role="button"
            tabindex="0"
            :title="`${m.name}（点击编辑）`"
            @click="editMachine(m)"
            @keydown.enter="editMachine(m)"
          >
            <div class="card-glow" :class="glowClass(m)"></div>

            <div class="card-top">
              <span class="status-led" :class="ledClass(m)"></span>
              <div class="card-title">{{ m.name }}</div>
            </div>

            <div class="card-tags">
              <span class="pill" :class="m.health == 1 ? 'pill-ok' : 'pill-fault'">
                {{ m.health == 1 ? '正常' : '故障' }}
              </span>
              <span class="pill" :class="m.running ? 'pill-run' : 'pill-idle'">
                {{ m.running ? `生产中 ×${m.tickets.length}` : '空闲' }}
              </span>
            </div>

            <div class="card-tickets">
              <template v-if="m.tickets.length">
                <a-tooltip v-for="t in m.tickets" :key="t.id" :title="t.label">
                  <span class="ticket-tag" @click.stop>#{{ t.id }} {{ t.label }}</span>
                </a-tooltip>
              </template>
              <span v-else class="ticket-empty">无在产工单</span>
            </div>
          </div>
        </a-col>
      </a-row>
    </a-spin>

    <component :is="mc" v-bind="mp" @submit="onEditSubmit" />
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch, onMounted, onBeforeUnmount } from 'vue'
import lib from '@libs/lib'
import { useModal } from '@/libs/modal'

interface TicketBrief {
  id: number
  label: string
}

interface MachineItem {
  id: number
  name: string
  remark: string
  health: number
  running: boolean
  tickets: TicketBrief[]
}

type SortMode = 'default' | 'fault' | 'idle' | 'running'

const sortOptions = [
  { value: 'default', label: '默认排序' },
  { value: 'fault', label: '故障设备靠前' },
  { value: 'idle', label: '空闲设备靠前' },
  { value: 'running', label: '生产设备靠前' },
]

const q = ref('')
const sortMode = ref<SortMode>('default')
const loading = ref(false)
const machines = ref<MachineItem[]>([])
const displayed = ref<MachineItem[]>([])
const isMobile = ref(false)

const { mc, mp, loadModal } = useModal()

const gutter = computed(() => (isMobile.value ? [8, 8] : [12, 12]))

const filtered = computed(() => {
  const keyword = q.value.trim().toLowerCase()
  if (!keyword) return machines.value
  return machines.value.filter(m => m.name.toLowerCase().includes(keyword))
})

const sorted = computed(() => {
  const list = [...filtered.value]
  const byName = (a: MachineItem, b: MachineItem) =>
    a.name.localeCompare(b.name, 'zh-CN')

  switch (sortMode.value) {
    case 'fault':
      return list.sort((a, b) => {
        const cmp = (a.health === 0 ? 0 : 1) - (b.health === 0 ? 0 : 1)
        return cmp !== 0 ? cmp : byName(a, b)
      })
    case 'idle':
      return list.sort((a, b) => {
        const cmp = (a.running ? 1 : 0) - (b.running ? 1 : 0)
        return cmp !== 0 ? cmp : byName(a, b)
      })
    case 'running':
      return list.sort((a, b) => {
        const cmp = (b.running ? 1 : 0) - (a.running ? 1 : 0)
        return cmp !== 0 ? cmp : byName(a, b)
      })
    default:
      return list
  }
})

const stats = computed(() => {
  const list = machines.value
  return {
    total: list.length,
    ok: list.filter(m => m.health == 1).length,
    fault: list.filter(m => m.health == 0).length,
    running: list.filter(m => m.running).length,
  }
})

let rafId = 0
const renderStep = 60

function startProgressive() {
  cancelAnimationFrame(rafId)
  displayed.value = []
  const src = sorted.value
  let idx = 0
  const tick = () => {
    const end = Math.min(idx + renderStep, src.length)
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

function updateMobile() {
  isMobile.value = window.innerWidth < 576
}

function cardClass(m: MachineItem): string {
  if (m.health == 0) return m.running ? 'card-fault-run' : 'card-fault'
  return m.running ? 'card-ok-run' : 'card-ok'
}

function glowClass(m: MachineItem): string {
  if (m.health == 0) return 'glow-fault'
  return m.running ? 'glow-run' : 'glow-ok'
}

function ledClass(m: MachineItem): string {
  if (m.health == 0) return 'led-fault'
  return m.running ? 'led-run' : 'led-ok'
}

function editMachine(m: MachineItem) {
  loadModal(`edit?id=${m.id}`)
}

async function onEditSubmit() {
  await loadData()
}

async function loadData() {
  loading.value = true
  const r = await lib.curl('board-data')
  if (r?.code === 1) {
    machines.value = (r.data || []).map((m: MachineItem) => ({
      ...m,
      tickets: m.tickets || [],
    }))
  }
  loading.value = false
}

watch(sorted, startProgressive)

onMounted(() => {
  updateMobile()
  window.addEventListener('resize', updateMobile)
  loadData()
})

onBeforeUnmount(() => {
  window.removeEventListener('resize', updateMobile)
  cancelAnimationFrame(rafId)
})
</script>

<style scoped>
.board-page {
  padding: 2px 0 12px;
}

/* ---- toolbar (light) ---- */
.board-toolbar {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  gap: 10px 14px;
  margin-bottom: 14px;
  padding: 12px 14px;
  border-radius: 8px;
  background: #f7f8fa;
  border: 1px solid #e8e8e8;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.04);
}

.board-search {
  width: 200px;
  flex-shrink: 0;
}

.board-sort {
  width: 148px;
  flex-shrink: 0;
}

.stat-total {
  font-size: 14px;
  color: #595959;
  white-space: nowrap;
}

.stat-total b {
  color: #1677ff;
  font-size: 16px;
  margin: 0 2px;
}

.board-stats {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  flex: 1;
  min-width: 0;
}

.stat-chip {
  font-size: 13px;
  padding: 4px 10px;
  border-radius: 14px;
  background: #fff;
  color: #595959;
  border: 1px solid #e8e8e8;
  white-space: nowrap;
}

.stat-chip i {
  font-size: 7px;
  margin-right: 5px;
  vertical-align: middle;
}

.stat-ok i { color: #52c41a; }
.stat-fault i { color: #ff4d4f; }
.stat-run i { color: #faad14; }

.board-actions {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-left: auto;
}

.btn-tiny {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 4px;
  height: 28px;
  padding: 0 12px;
  font-size: 13px;
  line-height: 1;
  border-radius: 4px;
  border: 1px solid #d9d9d9;
  cursor: pointer;
  text-decoration: none;
  white-space: nowrap;
  background: #fff;
  color: #595959;
  transition: border-color 0.15s, color 0.15s;
}

.btn-tiny:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.btn-tiny-primary {
  color: #1677ff;
  border-color: #91caff;
  background: #e6f4ff;
}

.btn-tiny-default:hover {
  color: #1677ff;
  border-color: #1677ff;
}

.board-empty {
  margin-top: 40px;
}

/* ---- grid & card ---- */
.board-col {
  display: flex;
}

.machine-card {
  position: relative;
  width: 100%;
  height: 118px;
  display: flex;
  flex-direction: column;
  padding: 8px 10px 6px;
  border-radius: 8px;
  background: #fff;
  border: 1px solid #e8e8e8;
  overflow: hidden;
  cursor: pointer;
  transition: box-shadow 0.15s ease, border-color 0.15s ease;
  user-select: none;
}

.machine-card:focus-visible {
  outline: 2px solid #1677ff;
  outline-offset: 1px;
}

@media (hover: hover) {
  .machine-card:hover {
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
    border-color: #91caff;
  }
}

.card-glow {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 3px;
}

.glow-ok { background: linear-gradient(90deg, #52c41a, #95de64); }
.glow-run { background: linear-gradient(90deg, #faad14, #ffc53d); }
.glow-fault { background: linear-gradient(90deg, #ff4d4f, #ff7875); }

.card-ok { border-color: #d9f7be; }
.card-ok-run { border-color: #ffe58f; background: linear-gradient(180deg, #fffbe6 0%, #fff 40%); }
.card-fault { border-color: #ffccc7; background: linear-gradient(180deg, #fff1f0 0%, #fff 40%); }
.card-fault-run { border-color: #ffa39e; background: linear-gradient(180deg, #fff2f0 0%, #fffbe6 40%); }

.card-top {
  position: relative;
  flex-shrink: 0;
  padding-right: 10px;
}

.status-led {
  position: absolute;
  top: 4px;
  right: 0;
  width: 8px;
  height: 8px;
  border-radius: 50%;
  box-shadow: 0 0 5px currentColor;
}

.led-ok { background: #52c41a; color: #52c41a; }
.led-run { background: #faad14; color: #faad14; animation: pulse 1.5s ease-in-out infinite; }
.led-fault { background: #ff4d4f; color: #ff4d4f; }

@keyframes pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.45; }
}

.card-title {
  font-size: 14px;
  font-weight: 600;
  color: #262626;
  line-height: 1.35;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.card-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 4px;
  flex-shrink: 0;
  margin: 6px 0 4px;
}

.pill {
  font-size: 12px;
  line-height: 1.3;
  padding: 2px 8px;
  border-radius: 10px;
  font-weight: 500;
}

.pill-ok { background: #f6ffed; color: #389e0d; border: 1px solid #b7eb8f; }
.pill-fault { background: #fff1f0; color: #cf1322; border: 1px solid #ffa39e; }
.pill-run { background: #fff7e6; color: #d46b08; border: 1px solid #ffd591; }
.pill-idle { background: #f5f5f5; color: #8c8c8c; border: 1px solid #d9d9d9; }

.card-tickets {
  flex: 1;
  min-height: 0;
  overflow-y: auto;
  overflow-x: hidden;
  -webkit-overflow-scrolling: touch;
  display: flex;
  flex-direction: column;
  gap: 3px;
}

.card-tickets::-webkit-scrollbar {
  width: 3px;
}

.card-tickets::-webkit-scrollbar-thumb {
  background: #d9d9d9;
  border-radius: 2px;
}

.ticket-tag {
  display: block;
  font-size: 12px;
  line-height: 1.35;
  padding: 2px 6px;
  border-radius: 4px;
  background: rgba(250, 173, 20, 0.12);
  color: #ad6800;
  border: 1px solid rgba(250, 173, 20, 0.35);
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  cursor: default;
}

.ticket-empty {
  font-size: 12px;
  color: #bfbfbf;
  line-height: 1.4;
}

/* ---- mobile ---- */
@media (max-width: 575.98px) {
  .board-page {
    margin: 0 -4px;
  }

  .board-toolbar {
    flex-direction: column;
    align-items: stretch;
    gap: 8px;
    padding: 10px;
  }

  .board-search,
  .board-sort {
    width: 100%;
  }

  .stat-total {
    text-align: center;
  }

  .board-stats {
    justify-content: space-between;
  }

  .stat-chip {
    flex: 1;
    text-align: center;
    font-size: 12px;
    padding: 4px 6px;
  }

  .board-actions {
    margin-left: 0;
    justify-content: stretch;
  }

  .btn-tiny {
    flex: 1;
  }

  .machine-card {
    height: 112px;
    padding: 7px 8px 5px;
  }

  .card-title {
    font-size: 13px;
  }
}

@media (max-width: 359.98px) {
  .board-grid :deep(.ant-col) {
    flex: 0 0 50%;
    max-width: 50%;
  }

  .machine-card {
    height: 108px;
  }
}
</style>
