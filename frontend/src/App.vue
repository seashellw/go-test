<script setup lang="ts">
import { Aside, Content, Layout } from "tdesign-vue-next";
import { Suspense } from "vue";
import { RouterView } from "vue-router";
import NavigationBar from "./components/NavigationBar.vue";
import Config from "./config/Config.vue";
</script>

<template>
  <Layout class="layout">
    <Aside width="4rem" class="aside">
      <NavigationBar />
    </Aside>
    <Content class="content">
      <RouterView />
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
}

.content {
  height: 100vh;
  width: calc(100vw - 4rem);
  overflow-y: auto;
  overflow-x: hidden;
}
</style>
