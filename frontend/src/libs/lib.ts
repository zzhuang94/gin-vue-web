import { defineAsyncComponent } from 'vue'
import { isEmpty, isObject } from 'lodash'
import axios from 'axios'
import swal from '@libs/swal.ts'

const pages = import.meta.glob('../**/*.vue')

function isInternalUrl(url: string): boolean {
  return !/^(https?:|mailto:|tel:|\/\/)/.test(url)
}

function smartUrl(url: string): string {
  if (!isInternalUrl(url)) {
    return url
  }
  if (! url.startsWith('/')) {
    const path = window.location.pathname
    const lastSlash = path.lastIndexOf('/')
    url = path.substring(0, lastSlash + 1) + url
  }
  return '/web' + url
}

async function post(url: string, params?: any): Promise<any> {
  return axios.post(smartUrl(url), params)
}

async function loadAvatar(): Promise<string> {
  const r = await curl('/base/user/get-avatar')
  if (r?.code === 1) {
    return 'data:image/jpeg;base64,' + r.data
  }
  return ''
}

async function loadModal(url: string, m: any, p: any, args: any = {}): Promise<void> {
  try {
    swal.startLoading()
    const r = await post(url, args)
    if (r.data.code === -1) {
      swal.warn("抱歉，您无权限操作！")
      return
    }
    const { page, props } = r.data
    m.value = loadComponent(page)
    p.value = props
  } catch (error) {
    console.error('Error fetching component:', error)
  } finally {
    swal.stopLoading()
  }
}

async function reloadModal(url: string, p: any): Promise<void> {
  try {
    const resp = await post(url)
    p.value = resp.data.props
  } catch (error) {
    console.error('Error fetching component:', error)
  }
}

function loadComponent(page: string): any {
  page = '../' + page + '.vue'
  const pageLoader = pages[page]
  if (!pageLoader) {
    throw new Error(`Component not found: ${page}`)
  }
  return defineAsyncComponent(() => pageLoader() as Promise<any>)
}

async function curl(url: string, params?: any): Promise<any> {
  try {
    const resp = await post(url, params)
    return resp.data
  } catch (error) {
    console.error("后端错误：" + error)
    return null
  }
}

async function ajax(url: string, params?: any): Promise<any> {
  try {
    const resp = await post(url, params)
    return checkResp(resp)
  } catch (error) {
    swal.error("后端错误：" + error)
    return false
  }
}

function checkResp(resp: any): boolean {
  const { code, data } = resp.data
  if (code === 1) {
    swal.succ("OK", data)
    return true
  }
  swal.error("Failed", data)
  return false
}

function redirect(url: string, target: string = '_blank', router: any = null): void {
  if (! isInternalUrl(url)) {
    window.open(url, target)
    return
  }

  // 如果是内部路由，并且传入了 router
  if (router) {
    router.push(url)
  } else {
    console.error("未传入 router 实例，无法执行路由跳转")
    // 降级方案：直接修改 location.href
    window.location.href = url
  }
}

function back(): void {
  window.history.back()
}

async function confirmCurl(url: string, msg: string = ''): Promise<any> {
  const confirmed = await swal.confirm('确定操作？', msg)
  if (! confirmed) {
    return
  }
  swal.startLoading()
  const ans = await ajax(url)
  swal.stopLoading()
  return ans
}

function mapToUriParams(map: any): any  {
  const ans: Record<string, any> = {}
  for (const k in map) {
    const v = map[k]
    if (Array.isArray(v)) {
      ans[k] = v
    } else {
      ans[k] = v
    }
  }
  return ans
}

function filterByLabel(inputValue: string, option: any): boolean {
  return option.label.toLowerCase().indexOf(inputValue.toLowerCase()) >= 0
}

function removeHtmlTags(html: string): string  {
    const parser = new DOMParser()
    const doc = parser.parseFromString(html, 'text/html')
    return doc.body.textContent || ''
}

function displayDK(r: any, v: any, autoBreak: boolean = true): string {
  if (r.limit && ! r.limit_map) {
    r.limit_map = {}
    for (let i = 0; i < r.limit.length; i++) {
      let l = r.limit[i]
      if (! l.badge) {
        l.badge = 'rand-' + i % 40
      }
      r.limit_map[l.key] = l
    }
  }
  if (r.limit_map) {
    if (r.split_sep) {
      const arr = []
      for (let dv of v.split(r.split_sep)) {
        arr.push(calcLimit(dv, r.limit_map))
      }
      return arr.join('<br/>')
    }
    return calcLimit(v, r.limit_map)
  }
  if (autoBreak) {
    return v.replace(/\n/g, '<br/>')
  }
  if (r.split_sep && r.textarea) {
    return v.split(r.split_sep).join('<br/>')
  }
  return v
}

function calcLimit(value: string, map: any): any {
  const v: any = map?.[value] ?? value
  if (! isObject(v)) {
    return v
  }
  if (! ('badge' in v) || isEmpty(v.badge)) {
    return v && 'label' in v ? v.label : ''
  }
  let ans = '<span class="badge badge-' + (v && 'badge' in v ? v.badge : '') + '">'
  ans += v && 'label' in v ? v.label : ''
  if (v && 'spin' in v && v.spin) {
    ans += ' <i class="fa fa-spinner fa-spin"></i>'
  }
  return ans + '</span>'
}

function isLimitBadge(obj: any): boolean {
  const firstValue = Object.values(obj)[0]
  if (firstValue && typeof firstValue === 'object' && firstValue !== null) {
    return firstValue.constructor === Object && 'badge' in firstValue
  }
  return false
}

export default {
  loadAvatar,
  loadModal,
  reloadModal,
  loadComponent,
  curl,
  ajax,
  confirm,
  redirect,
  back,
  confirmCurl,
  mapToUriParams,

  filterByLabel,
  removeHtmlTags,

  displayDK,
  calcLimit,
  isLimitBadge,
}
