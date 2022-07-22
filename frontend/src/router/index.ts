import {
  Armchair,
  Command,
  Edit,
  PlugConnected,
} from "@vicons/tabler";
import { Component } from "vue";
import { createRouter, createWebHistory } from "vue-router";

interface RouteItem {
  path: string;
  component: Component;
  meta?: {
    // 标题
    title?: string;
    // 图标
    icon?: Component;
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
      title: "主屏",
      icon: Armchair,
      show: true,
    },
  },
  {
    path: "/text",
    component: () => import("@/pages/Text"),
    meta: {
      title: "文本处理",
      icon: Edit,
      show: true,
    },
  },
  {
    path: "/http",
    component: () => import("@/pages/HTTP"),
    meta: {
      title: "HTTP接口测试",
      icon: PlugConnected,
      show: true,
    },
  },
  {
    path: "/system",
    component: () => import("@/pages/System"),
    meta: {
      icon: Command,
      title: "系统",
      show: true,
    },
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes: Routes,
});
