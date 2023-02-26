import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import VueGoogleMaps from "@fawmi/vue-google-maps";
import "./plugins/fontawesome";

import "@/sass/bootstrap.scss";
import "@/sass/main.scss";

createApp(App)
  .use(router)
  .use(VueGoogleMaps, {
    load: {
      key: import.meta.env.VITE_GOOGLE_API_KEY,
    },
  })
  .mount("#app");
