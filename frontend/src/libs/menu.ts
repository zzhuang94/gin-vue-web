import type { Menu } from '@libs/frm'
import type { LoadModalFunc } from '@libs/modal'
import type { FetchFunc } from '@libs/fetch'

import lib from '@libs/lib'

export function useMenu(loadModal: LoadModalFunc, fetch: FetchFunc) {
  const runMenu = async (m: Menu) => {
    if (m.type == 'modal') {
      loadModal(m.url)
    } else if (m.type == 'link') {
      lib.redirect(m.url)
    } else if (m.type == 'async') {
      const ok = await lib.confirmCurl(m.url)
      if (ok) {
        fetch()
      }
    }
  }
  return { runMenu }
}