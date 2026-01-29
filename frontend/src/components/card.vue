<template>
  <div :class="['card', `card--${color}`]">
    <div class="card__header">
        <div class="card__header-content">
          <span class="card__icon" v-if="icon">
            <i :class="`fa fa-${icon}`"></i>
          </span>
          <span class="card__title">{{ title }}</span>

          <div style="margin-left: auto">
            <slot name="more-op"></slot>
          </div>
        </div>
    </div>
    <div class="card__body">
      <slot></slot>
    </div>
  </div>
</template>

<script setup lang="ts">
interface Props {
  title?: string
  icon?: string
  color?: 'default' | 'brand' | 'metal' | 'accent' | 'primary' | 'success' | 'info' | 'warning' | 'danger' | 'focus'
}

withDefaults(defineProps<Props>(), {
  title: '卡片',
  icon: 'table',
  color: 'info'
})
</script>

<style scoped>
.card {
  border-radius: 4px;
  box-shadow: 0 1px 15px 1px rgba(69, 65, 78, 0.1);
  margin-bottom: 10px;
  background: #fff;
  overflow: hidden;
}

.card__header {
  padding: 6px 10px;
  display: flex;
  align-items: center;
  color: #fff;
}

.card__header-content {
  display: flex;
  align-items: center;
  width: 100%;
}

.card__icon {
  width: 2rem;
  font-size: 1.2rem;
}

.card__title {
  margin: 0;
  font-size: 1.15rem;
  font-weight: bold;
}

.card__body {
  padding: 10px;
  background: #fff;
}

/* 颜色主题 - 与Metronic保持一致 */
.card--default .card__header { background: #f4f5f8; color: #575962; }
.card--brand .card__header { background: #716aca; }
.card--metal .card__header { background: #c4c5d6; color: #575962; }
.card--accent .card__header { background: #00c5dc; }
.card--primary .card__header { background: #5867dd; }
.card--success .card__header { background: #1dc9b7; }
.card--info .card__header { background: #36a3f7; }
.card--warning .card__header { background: #ffb822; color: #111; }
.card--danger .card__header { background: #fd397a; }
.card--focus .card__header { background: #9816f4; }
</style>
