<template>
  <div>
    <template v-for="(l, i) in l2" :key="i">
      <template v-if="! isEmpty(l.subs)">
        <div class="l2-folder" :class="l.active ? 'active' : ''" @click="toggle(l.path)">
          <i :class="'fas fa-' + l.icon"></i>{{ l.name }}
          <span class="pull-right" :class="`fa fa-angle-` + (l.path == openPath ? 'down' : 'right')"></span>
        </div>

        <div v-if="l.path == openPath" class="submenu">
          <div v-for="(s, si) in l.subs" :key="si" class="sub-item">
            <router-link :to="s.path" class="link">
              <span class="sub-name" :class="s.active ? 'active' : ''">
                <b style="margin-right: 0.5rem">·</b>{{ s.name }}
              </span>
            </router-link>
          </div>
        </div>
      </template>
    </template>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'
import { isEmpty } from 'lodash'

const props = defineProps(['l2', 'env'])
const openPath = ref('')

function updateOpenPath() {
  if (!props.l2 || props.l2.length === 0) {
    openPath.value = ''
    return
  }
  for (const l of props.l2) {
    if (l.active) {
      openPath.value = l.path
      return
    }
  }
  openPath.value = ''
}

// 初始化
updateOpenPath()

// 监听 l2 变化，响应式更新 openPath
watch(() => props.l2, () => {
  updateOpenPath()
}, { deep: true })

const toggle = (path) => {
  openPath.value = openPath.value === path ? '' : path
}
</script>

<style scoped>
.l2-folder {
  margin: 0.7rem 0;
  transition: all 0.3s;
  font-size: 1.2rem;
  font-weight: bold;
}
.l2-folder.active {
  color: yellow;
}
.l2-folder:hover {
  color: gray;
}
i {
  font-size: 1.2rem;
  width: 2.5rem;
}
.fa.pull-right {
  margin-top: 0.3rem;
  font-size: 1.2rem;
}
.submenu {
  overflow: hidden;
}
.sub-item {
  margin: 0.7rem 0;
}
.sub-item:first-child {
  margin-top: 0.1rem;
}
.sub-item:last-child {
  margin-bottom: 0.1rem;
}
.sub-name {
  color: white;
  margin-left: 1.5rem;
  font-size: 1rem;
  font-weight: bold;
  text-decoration: none;
}
.sub-name.active {
  color: yellow !important;
}
.sub-name:hover {
  color: gray;
}
</style>
