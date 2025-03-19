import { createApp } from "vue";
import App from "@/App.vue";
import "@/css/main.css";
import "@/plugins/plausible";
import router from "@/plugins/router";

createApp(App).use(router).mount("#app");
