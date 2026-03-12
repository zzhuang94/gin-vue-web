
import { ref, shallowRef } from 'vue'
import lib from '@libs/lib'
import swal from '@libs/swal'

export type LoadModalFunc = (url: string, args?: any) => Promise<void>

export function useModal() {
  const mc = shallowRef(null)
  const mp = ref({})
  const modalUrl = ref('')

  const loadModal = async (url: string, args: any = {}): Promise<void> => {
    modalUrl.value = url
    try {
      swal.startLoading()
      const r = await lib.post(url, args)
      if (r.data.code === -1) {
        swal.warn("抱歉，您无权限操作！")
        return
      }
      const { page, props } = r.data
      mc.value = lib.loadComponent(page)
      mp.value = props
    } catch (error) {
      console.error('Error fetching component:', error)
    } finally {
      swal.stopLoading()
    }
  }
  
  const reloadModal = async (): Promise<void> => {
    try {
      swal.startLoading()
      const resp = await lib.post(modalUrl.value)
      mp.value = resp.data.props
    } catch (error) {
      console.error('Error fetching component:', error)
    } finally {
      swal.stopLoading()
    }
  }

  return {
    mc,
    mp,
    loadModal,
    reloadModal,
  }
}