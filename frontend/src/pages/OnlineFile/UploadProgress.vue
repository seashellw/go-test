<script setup lang="ts">
import { computed } from "vue";
import { useFileList } from "./state";
import { NProgress } from "naive-ui";
import CollapseTransition from "@/components/CollapseTransition.vue";

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
  <CollapseTransition :show="list.length">
    <ul>
      <li v-for="item in list">
        <NProgress
          type="line"
          :percentage="item.percent"
          :status="item.status"
        />
      </li>
    </ul>
  </CollapseTransition>
</template>

<style scoped></style>
