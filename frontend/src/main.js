import { createApp } from "vue";
import App from "@/App.vue";
import "@/plugins/plausible";
import router from "@/plugins/router";
import "@/scss/main.scss";

createApp(App).use(router).mount("#app");
