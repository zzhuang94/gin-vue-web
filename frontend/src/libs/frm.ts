export interface RuleLimit {
  key: string | number
  label: string
  badge: string
  spin: boolean
}

export interface RuleTrans {
  ajax: boolean
  db: string
  sql: string
  table: string
  key: string
  val: string
}

export interface RuleValidation {
  is_int: boolean
  int_range: boolean
  int_min: number
  int_max: number

  is_float: boolean
  float_range: boolean
  float_min: number
  float_max: number

  is_ip: boolean
  is_ipv4: boolean
  is_ipv6: boolean

  regex: string
}

// 统一的前端字段规则定义，对应后端 g.Rule，大多数字段保持 snake_case
export interface Rule {
  // 基本信息
  key: string
  name: string

  // 展示 & 交互
  default: any
  readonly: boolean
  required: boolean
  describe: string

  // 文本类型
  textarea: boolean
  date: boolean
  datetime: boolean
  json: boolean

  // 分割符
  split_sep: string

  // 宽度 & 自动隐藏 & 隐藏 & 操作使用
  width: string
  auto_hide: string
  hide: boolean

  // 操作使用，不显示，但是后端要输出这个字段
  op_use: boolean

  // 搜索，0: none, 1: eq, 2: like, 3: in
  search: number

  // 下拉选项
  limit: RuleLimit[] | null
  limit_map: Record<string, RuleLimit> | null

  // 转译与校验
  trans: RuleTrans | null
  validation: RuleValidation | null

  // table / td 里的扩展字段
  bold: boolean
  textcolor: string
  suffix: string
  prefix: string

  no_sort: boolean
  link_prefix: string
  format_func: string
  click_swal_modal: boolean
  swal_width: string
  click_swal_info: boolean
  wrap_badge: any
}

