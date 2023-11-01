import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import path from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "src"),
      "@views": path.resolve(__dirname, "src/views"),
      "@store": path.resolve(__dirname, "src/store"),
      "@components": path.resolve(__dirname, "src/components"),
    },
  },
  server: {
    host: "0.0.0.0",
    port: 3010,
  },
});
