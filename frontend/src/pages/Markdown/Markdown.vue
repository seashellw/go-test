<script setup lang="ts">
import { defineAsyncComponent } from "vue";
const Editor = defineAsyncComponent(() => import("./Editor.vue"));
</script>
<template>
  <Editor />
</template>

<style scoped></style>
