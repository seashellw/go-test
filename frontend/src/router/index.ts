import { createRouter, createWebHistory } from "vue-router";
import Home from "@/pages/Home.vue";
import Config from "@/pages/Config.vue";
import { Component } from "vue";

interface RouteItem {
  path: string;
  component: Component;
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
    component: Home,
    meta: {
      title: "面板",
      icon: "app",
      show: true,
    },
  },
  {
    path: "/config",
    component: Config,
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
