<script setup lang="ts">
import { Progress } from "tdesign-vue-next";
import { computed } from "vue";
import { useFileList } from "./state";

const fileList = useFileList();

const list = computed(() => {
  return fileList.uploadList.map((item) => {
    let name = item.name;
    if (name.length > 14) {
      name = `${name.slice(0, 7)}...${name.slice(name.length - 7)}`;
    }
    return {
      ...item,
      name,
    };
  });
});
</script>
<template>
  <ul v-if="list.length">
    <li v-for="item in list">
      <Progress :percentage="item.percent" :status="item.status" />
    </li>
  </ul>
</template>

<style scoped></style>
