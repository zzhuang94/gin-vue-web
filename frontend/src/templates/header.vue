<template>
  <div :class="env">
    <a-space :size="30" class="l1 space">
      <template v-for="l, i in l1" :key="i">
        <router-link v-if="l.path" :to="l.path">
          <span class="a" :class="l.active ? 'active' : ''">
            <i :class="'fas fa-' + l.icon" style="margin-right: 0.4rem"></i>
            <span class="auto-hide-l1">{{ l.name }}</span>
          </span>
        </router-link>
      </template>
    </a-space>
    <a-space :size="20" class="pull-right space">
      <a-input-search v-model:value="ip" :style="{ width, transition: 'width 0.3s ease' }" class="search-input"
        placeholder="全局搜索" enter-button
        @search="onSearch" @focus="() => width = '300px'" @blur="() => width = '200px'" />

      <a href="#" target="_blank">
        <Tooltip :icon="false" msg="使用文档" placement="bottom"><i class="fa fa-book" style="font-size: 1.4rem"></i></Tooltip>
      </a>
      <a href="#" target="_blank" >
        <Tooltip :icon="false" msg="提交反馈" placement="bottom"><i class="fa fa-commenting" style="font-size: 1.4rem"></i></Tooltip>
      </a>
      <a-popover placement="bottomRight" trigger="click">
        <template #title>
          <span style="font-size: 1.2rem"><i class="fa fa-user" style="width: 2rem"></i>
            {{ user.cn_name || user.username }}
          </span>
        </template>
        <template #content>
          <hr/>
          <router-link to="/base/user/edit" class="link"><i class="fa fa-user-gear" style="width: 2rem"></i> <b>个人设置</b></router-link>
          <hr/>
          <router-link to="/base/user/join" class="link"><i class="fa fa-user-group" style="width: 2rem"></i> <b>切换账号</b></router-link>
        </template>
        <img v-if="avatarUrl" :src="avatarUrl" style="width:40px;height:40px;cursor:pointer;border-radius:50%;object-fit:cover;" />
        <div v-else style="width:40px;height:40px;cursor:pointer;border-radius:50%;background:#1890ff;display:flex;align-items:center;justify-content:center;color:white;font-weight:bold;font-size:16px;">
          {{ (user.cn_name || user.username || 'A').charAt(0).toUpperCase() }}
        </div>
      </a-popover>
    </a-space>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import lib from '@libs/lib.ts'
import Tooltip from '@components/tooltip.vue'

const props = defineProps(['l1', 'env', 'user'])
const router = useRouter()

const ip = ref('')
const width = ref('200px')
const avatarUrl = ref('')

onMounted(async () => {
  avatarUrl.value = await lib.loadAvatar()
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
</style>
<style>
.search-input .ant-input {
  height: 32px;
}
</style>
