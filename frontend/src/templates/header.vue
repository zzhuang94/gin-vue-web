<template>
  <div :class="env">
    <button
      type="button"
      class="hamburger"
      aria-label="打开菜单"
      @click="openDrawer"
    >
      <i class="fa fa-bars"></i>
    </button>
    <a-space :size="30" class="l1 space">
      <template v-for="l, i in l1" :key="i">
        <router-link v-if="l.path" :to="l.path">
          <span class="a" :class="l.active ? 'active' : ''">
            <i :class="'fas fa-' + l.icon" class="l1-icon"></i>
            <span class="auto-hide-l1">{{ l.name }}</span>
          </span>
        </router-link>
      </template>
    </a-space>
    <a-space :size="20" class="pull-right space header-right">
      <button
        v-if="!isSmallScreen && searchCollapsed"
        type="button"
        class="search-icon-btn"
        aria-label="搜索"
        @click="searchCollapsed = false"
      >
        <i class="fa fa-search"></i>
      </button>
      <a-input-search
        v-else-if="!isSmallScreen"
        v-model:value="ip"
        :style="{ width: searchWidth, transition: 'width 0.3s ease' }"
        class="search-input"
        placeholder="全局搜索"
        enter-button
        @search="onSearch"
        @focus="() => { searchWidth = '300px' }"
        @blur="() => { searchWidth = '200px'; onSearchBlur() }"
      />
      <a v-if="!isSmallScreen" href="#" target="_blank">
        <Tooltip :icon="false" msg="使用文档" placement="bottom"><i class="fa fa-book" style="font-size: 1.4rem"></i></Tooltip>
      </a>
      <a v-if="!isSmallScreen" href="#" target="_blank" >
        <Tooltip :icon="false" msg="提交反馈" placement="bottom"><i class="fa fa-commenting" style="font-size: 1.4rem"></i></Tooltip>
      </a>
      <a-popover placement="bottomRight" trigger="click">
        <template #title>
          <span style="font-size: 1.2rem"><i class="fa fa-user" style="width: 2rem"></i>
            {{ user?.cn_name || user?.username }}
          </span>
        </template>
        <template #content>
          <hr/>
          <router-link to="/base/user/edit" class="link"><i class="fa fa-user-gear" style="width: 2rem"></i> <b>个人设置</b></router-link>
          <hr/>
          <router-link to="/base/user/join" class="link"><i class="fa fa-user-group" style="width: 2rem"></i> <b>切换账号</b></router-link>
        </template>
        <img v-if="avatarUrl" :src="avatarUrl" style="width:40px;height:40px;cursor:pointer;border-radius:50%;object-fit:cover;" />
        <span v-else class="header-user-name text-primary" style="font-size: 1.4rem; cursor: pointer">
          <i class="fa fa-circle-user" style="margin-right: 0.3rem;"></i>{{ user?.cn_name || user?.username || '' }}
        </span>
      </a-popover>
    </a-space>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, inject } from 'vue'
import { useRouter } from 'vue-router'
import lib from '@libs/lib.ts'
import Tooltip from '@components/tooltip.vue'

interface L1Item {
  path?: string
  name: string
  icon: string
  active?: boolean
}

interface User {
  cn_name?: string
  username: string
}

interface Props {
  l1?: L1Item[]
  env?: string
  user?: User
}

defineProps<Props>()
const router = useRouter()

const ip = ref('')
const searchWidth = ref('200px')
const avatarUrl = ref('')

const toggleDrawer = inject<() => void>('toggleDrawer')
const openDrawer = () => toggleDrawer?.()

const isSmallScreen = ref(false)
const searchCollapsed = ref(false)

function checkScreen() {
  isSmallScreen.value = window.matchMedia('(max-width: 576px)').matches
  if (isSmallScreen.value) {
    searchCollapsed.value = true
  } else {
    searchCollapsed.value = false
  }
}

function onSearchBlur() {
  if (isSmallScreen.value) {
    setTimeout(() => { searchCollapsed.value = true }, 150)
  }
}

onMounted(async () => {
  avatarUrl.value = await lib.loadAvatar()
  checkScreen()
  window.addEventListener('resize', checkScreen)
})
onUnmounted(() => {
  window.removeEventListener('resize', checkScreen)
})

const onSearch = () => {
  lib.redirect('/?ip[]=' + ip.value, '_self', router)
}
</script>

<style scoped>
.l1 .a {
  font-weight: bold;
  font-size: 1.2rem;
}
.l1 .a .l1-icon {
  margin-right: 0.4rem;
}

.prod {
  background: white;
}
.prod .l1 .a {
  color: black;
}
.prod .l1 .a.active {
  color: blue;
}

.dev {
  background: white;
}
.dev .l1 .a {
  color: #666;
}
.dev .l1 .a.active {
  color: #52c41a;
}

.test {
  background: #1a237e;
}
.test .l1 .a {
  color: white;
}
.test .l1 .a.active {
  color: yellow;
}
.test .header-user-name {
  color: white;
}

.space {
  display: flex;
  align-items: center;
}
.pull-right {
  margin-left: auto;
}
.pull-right .ant-input-search {
  display: flex;
  align-items: center;
}
.l1 .a:hover {
  color: gray;
}

.hamburger {
  display: none;
  align-items: center;
  justify-content: center;
  width: 40px;
  height: 40px;
  margin: 0 -8px 0 0;
  padding: 0;
  border: none;
  background: transparent;
  cursor: pointer;
  font-size: 1.4rem;
  color: inherit;
}
@media (max-width: 992px) {
  .hamburger {
    display: flex;
    margin-left: -6px;
    margin-right: 8px;
  }
}

.search-icon-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 36px;
  height: 36px;
  padding: 0;
  border: none;
  background: transparent;
  cursor: pointer;
  font-size: 1.2rem;
  color: inherit;
}

.header-user-name {
  cursor: pointer;
  font-weight: 600;
  font-size: 0.95rem;
  max-width: 120px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

@media (max-width: 768px) {
  .header-right.space {
    gap: 12px;
  }
}

/* 小屏时一级导航图标更紧凑（与 custom.css 中 .auto-hide-l1 隐藏文字对应） */
@media (max-width: 1000px) {
  .l1.space {
    gap: 10px !important;
  }
  .l1 .a .l1-icon {
    margin-right: 0 !important;
  }
  .l1 .a {
    padding: 4px 6px;
    min-width: auto;
  }
}
@media (max-width: 576px) {
  .l1.space {
    gap: 8px !important;
  }
  .l1 .a {
    padding: 2px 4px;
    font-size: 1.1rem;
  }
}
</style>
<style>
.search-input .ant-input {
  height: 32px;
}
</style>
