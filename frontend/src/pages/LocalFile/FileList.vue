<script setup lang="ts">
import { Circle, Folder } from "@vicons/tabler";
import { NIcon } from "naive-ui";
import { Button, List, ListItem } from "tdesign-vue-next";
import { watch } from "vue";
import { useFileList } from "./state";
import CollapseTransition from "@/components/CollapseTransition.vue";

const fileList = useFileList();

watch(
  () => fileList.dir,
  () => {
    fileList.getList();
  }
);

const handleClickName = (item: { path: string; isDir: boolean }) => {
  if (item.isDir) fileList.dir = item.path;
};
</script>
<template>
  <CollapseTransition :show="fileList.list.length">
    <List
      :split="true"
      size="small"
      class="list"
    >
      <ListItem v-for="item in fileList.list" :key="item.name">
        <p>
          <NIcon size="1.2rem" class="mr-2 align-middle">
            <Circle v-if="!item.isDir" class="text-blue-400" />
            <Folder v-if="item.isDir" class="text-yellow-400" />
          </NIcon>
          <button @click="handleClickName(item)">
            {{ item.name }}
          </button>
        </p>
        <template #action> </template>
      </ListItem>
    </List>
  </CollapseTransition>
</template>

<style scoped>
.list {
  border-radius: 0.3rem;
  max-height: 70vh;
}

.list li:last-child::after {
  display: none;
}

.list li {
  height: 2.5rem;
}
</style>
