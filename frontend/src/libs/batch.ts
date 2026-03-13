
import { ref } from 'vue'

import type { FetchFunc } from '@libs/fetch'
import type { LoadModalFunc } from '@libs/modal'
import type { Menu } from '@libs/menu'

import swal from '@libs/swal'
import lib from '@libs/lib'

export function useBatch(loadModal: LoadModalFunc, fetch: FetchFunc) {
  const ids = ref<string[]>([])
  const runBatch = async (m: Menu) => {
    if (ids.value.length == 0) {
      swal.warn("注意！", "请至少选择一条数据");
      return
    }
    if (m.type == 'batch-edit') {
      const url = 'batch-edit?count=' + ids.value.length + '&ids=' + ids.value.join(',')
      loadModal(url)
    }
    if (m.type == 'batch-modal') {
      const url = (m.url) + '?count=' + ids.value.length + '&ids=' + ids.value.join(',')
      loadModal(url)
    }
    if (m.type == 'batch-delete') {
      const url = 'batch-delete?ids=' + ids.value.join(',')
      const ok = await lib.confirmCurl(url, '您将删除 ' + ids.value.length + ' 条数据')
      if (ok) {
        fetch()
      }
    }
  }
  return { ids, runBatch }
}