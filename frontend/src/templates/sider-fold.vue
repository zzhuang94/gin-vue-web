<template>
  <div>
    <template v-for="l, i in l2" :key="i">
      <a-popover v-if="! isEmpty(l.subs)" placement="rightTop" :overlayClassName="`custom-popover ${env}`" style="padding: 10px">
        <template #title>
          <span style="font-size: 1.1rem; font-weight: bolder; color: white">{{ l.name }}</span>
        </template>
        <template #content>
          <div v-for="s, si in l.subs" :key="si" style="margin: 1rem 0; font-size: 0.95rem;">
            <router-link :to="s.path">
              <span class="a" :class="s.active ? 'active' : ''">
                <b style="margin-right: 0.5rem">·</b>{{ s.name }}
              </span>
            </router-link>
          </div>
        </template>
        <div style="padding-top: 0">
          <i :class="['fas fa-' + l.icon, l.active ? 'active' : '']"></i>
        </div>
      </a-popover>
    </template>
  </div>
</template>

<script setup lang="ts">
import { isEmpty } from 'lodash'

interface SubItem {
  path: string
  name: string
  active?: boolean
}

interface L2Item {
  name: string
  icon: string
  path: string
  active?: boolean
  subs?: SubItem[]
}

interface Props {
  l2?: L2Item[]
  env?: string
}

defineProps<Props>()
</script>

<style scoped>
i {
  margin: 0.7rem 0;
  font-size: 1.2rem;
}
i.active {
  color: yellow;
}
i:hover {
  color: gray;
}
.a {
  color: white;
  font-weight: bold;
}
.a.active {
  color: yellow;
}
.a:hover {
  color: gray;
}

@media (max-width: 1000px) {
  i {
    margin-top: 0.2rem;
    margin-bottom: 0.5rem;
    font-size: 1.3rem;
  }
  i.active {
    color: yellow;
  }
}
</style>

<style>
.custom-popover .ant-popover-inner {
  position: relative;
  top: -6px;
  margin-left: -1.15rem !important;
}
.custom-popover.prod .ant-popover-inner {
  background: black;
}
.custom-popover.dev .ant-popover-inner {
  background: green;
}
.custom-popover.test .ant-popover-inner {
  background: #1a237e;
}
.custom-popover .ant-popover-arrow {
  display: none !important;
}
</style>
