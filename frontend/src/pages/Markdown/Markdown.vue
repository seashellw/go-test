<script setup lang="ts">
import { defineAsyncComponent } from "vue";
import {
  FileGetCurrentPath,
  FileGetDesktopPath,
} from "wails/go/app/App";
const Editor = defineAsyncComponent(() => import("./Editor.vue"));
(async () => {
  console.log(await FileGetCurrentPath());
  console.log(await FileGetDesktopPath());
})();
</script>
<template>
  <Editor />
</template>

<style scoped></style>
