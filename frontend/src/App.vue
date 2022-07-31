<script setup lang="ts">
import { Aside, Content, Layout } from "tdesign-vue-next";
import { Suspense } from "vue";
import { RouterView } from "vue-router";
import NavigationBar from "./components/NavigationBar.vue";
import Config from "./config";
import { useRouterTransition } from "./hooks/useRouterTransition";

const transitionName = useRouterTransition();
</script>

<template>
  <Layout class="layout">
    <Aside width="4rem" class="aside">
      <NavigationBar />
    </Aside>
    <Content class="content">
      <RouterView v-slot="{ Component, route }">
        <transition :name="transitionName">
          <KeepAlive>
            <component :is="Component" :key="route.path" />
          </KeepAlive>
        </transition>
      </RouterView>
      <Suspense>
        <Config />
      </Suspense>
    </Content>
  </Layout>
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
  overflow-y: auto;
  overflow-x: hidden;
}

.aside::-webkit-scrollbar {
  display: none;
}

.content {
  height: 100vh;
  width: calc(100vw - 4rem);
  overflow: hidden;
  position: relative;
}
</style>
