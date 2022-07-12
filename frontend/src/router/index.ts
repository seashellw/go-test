import { createRouter, createWebHistory } from "vue-router";

interface RouteItem {
  path: string;
  component: () => Promise<any>;
  meta?: {
    title?: string;
    icon?: string;
    show?: boolean;
  };
  children?: RouteItem[];
}

export const Routes: RouteItem[] = [
  {
    path: "/",
    component: () => import("../pages/Home.vue"),
    meta: {
      title: "面板",
      icon: "app",
      show: true,
    },
  },
  {
    path: "/config",
    component: () => import("../pages/Config.vue"),
    meta: {
      icon: "control-platform",
      title: "配置",
      show: true,
    },
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes: Routes,
});
