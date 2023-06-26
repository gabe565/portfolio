import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import Icons from "unplugin-icons/vite";
import autoprefixer from "autoprefixer";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    Icons({
      compiler: "vue3",
      autoInstall: true,
    }),
  ],
  optimizeDeps: {
    include: ["@fawmi/vue-google-maps", "fast-deep-equal"],
  },
  css: {
    postcss: {
      plugins: [autoprefixer({})],
    },
  },
});
