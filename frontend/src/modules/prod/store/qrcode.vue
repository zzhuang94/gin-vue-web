<template>
  <a-modal :footer="false" v-model:open="open">
    <template #title>
      <i class="fa fa-qrcode"></i> 查看二维码
    </template>
    <hr />
    <div id="print-content" style="text-align: center; margin-bottom: 1rem;">
      <a-space style="width: 100%; justify-content: center; margin-bottom: 1rem;" size="large">
        <b>{{ props.store.category }}</b>
        <b>{{ props.store.material }}</b>
        <b>{{ props.store.color }}</b>
      </a-space>
      <a-qrcode :value="props.url" :size="300" style="margin: 0 auto;" />
    </div>
    <div style="text-align: center;">
      <button class="btn btn-lg btn-primary" v-print="'#print-content'">
        <i class="fa fa-print"></i> 打印二维码
      </button>
    </div>
  </a-modal>
</template>

<script setup lang="ts">
import { ref } from 'vue';
interface Store {
  category: string
  material: string
  color: string
}

interface Props {
  store: Store
  url: string
}

const props = defineProps<Props>()
const open = ref(true)

</script>

<style scoped>
b {
  font-size: 1.5rem;
}
</style>