import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'
import {resolve} from 'path'
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
      vue(),
      AutoImport({
        resolvers: [ElementPlusResolver()],
      }),
      Components({
        resolvers: [ElementPlusResolver()],
      }),
    ],
  resolve: {
      alias: {
          '@':resolve(__dirname,'src')
      }
  },
  server: {
    // http://localhost:5173/api/login -> http://localhost:3001/login
    proxy: {
      '/api': {
        target: 'http://localhost:18967',
        changeOrigin: true,
        rewrite: path => path.replace(RegExp(`^/api`), '')
      },
    }
  },
})
