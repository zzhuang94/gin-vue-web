<template>
  <a v-if="v.link_prefix" :href="`${v.link_prefix}${d[k]}`" target="_blank">{{ d[k] }}</a>
  <span v-else-if="v.format_func" v-html="strlib.formatFunc(v.format_func, d[k])" :class="calcClass(d, k, v)"
    @click="v.click_swal_modal && swal.modal(v.name, d[k], v.swal_width)">
  </span>
  <span v-else-if="v.click_swal_modal && ! v.format_func" class="fa fa-search text-info click-swal"
    @click="swal.modal(v.name, d[k], v.swal_width)">
  </span>
  <span v-else-if="v.click_swal_info" class="fa fa-search text-info click-swal" @click="swal.info(v.name, d[k])"></span>
  <span v-else-if="v.wrap_badge" v-html="strlib.wrapBadge(v.wrap_badge, d[k])" :class="calcClass(d, k, v)"></span>
  <span v-else-if="v.textarea && v.split_sep && d[k]" v-html="d[k].split(v.split_sep).join('<br/>')" :class="calcClass(d, k, v)"></span>
  <pre v-else-if="v.textarea && v.json && d[k]" :class="calcClass(d, k, v)">{{ strlib.formatJson(d[k]) }}</pre>
  <span v-else v-html="lib.displayDK(d, k, v)" :class="calcClass(d, k, v)"></span>
</template>

<script setup>
import lib from '@libs/lib.ts'
import strlib from '@libs/strlib.ts'
import swal from '@libs/swal.ts'

const props = defineProps(['d', 'k', 'v'])

const calcClass = (d, k, v) => {
  return [
    v.bold ? 'span-bold' : '',
    v.textcolor ? 'text-' + v.textcolor : '',
    v.dangers && v.dangers.includes(d[k]) ? 'text-danger' : '',
  ]
}
</script>

<style scoped>
.span-bold {
  font-weight: bold;
}
</style>
