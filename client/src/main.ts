import { createApp } from "vue";
import "./index.css";
import { createRouter, createWebHistory } from "vue-router";
import App from "./App.vue";
import { createPinia } from "pinia";

const router = createRouter({
  history: createWebHistory(),
  routes: [{ path: "/", component: () => import("./components/Home.vue") }],
});

const app = createApp(App);
app.use(router);
app.use(createPinia());
app.mount("#app");
