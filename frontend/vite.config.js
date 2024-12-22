import vue from "@vitejs/plugin-vue";
import autoprefixer from "autoprefixer";
import Icons from "unplugin-icons/vite";
import { defineConfig } from "vite";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    Icons({
      compiler: "vue3",
      autoInstall: true,
    }),
  ],
  css: {
    postcss: {
      plugins: [autoprefixer({})],
    },
  },
});
