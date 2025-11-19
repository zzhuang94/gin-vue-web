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

<script setup>
import { ref } from 'vue';
const open = ref(true)
const props = defineProps({
  width: { type: String, default: '50%' },
  isModal: Boolean,
  msg: String,
})
</script>
