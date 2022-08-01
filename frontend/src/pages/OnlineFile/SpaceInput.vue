<script setup lang="ts">
import { watchThrottled } from "@vueuse/core";
import { computed } from "vue";
import { useFileList } from "./state";
import { NInput } from "naive-ui";

const fileList = useFileList();

watchThrottled(
  () => fileList.space,
  () => {
    fileList.fetchList();
  },
  { throttle: 200 }
);

const inputStatus = computed(() => {
  if (!fileList.space) return undefined;
  if (fileList.isSpaceError) return "warning";
  return undefined;
});
</script>
<template>
  <NInput
    placeholder="请输入空间代码"
    v-model:value="fileList.space"
    :status="inputStatus"
  />
</template>

<style scoped></style>
