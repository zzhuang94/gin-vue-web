import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'
import Components from 'unplugin-vue-components/vite'
import { AntDesignVueResolver } from 'unplugin-vue-components/resolvers'

export default defineConfig({
  plugins: [
    vue(),
    Components({
      resolvers: [
        AntDesignVueResolver({ importStyle: 'less' }),
      ],
    }),
  ],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
      '@libs': path.resolve(__dirname, './src/libs'),
      '@styles': path.resolve(__dirname, './src/styles'),
      '@modules': path.resolve(__dirname, './src/modules'),
      '@templates': path.resolve(__dirname, './src/templates'),
      '@components': path.resolve(__dirname, './src/components'),
    }
  },
  server: {
    host: '0.0.0.0',
    allowedHosts: ['localhost', 'gin-vue.web.domain'],
    proxy: {
      '/web': {
        target: 'http://localhost:3000',
        changeOrigin: true,
      }
    }
  }
})
