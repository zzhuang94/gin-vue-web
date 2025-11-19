import { html } from 'diff2html';
import 'diff2html/bundles/css/diff2html.min.css';
import hljs from 'highlight.js';
import 'highlight.js/styles/github.css';


function genHtml(diff: string): string {
  if (!diff.trim()) {
    return ''
  }
  const div = document.createElement('div')
  div.innerHTML = html(diff, {drawFileList: false})
  div.querySelectorAll('pre code').forEach((el) => {
    hljs.highlightElement(el as HTMLElement)
  })
  return div.innerHTML
}

export default {
  genHtml,
}
