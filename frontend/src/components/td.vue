<template>
  <span>
    <a v-if="r.link_prefix" :href="`${r.link_prefix}${v}`" target="_blank">{{ v }}</a>
    <span v-else-if="r.format_func" v-html="strlib.formatFunc(r.format_func ?? '', v)" :class="calcClass(r, v)"
      @click="r.click_swal_modal && r.name && swal.modal(r.name, String(v ?? ''), r.swal_width)">
    </span>
    <span v-else-if="r.click_swal_modal && !r.format_func" class="fa fa-search text-info click-swal"
      @click="r.name && swal.modal(r.name, String(v ?? ''), r.swal_width)">
    </span>
    <span v-else-if="r.click_swal_info" class="fa fa-search text-info click-swal" @click="r.name && swal.info(r.name, String(v ?? ''))"></span>
    <span v-else-if="r.wrap_badge" v-html="strlib.wrapBadge(r.wrap_badge, v)" :class="calcClass(r, v)"></span>
    <span v-else-if="r.textarea && r.split_sep && v" v-html="v.split(r.split_sep).join('<br/>')"
      :class="calcClass(r, v)"></span>
    <pre v-else-if="r.textarea && r.json && v" :class="calcClass(r, v)">{{ strlib.formatJson(v) }}</pre>
    <span v-else v-html="lib.displayDK(r, v)" :class="calcClass(r, v)"></span>
  </span>
</template>

<script setup lang="ts">
import lib from '@libs/lib.ts'
import strlib from '@libs/strlib.ts'
import swal from '@libs/swal.ts'

interface Rule {
  link_prefix?: string
  format_func?: string
  click_swal_modal?: boolean
  swal_width?: string
  click_swal_info?: boolean
  wrap_badge?: any
  textarea?: boolean
  split_sep?: string
  json?: boolean
  bold?: boolean
  textcolor?: string
  dangers?: any[]
  name?: string
  [key: string]: any
}

interface Props {
  r: Rule
  v: any
}

defineProps<Props>()

const calcClass = (r: Rule, v: any): string[] => {
  return [
    r.bold ? 'span-bold' : '',
    r.textcolor ? 'text-' + r.textcolor : '',
    r.dangers && r.dangers.includes(v) ? 'text-danger' : '',
  ].filter(Boolean) as string[]
}
</script>

<style scoped>
.span-bold {
  font-weight: bold;
}
</style>
