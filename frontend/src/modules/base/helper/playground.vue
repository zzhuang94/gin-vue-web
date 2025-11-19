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

<script>
export default {
  data() {
    return {
      token: '', // 随机生成的验证码
      userInput: '', // 用户输入的验证码内容
      canvasWidth: 100, // 验证码 Canvas 的宽度
      canvasHeight: 40, // 验证码 Canvas 的高度
      validationMessage: '', // 验证结果消息
    };
  },
  mounted() {
    this.generateCaptcha();
  },
  methods: {
    // 验证码字符串生成器
    generateRandomString(length) {
      const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz';
      let result = '';
      const charactersLength = characters.length;
      for (let i = 0; i < length; i++) {
        result += characters.charAt(Math.floor(Math.random() * charactersLength));
      }
      return result;
    },

    // 生成验证码
    generateCaptcha() {
      const canvas = this.$refs.captchaCanvas;
      const ctx = canvas.getContext('2d');

      // 重置画布大小
      ctx.clearRect(0, 0, this.canvasWidth, this.canvasHeight);
      ctx.fillStyle = '#f0f0f0';
      ctx.fillRect(0, 0, this.canvasWidth, this.canvasHeight);

      // 生成随机字符串
      this.token = this.generateRandomString(5).toUpperCase();

      // 绘制验证码（随机的颜色和位置）
      ctx.font = '25px Arial';
      for (let i = 0; i < this.token.length; i++) {
        ctx.fillStyle = this.randomColor();
        ctx.fillText(this.token[i], 10 + i * 18, 25 + Math.random() * 10);
      }

      // 添加一些干扰线
      for (let i = 0; i < 2; i++) {
        ctx.strokeStyle = this.randomColor();
        ctx.beginPath();
        ctx.moveTo(Math.random() * this.canvasWidth, Math.random() * this.canvasHeight);
        ctx.lineTo(Math.random() * this.canvasWidth, Math.random() * this.canvasHeight);
        ctx.stroke();
      }
    },

    // 随机颜色生成
    randomColor() {
      const r = Math.floor(Math.random() * 256);
      const g = Math.floor(Math.random() * 256);
      const b = Math.floor(Math.random() * 256);
      return `rgb(${r}, ${g}, ${b})`;
    },

    // 验证用户输入的验证码
    validateCaptcha() {
      if (this.userInput.toUpperCase() === this.token) {
        this.validationMessage = '验证成功 🎉';
      } else {
        this.validationMessage = '验证码错误，请重试！';
        this.generateCaptcha(); // 验证失败后重新生成验证码
      }
    },
  },
};
</script>
