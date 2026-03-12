import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import type { Arg, Page, Data, Sort } from '@libs/frm'
import lib from '@libs/lib'

export type FetchFunc = () => Promise<void>

export function useFetch(opts: { arg: Arg, pageSize: number, sort?: Sort }) {
  const route = useRoute()
  const router = useRouter()

  const loading = ref<boolean>(true)
  const data = ref<Data[]>([])
  const arg = ref<Arg>(opts.arg)
  const page = ref<Page>({ curr: 1, size: opts.pageSize, total: 0 })
  const sort = ref<Sort>(opts.sort ?? { key: 'id', order: 'DESC' })

  const updateUrl = () => {
    router.replace({
      path: route.path,
      query: lib.mapToUriParams(arg.value),
    })
  }

  const fetch = async () => {
    loading.value = true
    try {
      const params: any = { arg: arg.value, page: page.value, sort: sort.value }
      const resp = await lib.curl('fetch', params)
      if (resp) {
        data.value = resp.data
        page.value = resp.page
      }
    } catch (error) {
      console.error(error)
    } finally {
      updateUrl()
      loading.value = false
    }
  }

  const reFetch = async () => {
    page.value.curr = 1
    await fetch()
  }

  return {
    loading,
    data,
    arg,
    page,
    sort,
    fetch,
    reFetch,
  }
}