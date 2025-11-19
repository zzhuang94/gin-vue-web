<template>
  <div>
    <a-input-search
      v-if="search"
      v-model:value="searchValue"
      placeholder="请输入关键字后回车进行搜索"
      style="margin-bottom: 8px"
      @search="onSearch"
    />
    <a-tree
      :tree-data="filteredTreeData"
      :checkable :checkStrictly
      v-model:checkedKeys="ckeys"
      v-model:expandedKeys="expandedKeys"
    >
      <template #title="{ key, title, remark, icon, color, children, expanded, size, hover }">
        <component :is="hover ? Tooltip : 'span'" v-bind="hover ? { msg: hover, icon: false, placement: 'right' } : {}">
          <a-dropdown :trigger="['contextmenu']">
            <span
              @click="(e) => handleTitleClick(e, key, children)"
              @dblclick="(e) => handleDoubleClick(e, key, children)"
              >
              <span style="width: 1rem; margin-right: 0.5rem" :class="calcIconClass(icon, color, children, expanded, size)"></span>
              <span :style="{ fontSize: size + 'rem', fontWeight: isEmpty(children) ? 'normal' : 'bold' }">
                <template v-if="searchValue">
                  <span v-html="highlightText(title, searchValue)"></span>
                </template>
                <template v-else>
                  {{ title }}
                </template>
              </span>
              <span v-if="remark" v-html="remark"></span>
            </span>
            <template #overlay>
              <a-menu>
                <a-menu-item v-for="o, i in ops" :key="i" @click="emit('right-click', key, o.op)">
                  <i style="margin-right: 0.5rem" :class="`fa fa-${o.icon}`"></i>{{ o.title }}
                </a-menu-item>
              </a-menu>
            </template>
          </a-dropdown>
        </component>
      </template>
    </a-tree>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { isEmpty } from 'lodash'
import Tooltip from '@components/tooltip.vue'

const emit = defineEmits(['left-click', 'right-click', 'update:checkedKeys', 'double-click'])

const props = defineProps({
  data: Array,
  ops: Array,
  checkable: Boolean,
  checkStrictly: Boolean,
  checkedKeys: { type: Array, default: [] },
  search: { type: Boolean, default: false },
})

const searchValue = ref('')
const expandedKeys = ref(initKeys('expanded'))
const ckeys = ref([...props.checkedKeys])
const matchedKeys = ref(new Set()) // 存储匹配的节点key

watch(ckeys, (newVal) => { emit('update:checkedKeys', newVal) })
watch(() => props.checkedKeys, (newVal) => { ckeys.value = newVal || [] })

const treeData = computed(() => { return buildTree() })

// 纯过滤函数，不再修改展开状态
const filteredTreeData = computed(() => {
  if (!searchValue.value || !props.search) {
    return treeData.value
  }
  return filterTree(treeData.value, searchValue.value.toLowerCase())
})

// 纯过滤函数
function filterTree(nodes, keyword) {
  const result = []

  nodes.forEach(node => {
    const newNode = { ...node }
    const children = filterTree(node.children, keyword)

    const isMatched = node.title.toLowerCase().includes(keyword) || children.length > 0

    if (isMatched) {
      newNode.children = children
      result.push(newNode)

      // 记录匹配的节点
      if (node.title.toLowerCase().includes(keyword)) {
        matchedKeys.value.add(node.key)
      }
    }
  })
  return result
}

// 高亮匹配文本
function highlightText(text, keyword) {
  if (!keyword) return text
  const regex = new RegExp(keyword, 'gi')
  return text.replace(regex, match => `<span style="color: #f50; font-weight: bold">${match}</span>`)
}

// 搜索处理函数 - 单独处理展开逻辑
function onSearch() {
  if (searchValue.value) {
    const keyword = searchValue.value.toLowerCase()
    matchedKeys.value = new Set() // 重置匹配节点
    const nodesToExpand = new Set()

    // 先执行过滤，收集匹配的节点
    filterTree(treeData.value, keyword)

    // 然后查找需要展开的父节点
    const findParents = (nodes, targetKeys, parentKeys = []) => {
      nodes.forEach(node => {
        if (targetKeys.has(node.key)) {
          // 添加所有父节点到展开集合
          parentKeys.forEach(k => nodesToExpand.add(k))
        }
        findParents(node.children, targetKeys, [...parentKeys, node.key])
      })
    }

    findParents(treeData.value, matchedKeys.value)

    // 合并原有展开的节点和新需要展开的节点
    const finalExpanded = new Set([...expandedKeys.value, ...nodesToExpand])
    expandedKeys.value = Array.from(finalExpanded)
  } else {
    // 清空搜索时恢复默认展开状态
    expandedKeys.value = initKeys('expanded')
    matchedKeys.value = new Set()
  }
}

// 双击事件处理
const handleDoubleClick = (e, key, children) => {
  e.stopPropagation()
  emit('double-click', key, children)
}

// 以下是原有函数保持不变
function buildTree() {
  const tree = []
  const nodes = {}

  Object.keys(props.data).forEach(id => {
    const node = props.data[id]
    nodes[id] = {
      key: id,
      title: node.name,
      icon: node.icon,
      color: node.color,
      expanded: node.expanded,
      checked: node.checked,
      children: [],
    }
    if (node.hasOwnProperty('rank')) {
      nodes[id].rank = node.rank
    }
    if (node.hasOwnProperty('hover')) {
      nodes[id].hover = node.hover
    }
    if (node.hasOwnProperty('remark')) {
      nodes[id].remark = node.remark
    }
    if (node.hasOwnProperty('size')) {
      nodes[id].size = node.size
    } else {
      nodes[id].size = 1
    }
  })

  Object.keys(nodes).forEach(id => {
    const node = nodes[id]
    const parent = nodes[props.data[id].parent_id]
    if (parent) {
      parent.children.push(node)
    } else {
      tree.push(node)
    }
  })

  const sortTree = (nodes) => {
    if (!nodes || !nodes.length) {
      return
    }
    nodes.sort((a, b) => {
      if (a.hasOwnProperty('rank') && b.hasOwnProperty('rank')) {
        return a.rank < b.rank ? -1 : 1
      }
      if (a.children.length == 0 && b.children.length != 0) return 1
      if (a.children.length != 0 && b.children.length == 0) return -1
      return a.title.localeCompare(b.title)
    })
    nodes.forEach(node => sortTree(node.children))
  }

  sortTree(tree)
  return tree
}

function initKeys(type) {
  const ans = []
  for (let id in props.data) {
    if (props.data[id][type]) {
      ans.push(id)
    }
  }
  return ans
}

const calcIconClass = (icon, color, children, expanded) => {
  let ans = 'fa fa-'
  if (icon) {
    ans += icon
  } else if (!isEmpty(children)) {
    ans += 'folder'
    if (expanded) {
      ans += '-open'
    }
  } else {
    ans += 'file'
  }
  if (color) {
    ans += ' text-' + color
  }
  return ans
}

const handleTitleClick = (e, key, children) => {
  e.stopPropagation()

  if (isEmpty(children)) {
    emit('left-click', key)
    return
  }

  const newExpandedKeys = [...expandedKeys.value]
  const index = newExpandedKeys.indexOf(key)
  if (index > -1) {
    newExpandedKeys.splice(index, 1)
  } else {
    newExpandedKeys.push(key)
  }
  expandedKeys.value = newExpandedKeys
}
</script>

<style>
.ant-tree .ant-tree-checkbox {
  margin-inline-end: 4px !important;
  margin-block-start: 1px !important;
}
</style>
