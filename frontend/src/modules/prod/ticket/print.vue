<template>
  <div class="print-source" aria-hidden="true">
    <article id="ticket-print-sheet" ref="sheet" class="ticket-sheet">
      <header class="ticket-heading">
        <div>
          <h1>生产工单</h1>
          <div class="ticket-meta">
            <span class="ticket-number">工单编号 #{{ ticket.id }}</span>
            <span v-if="statusRule" class="ticket-status">{{ fieldValue(statusRule) }}</span>
          </div>
        </div>
        <figure>
          <QRCode :value="qrUrl" type="svg" :size="116" :bordered="false" color="#000" bg-color="#fff" />
          <figcaption>扫码查看工单</figcaption>
        </figure>
      </header>

      <section>
        <h2>生产要求</h2>
        <dl v-for="group in fieldGroups" :key="group.className" class="ticket-fields" :class="group.className">
          <div v-for="rule in group.rules" :key="rule.key" class="ticket-field">
            <dt>{{ rule.name }}</dt>
            <dd>{{ fieldValue(rule) }}</dd>
          </div>
        </dl>
      </section>

      <section>
        <h2>流程记录</h2>
        <table class="ticket-workflow">
          <thead>
            <tr>
              <th scope="col">环节</th>
              <th scope="col">操作人</th>
              <th scope="col">操作时间</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="step in workflow" :key="step.label">
              <th scope="row">{{ step.label }}</th>
              <td>{{ step.operator }}</td>
              <td>{{ step.time }}</td>
            </tr>
          </tbody>
        </table>
      </section>

      <section>
        <h2>设备与备注</h2>
        <dl class="ticket-details">
          <div v-for="rule in detailRules" :key="rule.key">
            <dt>{{ rule.name }}</dt>
            <dd>{{ fieldValue(rule) }}</dd>
          </div>
        </dl>
      </section>
    </article>
    <button ref="printButton" type="button" hidden v-print="printOptions" />
  </div>
</template>

<script setup lang="ts">
import { computed, nextTick, onMounted, ref } from 'vue'
import { QRCode } from 'ant-design-vue'
import type { Data, Rule } from '@libs/frm'

const props = defineProps<{
  ticket: Data
  rules: Rule[]
  host: string
  closingOperation?: { operator_name: string; operation_time: string } | null
}>()
const sheet = ref<HTMLElement | null>(null)
const printButton = ref<HTMLButtonElement | null>(null)
const detailKeys = ['machine_plan', 'machine_list', 'remark']
const metricKeys = ['quantity', 'lead_time', 'progress', 'rate']
const statusRule = computed(() => props.rules.find(rule => rule.key === 'status'))
const fieldGroups = computed(() => [
  {
    className: 'ticket-product',
    rules: props.rules.filter(rule =>
      rule.key !== 'status' && !metricKeys.includes(rule.key) && !detailKeys.includes(rule.key)),
  },
  { className: 'ticket-metrics', rules: props.rules.filter(rule => metricKeys.includes(rule.key)) },
])
const detailRules = computed(() => props.rules.filter(rule => detailKeys.includes(rule.key)))

const workflowSteps = [
  { key: 'apply', label: '提交需求' },
  { key: 'plan', label: '编排机器' },
  { key: 'prepare', label: '准备机器' },
  { key: 'run', label: '开始生产' },
]
const closingLabels: Record<string, string> = { FINISHED: '完成生产', STOPPED: '终止生产' }
const workflow = computed(() => {
  const steps = workflowSteps.map(({ key, label }) => ({
    label,
    operator: displayText(props.ticket[`${key}_user`]),
    time: displayText(props.ticket[`${key}_time`]),
  }))
  const closingLabel = closingLabels[props.ticket.status]
  if (closingLabel) {
    steps.push({
      label: closingLabel,
      operator: displayText(props.closingOperation?.operator_name),
      time: displayText(props.closingOperation?.operation_time),
    })
  }
  return steps
})

const ticketUrl = new URL('/prod/ticket/index', scanOrigin(props.host))
ticketUrl.searchParams.set('id', props.ticket.id)
const qrUrl = ticketUrl.href

const printOptions = {
  id: '#ticket-print-sheet',
  popTitle: '生产工单',
  extraHead: '<meta charset="utf-8"><style>@page { size: A4 portrait; margin: 12mm; } body { margin: 0; }</style>',
}

