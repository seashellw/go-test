import { Component } from "vue";
import { createRouter, createWebHistory } from "vue-router";
import {
  Armchair,
  Cloud,
  Compass,
  DeviceDesktopAnalytics,
  Edit,
  Bible,
  Folders,
  PlugConnected,
} from "@vicons/tabler";

interface RouteItem {
  path: string;
  component: Component;
  meta?: {
    // 标题
    title?: string;
    // 图标
    icon?: any;
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
    path: "/localFile",
    component: () => import("@/pages/LocalFile"),
    meta: {
      title: "本地文件",
      icon: Folders,
      show: true,
    },
  },
  {
    path: "/onlineFile",
    component: () => import("@/pages/OnlineFile"),
    meta: {
      title: "在线文件",
      icon: Cloud,
      show: true,
    },
  },
  {
    path: "/markdown",
    component: () => import("@/pages/Markdown"),
    meta: {
      icon: Bible,
      title: "文档编辑器",
      show: true,
    },
    children: [
      {
        path: "",
        component: () => import("@/pages/Markdown/Home"),
        meta: {
          title: "目录",
        },
      },
      {
        path: "editor",
        component: () => import("@/pages/Markdown/Editor"),
        meta: {
          title: "编辑",
        },
      },
    ],
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
  {
    path: "/other",
    component: () => import("@/pages/Other.vue"),
    meta: {
      icon: Compass,
      title: "其他操作",
      show: true,
    },
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes: Routes,
});
