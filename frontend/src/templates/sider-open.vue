<template>
  <div>
    <template v-for="(l, i) in l2" :key="i">
      <template v-if="! isEmpty(l.subs)">
        <div class="l2-folder" :class="l.active ? 'active' : ''" @click="toggle(l.path)">
          <i :class="'fas fa-' + l.icon"></i>{{ l.name }}
          <span class="pull-right" :class="`fa fa-angle-` + (l.path == openPath ? 'down' : 'right')"></span>
        </div>

        <div v-if="l.path == openPath" class="submenu">
          <div v-for="(s, si) in l.subs" :key="si" style="margin: 1rem 0; font-size: 0.9rem;">
            <router-link :to="s.path" class="link">
              <span class="sub-name" :class="s.active ? 'active' : ''">
                <b style="margin: 0 0.6rem 0 0.6rem">·</b>{{ s.name }}
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
  margin-top: 0.5rem;
  margin-bottom: 1.2rem;
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
  font-size: 1.4rem;
  width: 3rem;
}
.fa.pull-right {
  margin-top: 3px;
  font-size: 1.2rem;
}
.sub-name {
  color: white;
  margin-left: 1.3rem;
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
