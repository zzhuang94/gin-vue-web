<template>
  <div style="font-size: 1.1rem; color: gray">
    <router-link to="/" class="link"><i class="fa fa-home" style="font-size: 1.4rem"></i> 首页</router-link>
    <template v-for="t, i in title" :key="i">
      <span style="margin: 0 1rem">/</span><router-link :to="t.path" class="link">{{ t.name }}</router-link>
    </template>
    <span @click="lib.back()" class="pull-right link">
      <i class="fa fa-arrow-left"></i> 后退
    </span>
  </div>
</template>

<script setup>
import lib from '@libs/lib.ts'

defineProps(['title'])
</script>
