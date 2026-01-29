<template>
  <a-card>
    <div style="display:flex; align-items: center;">
      <a-input
        v-model="userInput"
        placeholder="请输入验证码"
        style="width: 200px; margin-right: 20px;"
      />
      <canvas
        ref="captchaCanvas"
        :width="canvasWidth"
        :height="canvasHeight"
        style="border: 1px solid #d9d9d9; cursor: pointer;"
        @click="generateCaptcha"
      >
        您的浏览器不支持 Canvas，请升级！
      </canvas>
    </div>
    <p style="margin-top: 10px;">
      <a-button type="primary" @click="validateCaptcha">验证</a-button>
      <a-text v-if="validationMessage" style="margin-left: 16px; color: red;">{{ validationMessage }}</a-text>
    </p>
  </a-card>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

const token = ref('') // 随机生成的验证码
const userInput = ref('') // 用户输入的验证码内容
const canvasWidth = ref(100) // 验证码 Canvas 的宽度
const canvasHeight = ref(40) // 验证码 Canvas 的高度
const validationMessage = ref('') // 验证结果消息
const captchaCanvas = ref<HTMLCanvasElement | null>(null)

// 验证码字符串生成器
const generateRandomString = (length: number): string => {
  const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz';
  let result = '';
  const charactersLength = characters.length;
  for (let i = 0; i < length; i++) {
    result += characters.charAt(Math.floor(Math.random() * charactersLength));
  }
  return result;
}

// 随机颜色生成
const randomColor = (): string => {
  const r = Math.floor(Math.random() * 256);
  const g = Math.floor(Math.random() * 256);
  const b = Math.floor(Math.random() * 256);
  return `rgb(${r}, ${g}, ${b})`;
}

// 生成验证码
const generateCaptcha = () => {
  if (!captchaCanvas.value) return
  const canvas = captchaCanvas.value
  const ctx = canvas.getContext('2d');
  if (!ctx) return

  // 重置画布大小
  ctx.clearRect(0, 0, canvasWidth.value, canvasHeight.value);
  ctx.fillStyle = '#f0f0f0';
  ctx.fillRect(0, 0, canvasWidth.value, canvasHeight.value);

  // 生成随机字符串
  token.value = generateRandomString(5).toUpperCase();

  // 绘制验证码（随机的颜色和位置）
  ctx.font = '25px Arial';
  const tokenStr = token.value ?? ''
  for (let i = 0; i < tokenStr.length; i++) {
    ctx.fillStyle = randomColor();
    ctx.fillText(tokenStr[i] ?? '', 10 + i * 18, 25 + Math.random() * 10);
  }

  // 添加一些干扰线
  for (let i = 0; i < 2; i++) {
    ctx.strokeStyle = randomColor();
    ctx.beginPath();
    ctx.moveTo(Math.random() * canvasWidth.value, Math.random() * canvasHeight.value);
    ctx.lineTo(Math.random() * canvasWidth.value, Math.random() * canvasHeight.value);
    ctx.stroke();
  }
}

  // 验证用户输入的验证码
const validateCaptcha = () => {
  if (userInput.value.toUpperCase() === (token.value ?? '')) {
    validationMessage.value = '验证成功 🎉';
  } else {
    validationMessage.value = '验证码错误，请重试！';
    generateCaptcha(); // 验证失败后重新生成验证码
  }
}

onMounted(() => {
  generateCaptcha();
})
</script>
