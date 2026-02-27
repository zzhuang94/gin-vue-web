<template>
  <a-modal cancel-text="取消" ok-text="保存" :maskClosable="false" @ok="submit" v-model:open="open"
    :confirm-loading="submitting">
    <template #title>
      <i class="fa fa-plus"></i> 批量新增
    </template>
    <hr />
    <a-textarea v-model:value="names" placeholder="请输入名称，每行一个，系统会自动处理空行和重复名称" :rows="10" />
    <hr />
  </a-modal>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import lib from '@libs/lib.ts'
import strlib from '@libs/strlib.ts'

const emit = defineEmits<{'submit': []}>()
const open = ref(true)
const submitting = ref(false)
const names = ref<string>('')

const submit = async () => {
  submitting.value = true
  const namesList = strlib.lineExplode(names.value)
  const ok = await lib.ajax('batch-add', { names: namesList })
  if (ok) {
    emit('submit')
    open.value = false
  }
  submitting.value = false
}
</script>