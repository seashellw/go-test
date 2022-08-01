<script setup lang="ts">
import {
  darkTheme,
  dateZhCN,
  NConfigProvider,
  NLayout,
  NLayoutContent,
  NDialogProvider,
  NLayoutSider,
  NMessageProvider,
  zhCN,
} from "naive-ui";
import { ref, Suspense } from "vue";
import { RouterView } from "vue-router";
import NavigationBar from "./components/NavigationBar.vue";
import { useRouterTransition } from "./hooks/useRouterTransition";
import Config from "./config";

const transitionName = useRouterTransition();

const collapsed = ref(true);
</script>

<template>
  <NConfigProvider
    :locale="zhCN"
    :date-locale="dateZhCN"
    :theme="darkTheme"
  >
    <NMessageProvider>
      <NDialogProvider>
        <NLayout has-sider class="layout">
          <NLayoutSider
            collapse-mode="width"
            :collapsed-width="58"
            :width="240"
            class="aside"
            :native-scrollbar="false"
            :collapsed="collapsed"
            show-trigger
            @collapse="collapsed = true"
            @expand="collapsed = false"
          >
            <NavigationBar v-model:collapsed="collapsed" />
          </NLayoutSider>
          <NLayoutContent :native-scrollbar="false" class="content">
            <RouterView v-slot="{ Component, route }">
              <transition :name="transitionName">
                <KeepAlive>
                  <component :is="Component" :key="route.path" />
                </KeepAlive>
              </transition>
            </RouterView>
            <Suspense><Config /></Suspense>
          </NLayoutContent>
        </NLayout>
      </NDialogProvider>
    </NMessageProvider>
  </NConfigProvider>
</template>

<style scoped>
.layout {
  height: 100vh;
  width: 100vw;
  overflow: hidden;
  background-color: rgba(0, 0, 0, 0.7);
}

.aside {
  height: 100vh;
  background-color: transparent;
}

.content {
  height: 100vh;
  position: relative;
  background-color: transparent;
  border-top-left-radius: 3px;
}

.aside ::v-deep(.n-layout-toggle-button) {
  background-color: rgba(255, 255, 255, 0.126);
}
</style>
