<template>
  <head>
    <meta charset="utf-8">
    <title>{{ headTitle }}</title>
    <meta name="description" content="Base button default style">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
  </head>

  <div class="container">
    <div :class="['sider', env, fold ? 'fold' : 'open']">
      <div class="sider-header">
        <i @click="toggleFold" :class="'fa fa-' + (fold ? 'indent' : 'dedent')" style="cursor: pointer; font-size: 1.5rem; width: 2.3rem; margin-right: 0.2rem"></i>
        <span v-if="!fold" class="sider-title">{{ name }}</span>
      </div>
      <SiderFold v-if="fold" :l2="l2" :env="env" />
      <SiderOpen v-else :l2="l2" :env="env" />
    </div>
    <div class="main-content">
      <Header class="header" :l1="l1" :env="env" :user="user" />
      <div class="scroll-container">
        <div class="content-wrapper">
          <Title class="title" :title="title" />
          <div class="body">
            <slot></slot>
          </div>
          <Footer class="footer" />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import lib from '@libs/lib.ts'
import Header from './header.vue'
import Title from './title.vue'
import Footer from './footer.vue'
import SiderFold from './sider-fold.vue'
import SiderOpen from './sider-open.vue'

const props = defineProps(['name', 'env', 'fold', 'l1', 'l2', 'title', 'user'])
const fold = ref(props.fold)

watch(() => props.fold, (newVal) => {
  fold.value = newVal
})

const headTitle = computed(() => {
  let ans = props.name
  if (props.title && props.title.length > 0) {
    ans += ' - ' + props.title[props.title.length - 1].name
  }
  return ans
})

const toggleFold = () => {
  fold.value = !fold.value
  const url = `/base/user/set?key=fold&val=` + (fold.value ? '1' : '0')
  lib.curl(url)
}
</script>

<style scoped>
.container {
  display: flex;
  min-height: 100vh;
}

.sider {
  width: 200px;
  height: 100vh;
  color: white;
  flex-shrink: 0;
  position: sticky;
  top: 0;
  overflow-y: auto;
}
.sider-title {
 font-size: 1.8rem;
 font-weight: bold;
 padding-top: 0.1rem;
}
.sider.prod {
  background: black;
}
.sider.dev {
  background: green;
}
.sider.test {
  background: #1a237e;
}

.sider.fold {
  text-align: center;
  cursor: pointer;
  width: 50px;
}
.sider.open {
  cursor: pointer;
  padding: 0 0.5rem;
}

.sider.fold>.sider-header {
  justify-content: center;
}
.sider-header {
  line-height: 50px;
  height: 50px;
  display: flex;
  align-items: center;
  transition: all 0.1s;
}

.main-content {
  display: flex;
  flex-direction: column;
  flex: 1;
  min-height: 100vh;
}

.header {
  position: sticky;
  top: 0;
  z-index: 100;
  padding: 0px 20px;
  height: 50px;
  line-height: 50px;
  border-bottom: 1px solid #eee;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.scroll-container {
  flex: 1;
  overflow-y: auto;
}

.content-wrapper {
  min-height: calc(100vh - 50px); /* 减去 header 高度 */
  display: flex;
  flex-direction: column;
}

.title {
  flex-shrink: 0;
  padding: 0px 20px;
  height: 40px;
  line-height: 40px;
  background: #ffffff;
  box-shadow: 0px 1px 15px 1px rgba(69, 65, 78, 0.08);
}

.body {
  flex: 1;
  background-color: white;
  margin: 10px 10px 5px 10px;
  box-shadow: -2px -8px 41px -14px rgba(202, 210, 222, 0.61);
}

.footer {
  padding: 0px 20px;
  height: 40px;
  line-height: 40px;
  background: #ffffff;
  flex-shrink: 0;
  box-shadow: -2px -8px 41px -14px rgba(202, 210, 222, 0.61);
}
</style>
