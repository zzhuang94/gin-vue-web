<template>
  <component :is="hasLayout ? LayoutShell : 'div'" v-bind="layout">
    <div v-if="loading" class="portlet">
      <div class="portlet-body">
        <a-skeleton />
      </div>
    </div>
    <component v-else-if="pc" :is="pc" :key="pageKey" v-bind="data" />
  </component>
</template>

<script setup lang="ts">
import { ref, watch, computed, defineAsyncComponent, onMounted } from 'vue'
import type { Component } from 'vue'
import { useRoute } from 'vue-router'
import { isEmpty } from 'lodash'
import lib from '@libs/lib'

const uri = ref(window.location.pathname + window.location.search)
const page = ref('')
const data = ref<Record<string, any>>({})
const layout = ref<Record<string, any>>({})

const loading = ref(true)
const pc = ref<Component | null>(null)
const LayoutShell = defineAsyncComponent(() => import('./templates/layout.vue'))
const hasLayout = computed(() => !isEmpty(layout.value))

const pageKey = computed(() => `${uri.value}-${page.value}`)

const loadResources = async () => {
  loading.value = true
  const r = await lib.curl(uri.value)
  if (r.code === -1) {
    page.value = 'templates/noauth'
  } else if (r != null) {
    page.value = r.page
    data.value = r.data || {}
    layout.value = r.layout || {}
  } else {
    layout.value = {}
  }

  pc.value = lib.loadComponent(page.value)

  loading.value = false
}

const route = useRoute()
watch(() => route.fullPath, (newUri: string) => {
  const newPath = newUri.split('?')[0]
  const oldPath = uri.value.split('?')[0]
  if (newUri == '/' || newPath === oldPath && newPath !== '/') {
    return
  }
  uri.value = newUri
  loadResources()
}, { immediate: false })

onMounted(loadResources)
</script>
