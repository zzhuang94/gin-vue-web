<template>
  <div class="portlet">
    <div class="portlet-body">
      <Card title="Button 按钮样式" icon="mouse-pointer" color="primary">
        <div style="margin-bottom: 10px">
          <p><strong>功能：</strong>展示按钮的各种样式和尺寸</p>
        </div>
        <div style="margin-bottom: 15px">
          <div style="font-weight: bold; margin-bottom: 5px">button class="btn btn-${color}"</div>
          <a-space>
            <button v-for="c, i in colors" :key="i" :class="`btn btn-${c}`">{{ c }}</button>
          </a-space>
        </div>
        <div style="margin-bottom: 15px">
          <div style="font-weight: bold; margin-bottom: 5px">button class="btn btn-outline-${color}"</div>
          <a-space>
            <button v-for="c, i in colors" :key="i" :class="`btn btn-outline-${c}`">{{ c }}</button>
          </a-space>
        </div>
        <div style="margin-bottom: 15px">
          <div style="font-weight: bold; margin-bottom: 5px">button class="btn btn-success btn-${size}"</div>
          <a-space>
            <button v-for="sz, si in ['sm', 'md', 'lg']" :key="si" :class="`btn btn-${sz} btn-success`">btn btn-success btn-{{ sz }}</button>
          </a-space>
        </div>
      </Card>

      <Card title="Badge 徽章样式" icon="tags" color="success">
        <div style="margin-bottom: 10px">
          <p><strong>功能：</strong>展示徽章的各种颜色样式</p>
        </div>
        <div style="margin-bottom: 15px">
          <div style="font-weight: bold; margin-bottom: 5px">class="badge badge-${color}"</div>
          <a-space>
            <span v-for="c, i in colors" :key="i" :class="`badge badge-${c}`">{{ c }}</span>
          </a-space>
        </div>
        <div style="margin-bottom: 15px">
          <div style="font-weight: bold; margin-bottom: 5px">class="badge badge-rand-${n}" (0-39 随机配色方案)</div>
          <p style="color: #666; font-size: 0.9rem; margin-bottom: 10px">
            <strong>用法：</strong>使用 <code>badge-rand-0</code> 到 <code>badge-rand-39</code> 可以生成40种不同的随机配色方案，适用于标签分类显示。所有配色均按照绿色→蓝色→红色→黄色→紫色→橙色的循环模式，前景色和背景色对比明显，确保良好的可读性。
          </p>
          <div style="display: flex; flex-wrap: wrap; gap: 8px; max-height: 400px; overflow-y: auto; padding: 10px; border: 1px solid #e0e0e0; border-radius: 4px;">
            <span v-for="n in 40" :key="n" :class="`badge badge-rand-${n - 1}`" style="margin: 2px 0;">
              rand-{{ (n - 1 < 10 ? '0' : '') + (n - 1) }}
            </span>
          </div>
        </div>
      </Card>

      <Card title="Text 文本样式" icon="font" color="info">
        <div style="margin-bottom: 10px">
          <p><strong>功能：</strong>展示文本的各种颜色样式</p>
        </div>
        <div style="margin-bottom: 15px">
          <div style="font-weight: bold; margin-bottom: 5px">Text text-${color}</div>
          <a-space>
            <span v-for="c, i in colors" :key="i" :class="`text-${c}`">{{ c }}</span>
          </a-space>
        </div>
        <div style="margin-bottom: 15px">
          <div style="font-weight: bold; margin-bottom: 5px">A Text text-${color}</div>
          <a-space>
            <a v-for="c, i in colors" :key="i" :class="`text-${c}`">{{ c }}</a>
          </a-space>
        </div>
      </Card>

      <Card title="Icon 动态图标" icon="star" color="warning">
        <div style="margin-bottom: 10px">
          <p><strong>功能：</strong>展示图标动画和尺寸样式</p>
        </div>
        <div style="margin-bottom: 15px">
          <div style="font-weight: bold; margin-bottom: 5px">Icon animation: fa-${animation}</div>
          <a-space>
            <button v-for="a, i in animations" :key="i" class="btn btn-info btn-lg">
              <i :class="`fa fa-star fa-${a}`"></i> {{ a }}
            </button>
          </a-space>
        </div>
        <div style="margin-bottom: 15px">
          <div style="font-weight: bold; margin-bottom: 5px">Icon size： fa-${size}</div>
          <a-space>
            <button v-for="sz, i in sizes" :key="i" class="btn btn-info btn-lg">
              <i :class="`fa fa-star fa-${sz}`"></i>&nbsp;&nbsp;&nbsp;&nbsp; {{ sz }}
            </button>
          </a-space>
        </div>
      </Card>

      <Card title="Image 图片样式" icon="image" color="primary">
        <div style="margin-bottom: 10px">
          <p><strong>功能：</strong>展示图片的各种样式</p>
        </div>
        <div style="margin-bottom: 15px">
          <a-image :src="ImgTree" alt="Logo" style="width: 100px; height: 100px;" />
        </div>
      </Card>

      <!-- BinSwitch 组件 -->
      <Card title="BinSwitch 开关" icon="toggle-on" color="success">
        <div style="margin-bottom: 10px">
          <p><strong>功能：</strong>支持自定义真假值的开关组件，常用于状态切换</p>
          <p><strong>用法：</strong>&lt;BinSwitch v-model="value" true-value="1" false-value="0" /&gt;</p>
          <p><strong>Props：</strong>modelValue(值), trueValue(真值,默认'1'), falseValue(假值,默认'0'), disabled(禁用)</p>
        </div>
        <BinSwitch v-model="switchValue1" true-value="1" false-value="0" />
        <span style="margin-left: 10px">当前值: {{ switchValue1 }}</span>
        <br />
        <BinSwitch v-model="switchValue2" true-value="yes" false-value="no" />
        <span style="margin-left: 10px">当前值: {{ switchValue2 }}</span>
      </Card>

      <!-- BinCheck 组件 -->
      <Card title="BinCheck 复选框" icon="check-square" color="info">
        <div style="margin-bottom: 10px">
          <p><strong>功能：</strong>支持自定义真假值的复选框组件</p>
          <p><strong>用法：</strong>&lt;BinCheck v-model="value"&gt;选项&lt;/BinCheck&gt;</p>
          <p><strong>Props：</strong>modelValue(值), trueValue(真值,默认'1'), falseValue(假值,默认'0')</p>
        </div>
        <BinCheck v-model="checkValue1">选项1</BinCheck>
        <span style="margin-left: 10px">当前值: {{ checkValue1 }}</span>
      </Card>

      <!-- BtnSelect 组件 -->
      <Card title="BtnSelect 按钮选择器" icon="list" color="accent">
        <div style="margin-bottom: 10px">
          <p><strong>功能：</strong>按钮组形式的选择器，支持单选和多选</p>
          <p><strong>用法：</strong>&lt;BtnSelect name="状态" :items="items" :default="'1'" :required="true" /&gt;</p>
          <p><strong>Props：</strong>name(名称), items(选项数组), default(默认值), required(是否必选)</p>
        </div>
        <BtnSelect name="状态" :items="btnSelectItems" :default="'1'" :required="true" @update="handleBtnSelect" />
      </Card>

      <!-- Input 组件 -->
      <Card title="Input 输入框" icon="keyboard" color="primary">
        <div style="margin-bottom: 10px">
          <p><strong>功能：</strong>统一的输入组件，支持文本、文本域、下拉选择、复选框</p>
          <p><strong>用法：</strong>&lt;Input label="名称" type="input" v-model="value" /&gt;</p>
          <p><strong>Props：</strong>label(标签), type(input/textarea/select/checkbox), value(值), placeholder, options(选项), multiple(多选)</p>
        </div>
        <Input label="名称" type="input" v-model="inputValue1" placeholder="请输入名称" />
        <br />
        <Input label="描述" type="textarea" v-model="inputValue2" :min-rows="3" />
        <br />
        <Input label="类型" type="select" v-model="inputValue3" :options="selectOptions" />
      </Card>

      <!-- TimeRange 组件 -->
      <Card title="TimeRange 时间范围选择器" icon="clock" color="focus">
        <div style="margin-bottom: 10px">
          <p><strong>功能：</strong>日期时间范围选择器</p>
          <p><strong>用法：</strong>&lt;TimeRange :start="start" :end="end" @change="handleChange" /&gt;</p>
          <p><strong>Props：</strong>start(开始时间), end(结束时间), format(格式,默认'YYYY-MM-DD HH:mm:00')</p>
        </div>
        <TimeRange :start="timeStart" :end="timeEnd" @change="handleTimeRange" />
        <div style="margin-top: 10px">开始: {{ timeStart }}, 结束: {{ timeEnd }}</div>
      </Card>

      <!-- Searcher 组件 -->
      <Card title="Searcher 搜索器" icon="search" color="info">
        <div style="margin-bottom: 10px">
          <p><strong>功能：</strong>动态筛选条件搜索组件，支持多种输入类型</p>
          <p><strong>用法：</strong>&lt;Searcher :rules="rules" :arg="arg" @search="handleSearch" /&gt;</p>
          <p><strong>Props：</strong>rules(规则配置), arg(参数对象), bind(绑定数据)</p>
          <p><strong>说明：</strong>rules 中每个字段可配置 limit(选项), trans(翻译), search(是否可搜索)</p>
        </div>
        <Searcher :rules="searcherRules" :arg="searcherArg" @search="handleSearcher" />
      </Card>

      <!-- Card 组件 -->
      <Card title="Card 卡片" icon="id-card" color="brand">
        <div style="margin-bottom: 10px">
          <p><strong>功能：</strong>卡片容器组件，支持多种颜色主题</p>
          <p><strong>用法：</strong>&lt;Card title="标题" icon="table" color="info"&gt;内容&lt;/Card&gt;</p>
          <p><strong>Props：</strong>title(标题), icon(图标), color(颜色主题同button)</p>
        </div>
        <div style="display: flex; gap: 10px; flex-wrap: wrap">
          <Card title="Info" icon="info" color="info" style="width: 200px">内容区域</Card>
          <Card title="Success" icon="check" color="success" style="width: 200px">内容区域</Card>
          <Card title="Warning" icon="exclamation" color="warning" style="width: 200px">内容区域</Card>
        </div>
      </Card>

      <!-- Tooltip 组件 -->
      <Card title="Tooltip 提示框" icon="comment-alt" color="focus">
        <div style="margin-bottom: 10px">
          <p><strong>功能：</strong>提示信息组件，支持自定义位置和颜色</p>
          <p><strong>用法：</strong>&lt;Tooltip msg="提示信息" placement="right"&gt;触发元素&lt;/Tooltip&gt;</p>
          <p><strong>Props：</strong>msg(提示内容), placement(位置), color(颜色), icon(是否显示图标)</p>
        </div>
        <Tooltip msg="这是一个提示信息" placement="right">
          <span style="cursor: pointer; color: #36a3f7">鼠标悬停查看提示</span>
        </Tooltip>
      </Card>

      <!-- TableV 组件 -->
      <Card title="TableV 垂直表格" icon="table" color="default">
        <div style="margin-bottom: 10px">
          <p><strong>功能：</strong>垂直布局的表格组件，用于展示键值对数据</p>
          <p><strong>用法：</strong>&lt;TableV :config="config" :data="data" /&gt;</p>
          <p><strong>Props：</strong>config(字段配置对象), data(数据对象)</p>
        </div>
        <TableV :rules="tableVRules" :data="tableVData" />
      </Card>

      <!-- Lock 组件 -->
      <Card title="Lock 锁定开关" icon="lock" color="warning">
        <div style="margin-bottom: 10px">
          <p><strong>功能：</strong>资源锁定开关组件，用于防止并发编辑</p>
          <p><strong>用法：</strong>&lt;Lock :type="type" :title="title" :locked="locked" :lockers="lockers" /&gt;</p>
          <p><strong>Props：</strong>type(资源类型), title(标题), locked(是否锁定), lockers(锁定用户列表)</p>
          <p><strong>说明：</strong>需要后端接口 /base/lock/lock 支持</p>
        </div>
        <Lock type="test" title="测试锁" :locked="lockValue" :lockers="lockUsers" />
        <Lock type="test" title="测试锁" :locked="! lockValue" :lockers="lockUsers" />
      </Card>

      <!-- Error 组件 -->
      <Card title="Error 错误提示" icon="exclamation-triangle" color="danger">
        <div style="margin-bottom: 10px">
          <p><strong>功能：</strong>错误信息展示组件，支持模态框和页面两种模式</p>
          <p><strong>用法：</strong>&lt;Error msg="错误信息" :is-modal="true" /&gt;</p>
          <p><strong>Props：</strong>msg(错误信息), isModal(是否模态框), width(宽度)</p>
        </div>
        <p style="color: #999">此组件通常用于错误场景，此处仅作展示</p>
      </Card>

    </div>
  </div>
