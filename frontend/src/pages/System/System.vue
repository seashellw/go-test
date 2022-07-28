<script setup lang="ts">
import Page from "@/components/Page.vue";
import { defineAsyncComponent } from "vue";
const AreaCPU = defineAsyncComponent(() => import("./AreaCPU.vue"));
</script>

<template>
  <Page>
    <AreaCPU class="mt-2" />
  </Page>
</template>

<style scoped></style>
