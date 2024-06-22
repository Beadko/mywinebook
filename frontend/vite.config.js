import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { fileURLToPath } from 'url'
import path from 'path'
import tailwindcss from "tailwindcss"


const filename = fileURLToPath(import.meta.url)
const pathSegments = path.dirname(filename)

export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': path.resolve(pathSegments, './src'),
    }
  },
  css: {
    postcss: {
      plugins: [tailwindcss()],
    },
  }
})
