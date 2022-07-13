<script setup lang="ts">
import { Aside, Content, Layout, Header } from "tdesign-vue-next";
import { ref, Suspense, watchEffect } from "vue";
import { RouterView, useRoute } from "vue-router";
import NavigationBar from "./components/NavigationBar.vue";
import Config from "./config/Config.vue";

const route = useRoute();
const title = ref("");
watchEffect(() => {
  title.value = `${route.matched[0].meta.title}`;
});
</script>

<template>
  <Layout data-wails-drag>
    <Aside width="auto" class="left-nav-bar">
      <NavigationBar />
    </Aside>
    <Layout>
      <Header height="2rem" class="header">
        <span class="title">
          {{ title }}
        </span>
      </Header>
      <Content data-wails-no-drag class="content">
        <RouterView />
        <Suspense>
          <Config />
        </Suspense>
      </Content>
    </Layout>
  </Layout>
</template>

<style scoped>
.left-nav-bar {
  height: 100vh;
  overflow: auto;
}

.title {
  margin: 0 1rem;
}
.header {
  display: flex;
  align-items: center;
}
.content {
  height: calc(100vh - 2rem);
  overflow-y: auto;
  overflow-x: hidden;
}
</style>
