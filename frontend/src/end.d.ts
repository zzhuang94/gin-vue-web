declare module '*.vue' {
  import { DefineComponent } from 'vue'
  const component: DefineComponent<{}, {}, any>
  export default component
}

declare module 'xlsx'
declare module 'diff2html'
declare module 'highlight.js'
