import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import tailwindcss from '@tailwindcss/vite'
import { fileURLToPath, URL } from 'node:url'
import Components from 'unplugin-vue-components/vite'
import pkg from './package.json'
// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    tailwindcss(),
    Components({
      // 默认会自动扫描 src/components 目录下的所有组件
      dirs: ['src/components'],
      // 支持的扩展名
      extensions: ['vue'],
      // 深度搜索子目录
      deep: true,
      // 自动生成 components.d.ts 以获得完整的 TS 类型提示
      dts: 'src/types/components.d.ts',
    })
  ],
  define: {
    "__APP_VERSION__": JSON.stringify(pkg.version)
  },
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server: {
    proxy: {
      "/api": {
        target: "http://localhost:8061"
      }
    }
  }
})