</template>

<script setup>
import { ref, provide } from 'vue'
import Card from '@components/card.vue'
import Button from '@components/button.vue'
import BinSwitch from '@components/bin-switch.vue'
import BinCheck from '@components/bin-check.vue'
import BadgeSelect from '@components/badge-select.vue'
import BtnSelect from '@components/btn-select.vue'
import Input from '@components/input.vue'
import AjaxSelect from '@components/ajax-select.vue'
import TimeRange from '@components/time-range.vue'
import Searcher from '@components/searcher.vue'
import Table from '@components/table.vue'
import List from '@components/list.vue'
import Edit from '@components/edit.vue'
import Tree from '@components/tree.vue'
import Pager from '@components/pager.vue'
import Chart from '@components/chart.vue'
import Tooltip from '@components/tooltip.vue'
import Td from '@components/td.vue'
import TableV from '@components/table-v.vue'
import CollapseTable from '@components/collapse-table.vue'
import Lock from '@components/lock.vue'
import Error from '@components/error.vue'
import ImgTree from '@assets/tree.png'

// 基础样式数据
const colors = ['default', 'secondary', 'metal', 'dark', 'accent', 'primary', 'info', 'focus', 'brand', 'success', 'danger', 'warning']
const animations = ['beat', 'beat-fade', 'bounce', 'fade', 'flip', 'shake', 'spin', 'spin-pulse']
const sizes = ['2xs', 'xs', 'sm', 'md', 'lg', 'xl', '2xl']

