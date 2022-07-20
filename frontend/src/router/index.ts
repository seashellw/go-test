import { Component } from "vue";
import { createRouter, createWebHistory } from "vue-router";

interface RouteItem {
  path: string;
  component: Component;
  meta?: {
    // 标题
    title?: string;
    // 图标
    icon?: string;
    // 是否显示于侧边栏
    show?: boolean;
  };
  children?: RouteItem[];
}

export const Routes: RouteItem[] = [
  {
    path: "/",
    component: () => import("@/pages/Home.vue"),
    meta: {
      title: "面板",
      icon: "app",
      show: true,
    },
  },
  {
    path: "/text",
    component: () => import("@/pages/Text/Text.vue"),
    meta: {
      title: "文本处理",
      icon: "edit",
      show: true,
    },
  },
  {
    path: "/http",
    component: () => import("@/pages/HTTP/HTTP.vue"),
    meta: {
      title: "HTTP接口测试",
      icon: "internet",
      show: true,
    },
  },
  {
    path: "/config",
    component: () => import("@/pages/Config.vue"),
    meta: {
      icon: "save",
      title: "配置",
      show: true,
    },
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes: Routes,
});
