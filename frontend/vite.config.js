import { defineConfig } from "vite";

// https://vitejs.dev/config/
export default defineConfig({
  build: {
    outDir: "../backend/dist/static",
  },
  server: {
    host: true,
    port: 53172,
  },
});
