import { createRouter, createWebHistory } from "vue-router";

export const MainRoutes: {
  path: string;
  icon: string;
  title: string;
}[] = [
  {
    path: "/",
    icon: "app",
    title: "面板",
  },
  {
    path: "/config",
    icon: "control-platform",
    title: "配置",
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/",
      component: () => import("../pages/Home.vue"),
    },
    {
      path: "/config",
      component: () => import("../pages/Config.vue"),
    },
  ],
});
