<script setup lang="ts">
import { ElButton } from "element-plus";
import { useAppStore } from "../store/app";
const appStore = useAppStore();
</script>
<template>
  <div>{{ appStore.count }}</div>
  <el-button @click="appStore.add">I am ElButton</el-button>
</template>
