<template>
  <a-modal cancel-text="取消" ok-text="保存" :maskClosable="false" @ok="submit" v-model:open="open"
    :confirm-loading="submitting">
    <template #title>
      <i class="fa fa-download"></i> 良品入库
    </template>
    <hr />
    <table class="table table-hover" style="width: 100%; border: 0;">
      <tbody>
        <tr>
          <td>类型</td>
          <td>{{ store?.category }}</td>
        </tr>
        <tr>
          <td>材质</td>
          <td>{{ store?.material }}</td>
        </tr>
        <tr>
          <td>颜色</td>
          <td>{{ store?.color }}</td>
        </tr>
        <tr>
          <td>当前库存</td>
          <td>{{ store?.goods }}</td>
        </tr>
        <tr>
          <td>产品来源</td>
          <td>
            <a-radio-group v-model:value="src">
              <a-radio value="gen">新生产</a-radio>
              <a-radio value="oth">退货/代工/其它</a-radio>
            </a-radio-group>
          </td>
        </tr>
        <tr>
          <td>本次入库</td>
          <td><a-input-number v-model:value="count" /></td>
        </tr>

        <tr>
          <td>备注</td>
          <td><a-input v-model:value="remark" /></td>
        </tr>
      </tbody>
    </table>
  </a-modal>
</template>


<script setup lang="ts">
import { ref } from 'vue';
import lib from '@libs/lib.ts'

interface Store {
  id: string | number
  category: string
  material: string
  color: string
  goods: number
}

interface Props {
  store: Store
}

const props = defineProps<Props>()
const emit = defineEmits<{'submit': []}>()
const submitting = ref(false)
const open = ref(true)
const count = ref(0)
const remark = ref('')
const src = ref('gen')

const submit = async () => {
  submitting.value = true
  const ok = await lib.ajax('/prod/store/op', { 
    op: 'PLUS',
    id: props.store.id, 
    count: count.value, 
    remark: remark.value, 
    src: src.value,
  })
  if (ok) {
    open.value = false
    emit('submit')
  }
  submitting.value = false
}
</script>

<style scoped>
tr td:last-child {
  font-weight: bold;
}
</style>