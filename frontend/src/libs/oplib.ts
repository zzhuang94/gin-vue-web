import { isEmpty, isObject } from 'lodash'

function filterOption(option: any, data: any): any {
  const ans = []
  for (const op of option) {
    if (isOpMatchData(op, data)) {
      ans.push(op)
    }
  }
  return ans
}

function isOpMatchData(op: any, d: any): boolean {
  if (isEmpty(op.cond)) {
    return true
  }
  const [k, t, v] = op.cond
  switch (t) {
    case 'EQ':
      return d[k] == v
    case 'NQ':
      return d[k] != v
    case 'LT':
      return d[k] < v
    case 'LE':
      return d[k] <= v
    case 'GT':
      return d[k] > v
    case 'GE':
      return d[k] >= v
    case 'IN':
      return v.includes(d[k])
    case 'NIN':
      return ! v.includes(d[k])
  }
  return true
}

function calcOp(op: any, d: any): any {
  let url = op.url
  if (op.args) {
    const params: Record<string, any> = {}
    for (let k of op.args) {
      if (isObject(k) && 'v' in k && 'k' in k) {
        params[(k as { v: string }).v] = d[(k as { k: string }).k]
      } else {
        params[k] = d[k]
      }
    }
    if (url.includes('?')) {
      url += '&' + (new URLSearchParams(params).toString())
    } else {
      url += '?' + (new URLSearchParams(params).toString())
    }
  }
  return { url, type: op.type, title: op.title }
}

export default {
  filterOption,
  calcOp,
}