function scanOrigin(host: string): string {
  const configured = host.trim()
  if (configured) {
    try {
      const base = new URL(configured.includes('://') ? configured : `${window.location.protocol}//${configured}`)
      if (['http:', 'https:'].includes(base.protocol) && !['localhost', '127.0.0.1', '[::1]'].includes(base.hostname)) {
        return base.origin
      }
    } catch {
      // 配置地址无效时，使用当前页面的地址。
    }
  }
  return window.location.origin
}

function fieldValue(rule: Rule): string {
  const value = String(props.ticket[rule.key] ?? rule.default ?? '')
  if (!value.trim()) return '—'
  if (rule.split_sep) {
    return value.split(rule.split_sep).map(part => part.trim()).join('、')
  }
  const label = rule.limit_map?.[value]?.label ?? value
  return `${rule.prefix ?? ''}${label}${rule.suffix ?? ''}`
}

function displayText(value: unknown): string {
  return String(value ?? '').trim() || '—'
}

onMounted(async () => {
  await nextTick()
  const paper = sheet.value
  if (!paper || !printButton.value) return
  const { width, height } = paper.getBoundingClientRect()
  // A4 扣除页边距后宽 186mm、高 273mm，留 1mm 余量避免分页舍入。
  const availableHeight = width * 272 / 186
  paper.style.zoom = String(Math.min(1, availableHeight / height))
  printButton.value.click()
})
</script>

<style scoped>
.print-source {
  position: fixed;
  top: 0;
  left: -10000px;
}

.ticket-sheet {
  box-sizing: border-box;
  width: 186mm;
  margin: 0 auto;
  color: #172b3a;
  font: 11pt/1.5 "PingFang SC", "Microsoft YaHei", "Noto Sans CJK SC", sans-serif;
  overflow-wrap: anywhere;
}

.ticket-heading {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 6mm;
  padding-bottom: 3mm;
  border-bottom: 2pt solid #172b3a;
}

.ticket-meta {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  gap: 2mm 4mm;
  margin-top: 2mm;
}

.ticket-number {
  font-weight: 600;
  font-variant-numeric: tabular-nums;
}

.ticket-status {
  padding: 0.4mm 2.5mm;
  border: 1px solid currentColor;
  border-radius: 3mm;
  color: #17645d;
  font-size: 9pt;
  font-weight: 600;
}

.ticket-heading h1 {
  margin: 0;
  font-size: 26pt;
  font-weight: 700;
  letter-spacing: 1mm;
}

.ticket-heading figure {
  flex: none;
  margin: 0;
  text-align: center;
}

.ticket-heading figcaption {
  font-size: 9pt;
}

.ticket-sheet h2 {
  margin: 5mm 0 2mm;
  padding-left: 2mm;
  border-left: 2pt solid #17645d;
  font-size: 12pt;
  font-weight: 700;
}

.ticket-fields {
  display: grid;
  gap: 0 5mm;
  margin: 0;
}

.ticket-product {
  grid-template-columns: repeat(3, minmax(0, 1fr));
}

.ticket-metrics {
  grid-template-columns: repeat(4, minmax(0, 1fr));
}

.ticket-field {
  padding: 2mm 0;
  border-bottom: 1px solid #bdc8cf;
}

.ticket-field dt {
  color: #455a6a;
  font-size: 9pt;
}

.ticket-field dd {
  margin: 0.5mm 0 0;
  font-weight: 600;
}

.ticket-metrics dd {
  font-size: 13pt;
  font-variant-numeric: tabular-nums;
}

.ticket-workflow {
  width: 100%;
  border-collapse: collapse;
  text-align: left;
}

.ticket-workflow th,
.ticket-workflow td {
  padding: 1.5mm 2.5mm;
  border-bottom: 1px solid #bdc8cf;
}

.ticket-workflow thead th {
  border-bottom-color: #455a6a;
  color: #455a6a;
  font-size: 9pt;
}

.ticket-workflow th:first-child {
  width: 26%;
}

.ticket-workflow th:last-child {
  width: 40%;
}

.ticket-workflow td:last-child {
  font-variant-numeric: tabular-nums;
}

.ticket-details {
  margin: 0;
}

.ticket-details > div {
  display: grid;
  grid-template-columns: 24mm minmax(0, 1fr);
  gap: 3mm;
  padding: 2mm 0;
  border-bottom: 1px solid #bdc8cf;
}

.ticket-details dt {
  font-weight: 600;
}

.ticket-details dd {
  margin: 0;
  white-space: pre-wrap;
}
</style>
