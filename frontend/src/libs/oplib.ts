import { isEmpty } from 'lodash'
import type { TableMenu, Menu, MenuCond } from '@libs/frm.ts'

function filterMenus(tms: TableMenu[], d: any): TableMenu[] {
  const ans = []
  for (const tm of tms) {
    if (matchMenu(tm, d)) {
      ans.push(tm)
    }
  }
  return ans
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

function calcMenu(tm: TableMenu, d: any): Menu {
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

export default {
  filterMenus,
  calcMenu,
}
