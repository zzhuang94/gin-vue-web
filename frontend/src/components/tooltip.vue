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

<script setup lang="ts">
interface Props {
  placement?: string
  color?: string
  icon?: boolean
  msg: string
}

withDefaults(defineProps<Props>(), {
  placement: 'right',
  icon: true
})
</script>
