import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import FontAwesomeIcon from "./plugins/fontawesome";
import VueGoogleMaps from "@fawmi/vue-google-maps";

import "@/sass/bootstrap.scss";
import "@/sass/main.scss";

createApp(App)
  .use(router)
  .use(VueGoogleMaps, {
    load: {
      key: import.meta.env.VITE_GOOGLE_API_KEY,
    },
  })
  .component("FontAwesomeIcon", FontAwesomeIcon)
  .mount("#app");
