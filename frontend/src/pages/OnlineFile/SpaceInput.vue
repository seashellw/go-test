<script setup lang="ts">
import { watchThrottled } from "@vueuse/core";
import { Input } from "tdesign-vue-next";
import { computed } from "vue";
import { useFileList } from "./state";

const fileList = useFileList();

watchThrottled(
  () => fileList.space,
  () => {
    fileList.fetchList();
  },
  { throttle: 200 }
);

const inputStatus = computed(() => {
  if (!fileList.space) return "default";
  if (fileList.isSpaceError) return "warning";
  return undefined;
});
</script>
<template>
  <div class="flex items-center gap-1">
    <span class="whitespace-nowrap"> 空间代码： </span>
    <Input v-model="fileList.space" :status="inputStatus" />
  </div>
</template>

<style scoped></style>
