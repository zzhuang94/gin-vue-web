<template>
  <div style="display: flex; gap: 8px; flex-wrap: wrap;">
    <div class="portlet" style="flex: 0 0 400px;">
      <div class="portlet-head">
        <b style="font-size: 1.2rem"><i class="fa fa-image"></i> 修改头像</b>
      </div>
      <div class="portlet-body">
        <div style="display: flex; flex-direction: column; align-items: center; gap: 24px;">
          <div v-if="currentAvatar" style="text-align: center;">
            <p style="margin-bottom: 12px; color: #666;">当前头像</p>
            <img :src="currentAvatar" alt="当前头像" style="width: 120px; height: 120px; border-radius: 50%; object-fit: cover; border: 2px solid #d9d9d9;" />
          </div>

          <div v-if="!previewImage" style="text-align: center;">
            <a-upload
              :show-upload-list="false"
              :before-upload="beforeUpload"
              accept="image/*"
            >
              <a-button type="primary" size="large">
                <i class="fa fa-upload"></i> 选择图片
              </a-button>
            </a-upload>
          </div>

          <div v-if="previewImage" style="text-align: center; width: 100%;">
            <p style="margin-bottom: 12px; color: #666;">预览</p>
            <img :src="previewImage" alt="预览" style="width: 200px; height: 200px; border-radius: 50%; object-fit: cover; border: 2px solid #d9d9d9; margin-bottom: 16px;" />
            <div style="display: flex; justify-content: center; gap: 12px;">
              <a-button @click="cancelPreview">重新选择</a-button>
              <a-button type="primary" @click="confirmUpload" :loading="uploading">
                <i class="fa fa-check"></i> 确认上传
              </a-button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="portlet" style="flex: 1; min-width: 400px;">
      <div class="portlet-head">
        <b style="font-size: 1.2rem"><i class="fa fa-user-gear"></i> 个人设置</b>
      </div>
      <div class="portlet-body">
        <a-form :label-col="{span: 6}" :wrapper-col="{span: 12}">

          <template v-for="v in rules" :key="v.key">
            <a-form-item :required="v.required">
              <template #label>
                <Tooltip v-if="v.describe" placement="left" color="blue" :icon="false" :msg="v.describe">
                  <i class="fa fa-info-circle text-info"></i>&nbsp;{{ v.name }}
                </Tooltip>
                <span v-else>{{ v.name }}</span>
              </template>

              <a-select v-if="v.limit" v-model:value="fd[v.key]" show-search allow-clear
                :filterOption="lib.filterByLabel" :placeholder="`请选择${v.name}`"
                :mode="v.split_sep ? 'multiple' : 'undefined'">
                <a-select-option v-for="lv in v.limit" :key="lv.key || lv" :value="(lv.key !== undefined ? lv.key : lv)"
                  :label="lv.label || lv">
                  {{ lv.label || lv }}
                </a-select-option>
              </a-select>

              <a-input v-else v-model:value="fd[v.key]"
                :placeholder="`请输入${v.name}`" :readonly="v.readonly" />

            </a-form-item>
          </template>

          <a-row>
            <a-col :span="6"></a-col>
            <a-col :span="18">
              <a-space>
                <button class="btn btn-primary" @click="handleSave" :disabled="submitting">
                  保存
                </button>
              </a-space>
            </a-col>
          </a-row>
        </a-form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import lib from '@libs/lib.ts'
import swal from '@libs/swal.ts'
import Tooltip from '@components/tooltip.vue'

const props = defineProps(['user', 'rules'])
const submitting = ref(false)
const fd = reactive(props.user)

const currentAvatar = ref('')
const previewImage = ref('')
const croppedBase64 = ref('')
const uploading = ref(false)

const loadCurrentAvatar = async () => {
  currentAvatar.value = await lib.loadAvatar()
}

const beforeUpload = async (file) => {
  if (!file.type.startsWith('image/')) {
    swal.error('错误', '只能上传图片文件')
    return false
  }

  try {
    const reader = new FileReader()
    const imageData = await new Promise((resolve, reject) => {
      reader.onload = (e) => resolve(e.target?.result)
      reader.onerror = reject
      reader.readAsDataURL(file)
    })

    if (typeof imageData !== 'string') {
      throw new Error('无法读取图片')
    }

    const img = new Image()
    await new Promise((resolve, reject) => {
      img.onload = () => resolve()
      img.onerror = reject
      img.src = imageData
    })

    const size = Math.min(img.naturalWidth, img.naturalHeight)
    const startX = (img.naturalWidth - size) / 2
    const startY = (img.naturalHeight - size) / 2

    const canvas = document.createElement('canvas')
    canvas.width = size
    canvas.height = size
    const ctx = canvas.getContext('2d')
    
    if (!ctx) {
      throw new Error('无法创建绘图上下文')
    }

    ctx.drawImage(img, startX, startY, size, size, 0, 0, size, size)

    const base64 = canvas.toDataURL('image/jpeg', 0.9)
    previewImage.value = base64
    croppedBase64.value = base64.split(',')[1]
  } catch (error) {
    swal.error('错误', '处理图片失败：' + (error?.message || '未知错误'))
  }

  return false
}

const cancelPreview = () => {
  previewImage.value = ''
  croppedBase64.value = ''
}

const confirmUpload = async () => {
  if (!croppedBase64.value) return

  uploading.value = true
  const ok = await lib.ajax('upload-avatar', { image: croppedBase64.value })
  if (ok) {
    await loadCurrentAvatar()
    cancelPreview()
    window.setTimeout(() => window.location.reload(), 100)
  }
  uploading.value = false
}

const handleSave = async () => {
  submitting.value = true
  await lib.ajax(`save?id=${props.user.id}`, fd)
  submitting.value = false
}

onMounted(() => {
  loadCurrentAvatar()
})
</script>
  
