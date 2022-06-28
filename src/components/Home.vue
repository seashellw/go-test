<script setup lang="ts">
import { useAppStore } from "../store/app";
import { Button } from "tdesign-vue-next";
const appStore = useAppStore();
</script>
<template>
  <div>{{ appStore.count }}</div>
  <Button class="m-4" @click="appStore.add">I am ElButton</Button>
</template>
