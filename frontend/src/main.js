import { createApp } from "vue";
import App from "./App.vue";
import router from "./plugins/router";
import VueGoogleMaps from "@fawmi/vue-google-maps";

import "./scss/main.scss";

createApp(App)
  .use(router)
  .use(VueGoogleMaps, {
    load: {
      key: import.meta.env.VITE_GOOGLE_API_KEY,
    },
  })
  .mount("#app");
