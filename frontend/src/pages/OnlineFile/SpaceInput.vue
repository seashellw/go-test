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
  if (fileList.isSpaceError) return "warning";
  return undefined;
});
</script>
<template>
  <div class="">
    <Input
      label="空间代码："
      v-model="fileList.space"
      :status="inputStatus"
    />
  </div>
</template>

<style scoped></style>