// 示例数据
const switchValue1 = ref('1')
const switchValue2 = ref('no')
const checkValue1 = ref('0')
const badgeValue = ref('1')
const badgeOptions = {
  '1': { label: '启用', badge: 'success' },
  '0': { label: '禁用', badge: 'danger' }
}
const btnSelectItems = [
  { key: '1', label: '启用' },
  { key: '0', label: '禁用' }
]
const inputValue1 = ref('')
const inputValue2 = ref('')
const inputValue3 = ref('')
const selectOptions = [
  { value: '1', label: '选项1' },
  { value: '2', label: '选项2' }
]
const ajaxValue = ref('')
const timeStart = ref('')
const timeEnd = ref('')
const searcherArg = ref({})
const searcherRules = [
  { key: 'name', name: '名称', search: true },
  { key: 'status', name: '状态', search: true, limit: [
    { key: '1', label: '启用' },
    { key: '0', label: '禁用' }
  ]}
]
const tableRules = [
  { key: 'id', name: 'ID' },
  { key: 'name', name: '名称' },
  { key: 'status', name: '状态' }
]
const tableData = [
  { id: 1, name: '测试1', status: '启用' },
  { id: 2, name: '测试2', status: '禁用' }
]
const tableOptions = []
const treeData = {
  '1': { name: '节点1', parent_id: '', icon: 'folder' },
  '2': { name: '节点2', parent_id: '1', icon: 'file' },
  '3': { name: '节点3', parent_id: '1', icon: 'file' }
}
const treeOps = [
  { title: '编辑', icon: 'edit', op: 'edit' },
  { title: '删除', icon: 'trash', op: 'delete' }
]
const pagerCurr = ref(1)
const pagerSize = ref(10)
const pagerTotal = ref(100)
const tableVRules = [
  { key: 'name', name: '名称' },
  { key: 'status', name: '状态' }
]
const tableVData = {
  id: 1,
  name: '测试数据',
  status: '启用'
}
const lockValue = ref(false)
const lockUsers = ref([])

// 为 Button 组件提供 toolClick 函数
const toolClick = (props) => {
  console.log('Button clicked:', props)
  if (props.url) {
    window.open(props.url, '_blank')
  }
}
provide('toolClick', toolClick)

// 事件处理
const handleBtnSelect = (value) => {
  console.log('BtnSelect changed:', value)
}
const handleTimeRange = (timeStrings) => {
  if (timeStrings && timeStrings.length === 2) {
    timeStart.value = timeStrings[0]
    timeEnd.value = timeStrings[1]
  }
}
const handleSearcher = () => {
  console.log('Searcher search:', searcherArg.value)
}
</script>

<style scoped>
.portlet-body {
  padding: 20px;
}
</style>
