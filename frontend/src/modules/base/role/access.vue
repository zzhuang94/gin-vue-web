<template>
    <a-modal width="40%" cancel-text="取消" ok-text="保存" :maskClosable="false" @ok="submit" v-model:open="open" :confirm-loading="submitting">
      <template #title>
        <i class="fa fa-edit"></i> <code>{{ role.name }}</code> 特权管理
      </template>
      <hr />
      <a-alert type="success" closable style="margin-bottom: 10px">
        <template #message><i class="fa fa-bell"></i> 绿色的Action所有用户都有权限，无需分配</template>
      </a-alert>
      <Tree checkable :data v-model:checkedKeys="ids" />
      <hr />
    </a-modal>
  </template>
  
  
  <script setup>
  import { ref, reactive, shallowRef, onMounted } from 'vue';
  import lib from '@libs/lib.ts'
  import Tree from '@components/tree.vue'
  
  const props = defineProps(['role', 'data', 'ids'])
  const data = ref(props.data)
  const ids = ref([...props.ids])
  const submitting = ref(false)
  const open = ref(true)
  
  const submit = async () => {
    submitting.value = true
    const params = { id: props.role.id, ids: ids.value }
    const ok = await lib.ajax('access-save', params)
    if (ok) {
      open.value = false
    }
    submitting.value = false
  }
  </script>
  