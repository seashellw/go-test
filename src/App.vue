<script setup lang="ts">
import { Aside, Content, Layout } from "tdesign-vue-next";
import { RouterView } from "vue-router";
import { defineAsyncComponent } from "vue";
const NavigationBar = defineAsyncComponent(
  () => import("./components/NavigationBar.vue")
);
</script>

<template>
  <Layout>
    <Aside width="auto" class="h-screen">
      <NavigationBar />
    </Aside>
    <Content class="h-screen">
      <RouterView />
    </Content>
  </Layout>
</template>

<style scoped>

.h-screen {
  height: 100vh;
  overflow: auto;
}
</style>
