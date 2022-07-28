import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import path from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    // 配置路径别名
    alias: {
      "@": path.resolve(__dirname, "./src"),
      wails: path.resolve(__dirname, "./wailsjs"),
    },
  },
  server: {
    port: 3000,
    hmr: {
      port: 3001,
    },
  },
  build: {
    reportCompressedSize: false,
  },
});
