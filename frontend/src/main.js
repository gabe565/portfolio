import "./plugins/plausible";
import { createApp } from "vue";
import App from "./App.vue";
import router from "./plugins/router";

import "./scss/main.scss";

createApp(App).use(router).mount("#app");
