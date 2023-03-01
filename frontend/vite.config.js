import { fileURLToPath, URL } from "node:url";
import AutoImport from "unplugin-auto-import/vite";
import Components from "unplugin-vue-components/vite";
import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    AutoImport({
      dts: true,
      imports: ["vue", "vue-router"],
      eslintrc: {
        enabled: true,
      },
    }),
    Components({
      dts: true,
      resolvers: [
        (name) => {
          if (name.match(/^FontAwesome/)) {
            return { name, from: "@fortawesome/vue-fontawesome" };
          }
        },
      ],
      types: [
        {
          from: "vue-router",
          names: ["RouterLink", "RouterView"],
        },
      ],
    }),
  ],
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
    },
  },
});
