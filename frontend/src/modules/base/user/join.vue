<template>
  <div class="portlet">
    <div class="portlet-head">
      <b style="font-size: 1.2rem"><i class="fa fa-user"></i>&nbsp;用户注册/登录</b>
    </div>
    <div class="portlet-body">

      <a-form :label-col="{span: 6}" :wrapper-col="{span: 12}">
        <a-form-item label="账号">
          <a-input v-model:value="username" placeholder="请输入账号（英文字母和数字组合）" />
        </a-form-item>

        <a-form-item label="密码">
          <a-input-password v-model:value="password" placeholder="请输入密码" />
        </a-form-item>

        <a-row>
          <a-col :span="6"></a-col>
          <a-col :span="18">
            <a-space>
              <button class="btn btn-primary" @click="handle('sign-up')" :disabled="submitting">
                注册
              </button>
              <button class="btn btn-success" @click="handle('log-in')" :disabled="submitting">
                登录
              </button>
            </a-space>
          </a-col>
        </a-row>
      </a-form>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import lib from '@libs/lib.ts'
import swal from '@libs/swal.ts'

const props = defineProps(['path'])
const submitting = ref(false)
const username = ref('')
const password = ref('')

const handle = async (action) => {
  if (!username.value || !password.value) {
    swal.error('错误', '请输入用户名和密码')
    return
  }

  submitting.value = true
  const ok = await lib.ajax(`/base/user/${action}`, {
    username: username.value,
    password: password.value
  })
  submitting.value = false

  if (ok) {
    window.location.href = props.path
  }
}
</script>
