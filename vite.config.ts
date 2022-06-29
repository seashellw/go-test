import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import { vanillaExtractPlugin } from "@vanilla-extract/vite-plugin";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue(), vanillaExtractPlugin()],
  server: {
    port: 3000,
    open: true,
    proxy: {
      "/api": "http://localhost:80",
    },
  },
});
