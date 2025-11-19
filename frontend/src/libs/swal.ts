import { createApp, h } from 'vue'
import {  Modal, Input, notification, Spin } from 'ant-design-vue'

function succ(message: string, description: string = ''): void {
  notification.success(wrapNotiObj(message, description, 2, '#d4edda'))
}
function error(message: string, description: string = ''): void {
  notification.error(wrapNotiObj(message, description, 5, '#f8d7da'))
}
function info(message: string, description: string = ''): void {
  notification.info(wrapNotiObj(message, description, 3, '#d1ecf1'))
}
function warn(message: string, description: string = ''): void {
  notification.warning(wrapNotiObj(message, description, 3, '#fff3cd'))
}
function wrapNotiObj(message: string, description: string, duration: number, bgColor: string): any {
  return {
    message: h('div', {
      innerHTML: message,
      style: {
        maxHeight: '60vh',
        overflowY: 'auto',
        paddingRight: '8px',
        lineHeight: '1.5',
      }
    }),
    description,
    placement: 'top',
    duration,
    style: {
      top: `${window.innerHeight / 8}px`,
      whiteSpace: 'pre-line',
      border: '1px solid #e8e8e8',
      padding: '16px',
      backgroundColor: bgColor,
    }
  }
}

function confirm(title: string, content: string): any {
  return new Promise((resolve) => {
     Modal.confirm({
       title,
       content,
       style: { top: `${window.innerHeight / 8}px`, },
       onOk: () => resolve(true),
       onCancel: () => resolve(false),
       okText: '确定',
       cancelText: '取消',
     })
  })
}

function prompt(title: string = '请填写操作事由', placeholder: string = '不填则视为放弃操作'): any {
  return new Promise((resolve) => {
    let inputValue = ''

    Modal.confirm({
      title,
      width: 600,
      style: { top: `${window.innerHeight / 8}px`, },
      content: h(Input, {
        placeholder,
        onChange: (e) => (inputValue = e.target.value ?? ''),
        autofocus: true,
      }),
      okText: '确定',
      cancelText: '取消',
      onOk: () => resolve(inputValue || ''),
      onCancel: () => resolve(false),
    })
  })
}

function modal(title: string, content: string, width = '50%'): void {
  let dc: any = content
  try {
    const obj = JSON.parse(content)
    dc = h('pre', { style: 'font-size: 13px' }, JSON.stringify(obj, null, 2))
  } catch (e) {
    dc = h('p', { style: 'font-size: 13px' }, content)
  }
  Modal.info({
    title,
    width,
    closable: true,
    footer: null,
    content: dc,
  })
}

function htmlModal(title: string, content: string, width = '60%'): any {
  // 创建 Modal 实例
  const modal = Modal.info({
    width,
    closable: true,
    footer: null,
    title: typeof title === 'string' 
      ? h('div', { innerHTML: title }) 
      : title,
    content: typeof content === 'string' 
      ? h('div', { innerHTML: content })
      : content
  })

  // 手动触发更新（解决初始化渲染问题）
  setTimeout(() => {
    modal.update({
      title: typeof title === 'string' 
        ? h('div', { innerHTML: title }) 
        : title,
      content: typeof content === 'string' 
        ? h('div', { innerHTML: content })
        : content
    })
  }, 50)

  return modal
}

let loadingInstance: any = null

function startLoading (options: any = {}): void {
    if (!loadingInstance) {
    const container = document.createElement('div')
    container.id = 'simple-loading-container'
    container.style.position = 'fixed'
    container.style.top = '50%'
    container.style.left = '50%'
    container.style.transform = 'translate(-50%, -50%)'
    container.style.zIndex = '9999'
    document.body.appendChild(container)

    loadingInstance = createApp({
      render: () => h(Spin, {
        size: 'large',
        spinning: true,
        delay: options.delay || 100
      })
    }).mount(container)
  }
}

function stopLoading(): void {
  if (loadingInstance) {
    const container = document.getElementById('simple-loading-container')
    if (container) {
      container.remove()
    }
    loadingInstance = null
  }
}

export default {
  succ,
  error,
  info,
  warn,
  confirm,
  modal,
  prompt,
  htmlModal,
  startLoading,
  stopLoading,
}
