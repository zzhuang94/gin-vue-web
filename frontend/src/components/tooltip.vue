<template>
  <span style="cursor: pointer">
    <a-tooltip :placement="placement" :color="color" :overlayStyle="{ maxWidth: '800px', width: 'auto' }">
      <template #title>
        <span v-html="msg"></span>
      </template>
      <slot />
      <span v-if="icon">&nbsp;<i v-if="icon" class="fa fa-info-circle text-warning"></i></span>
    </a-tooltip>
  </span>
</template>

<script setup>
const props = defineProps({
  placement: {
    type: String,
    default: 'right', // 默认位置
  },
  color: {
    type: String,
    default: undefined, // 默认不指定颜色
  },
  icon: {
    type: Boolean,
    default: true
  },
  msg: {
    type: String,
    required: true,
    default: '',
  },
})
</script>
