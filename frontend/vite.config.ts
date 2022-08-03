import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import { resolve } from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    // 配置路径别名
    alias: {
      "@": resolve("./src"),
      wails: resolve("./wailsjs"),
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
