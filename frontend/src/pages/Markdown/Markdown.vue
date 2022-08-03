<script setup lang="ts">
import { KeepAlive } from "vue";
import { RouterView } from "vue-router";
import Header from "./Header.vue";
</script>
<template>
  <div>
    <Header />
    <RouterView v-slot="{ Component, route }">
      <KeepAlive>
        <component :is="Component" :key="route.path" />
      </KeepAlive>
    </RouterView>
  </div>
</template>

<style scoped></style>
