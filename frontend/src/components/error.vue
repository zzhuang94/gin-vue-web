<template>
  <a-modal :width="width" :maskClosable="false" v-model:open="open" :footer="null" v-if="isModal">
    <template #title>
      <i class="fa fa-warning"></i> 出错啦~
    </template>
    <hr />
    <a-result status="500" :title="msg" />
  </a-modal>

  <a-result v-else status="500" :title="msg" />
</template>

<script setup lang="ts">
import { ref } from 'vue';

interface Props {
  width?: string
  isModal?: boolean
  msg?: string
}

const open = ref(true)
withDefaults(defineProps<Props>(), {
  width: '50%'
})
</script>
