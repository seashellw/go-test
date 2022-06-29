import { createPinia } from "pinia";
import "tdesign-vue-next/es/style/index.css";
import { createApp } from "vue";
import App from "./App.vue";
import { router } from "./router";
import "./style/reset.css";

const app = createApp(App);
app.use(router);
app.use(createPinia());
app.mount("#app");
