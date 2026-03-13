import type { LoadModalFunc } from '@libs/modal'
import type { FetchFunc } from '@libs/fetch'

import lib from '@libs/lib'
import { isEmpty } from 'lodash'

export interface Menu {
  title: string
  icon: string
  url: string
  type: string
  color: string
}

export interface TableMenu extends Menu {
  alone: boolean
  args: MenuArg[]
  conds: MenuCond[] | null
}

export interface MenuArg {
  key: string
  val: string
}

export interface MenuCond {
  key: string
  val: any
  comp: string
}

export function useMenu(loadModal: LoadModalFunc, fetch: FetchFunc) {
  const runMenu = async (m: Menu) => {
    if (m.type == 'modal') {
      loadModal(m.url)
    } else if (m.type == 'link') {
      lib.redirect(m.url)
    } else if (m.type == 'async') {
      const ok = await lib.confirmCurl(m.url)
      if (ok) {
        fetch()
      }
    }
  }
  return { runMenu }
}

export function filterMenus(tms: TableMenu[], d: any): TableMenu[] {
  const ans = []
  for (const tm of tms) {
    if (matchMenu(tm, d)) {
      ans.push(tm)
    }
  }
  return ans
}

export function calcMenu(tm: TableMenu, d: any): Menu {
  let url = tm.url
  if (! isEmpty(tm.args)) {
    const params: Record<string, any> = {}
    for (let a of tm.args) {
      params[a.val] = d[a.key]
    }
    if (url.includes('?')) {
      url += '&' + (new URLSearchParams(params).toString())
    } else {
      url += '?' + (new URLSearchParams(params).toString())
    }
  }
  return { url, type: tm.type, title: tm.title, icon: tm.icon, color: tm.color }
}

function matchMenu(tm: TableMenu, d: any): boolean {
  if (! tm.conds || tm.conds.length == 0) {
    return true
  }
  for (const cond of tm.conds) {
    if (! matchCond(cond, d)) {
      return false
    }
  }
  return true
}

function matchCond(cond: MenuCond, d: any): boolean {
  const { key, val, comp } = cond
  switch (comp) {
    case 'EQ':
      return d[key] == val
    case 'NQ':
      return d[key] != val
    case 'LT':
      return d[key] < val
    case 'LE':
      return d[key] <= val
    case 'GT':
      return d[key] > val
    case 'GE':
      return d[key] >= val
    case 'IN':
      return val.includes(d[key])
    case 'NIN':
      return ! val.includes(d[key])
  }
  return true
}