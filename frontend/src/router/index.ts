import {
  Armchair,
  DeviceDesktopAnalytics,
  Edit,
  Folders,
  Cloud,
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
    path: "/LocalFile",
    component: () => import("@/pages/LocalFile"),
    meta: {
      title: "本地文件",
      icon: Folders,
      show: true,
    },
  },
  {
    path: "/OnlineFile",
    component: () => import("@/pages/OnlineFile"),
    meta: {
      title: "在线文件",
      icon: Cloud,
      show: true,
    },
  },
  {
    path: "/system",
    component: () => import("@/pages/System"),
    meta: {
      icon: DeviceDesktopAnalytics,
      title: "计算机",
      show: true,
    },
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes: Routes,
});
