function lineExplode(text: string): string[] {
  if (! text) {
    return []
  }
  return text
    .split('\n')                      // 按行分割成数组
    .map(line => line.trim())         // 去除每行首尾空格
    .filter(line => line !== '')      // 过滤空行
    .filter((line, index, arr) => {   // 去重（保留首次出现的行）
      return arr.indexOf(line) === index
    })
}

function isJson(str: string): boolean {
  try {
    JSON.parse(str)
    return true
  } catch (e) {
    return false
  }
}

function formatJson(str: string): string {
  try {
    const jsonObj = JSON.parse(str)
    return JSON.stringify(jsonObj, null, 2)
  } catch (e) {
    return str
  }
}

function formatFunc(func: string, val: any): string {
  switch (func) {
    case 'formatXxx':
      return formatXxx(val)
  }
  return ''
}


function formatXxx(val: string): string {
  return 'Xxx:' + val
}

function wrapBadge(c: string, val: string): string {
  const ss = []
  for (let v of val.split(',')) {
    if (! v) {
      continue
    }
    ss.push(buildBadge(c, v))
  }
  return ss.join('&nbsp')
}

function buildBadge(badge: string, label: string, icon: string = '') {
  let ans = `<span class='badge badge-${badge}'>`
  if (icon) {
    ans += `<span class="fa fa-${icon}"></span> `
  }
  return ans + label + '</span>'
}

function formatDomain(domain: string): string {
    if (domain.endsWith(".")) {
        return domain
    }
    return domain + "."
}

function isDomain(str: string): boolean {
  if (str.length != 0) {
    var reg = /^(@\.)?[A-Za-z0-9\-\.\*]+$/
    if (str.match(reg) != null) {
      return true
    }
  }
  return false
}

function isIP(str: string): boolean {
  return isIPv4(str) || isIPv6(str)
}

function isIPv4(str: string): boolean {
  const segments = str.split('.')
  return segments.length === 4 && segments.every(seg => {
    const num = +seg
    return num >= 0 && num <= 255 && String(num) === seg
  })
}

function isIPv6(str: string): boolean {
  // 处理压缩格式（::）
  if (str.includes('::')) {
    const parts = str.split('::')
    if (parts.length > 2) return false; // 只能有一个::

    const leftParts = parts[0] ? parts[0].split(':').filter(Boolean) : []
    const rightParts = parts[1] ? parts[1].split(':').filter(Boolean) : []

    // 压缩后的总段数不能超过7
    if (leftParts.length + rightParts.length > 7) return false

    // 验证每一部分
    return [...leftParts, ...rightParts].every(part =>
      /^[0-9a-fA-F]{1,4}$/.test(part)
    )
  }

  // 标准格式（无压缩）
  const parts = str.split(':')
  if (parts.length !== 8) return false

  return parts.every(part =>
    /^[0-9a-fA-F]{1,4}$/.test(part)
  )
}

function isInt(str: string, validPositive: boolean = false): boolean {
  try {
    const num = BigInt(str)
    if (validPositive && num <= 0) {
      return false
    }
    return str === num.toString()
  } catch {
    return false
  }
}

function checkFirstDup(ss: string[]): string {
  const seen = new Set()
  for (const s of ss) {
    if (seen.has(s)) {
      return s
    }
    seen.add(s)
  }
  return ''
}

function prettyFlow(flow: any): string {
  if (typeof flow == 'string') {
    flow = parseFloat(flow)
  }
  if (flow >= 1000) {
    flow = flow / 1000
    return flow.toFixed(2) + 'Gb'
  }
  if (flow < 10) {
    return flow.toFixed(2) + 'Mb'
  }
  return flow.toFixed(0) + 'Mb'
}

export default {
  lineExplode,
  formatFunc,
  wrapBadge,
  buildBadge,
  formatDomain,
  isJson,
  formatJson,
  isDomain,
  isIP,
  isIPv4,
  isIPv6,
  isInt,
  checkFirstDup,
  prettyFlow,
}
