<script setup lang="ts">
import CollapseTransition from "@/components/CollapseTransition.vue";
import { Circle, Folder } from "@vicons/tabler";
import { NIcon, NScrollbar } from "naive-ui";
import { watch } from "vue";
import { useFileList } from "./state";

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
    <NScrollbar style="max-height: 70vh">
      <ul
        class="divide-y divide-slate-500 divide-dashed rounded overflow-hidden"
      >
        <li
          class="item flex items-center px-1 py-2 gap-2"
          v-for="item in fileList.list"
          :key="item.name"
        >
          <p>
            <NIcon size="1.2rem" class="mr-2 align-middle">
              <Circle v-if="!item.isDir" class="text-blue-400" />
              <Folder v-if="item.isDir" class="text-yellow-400" />
            </NIcon>
            <button @click="handleClickName(item)">
              {{ item.name }}
            </button>
          </p>
        </li>
      </ul>
    </NScrollbar>
  </CollapseTransition>
</template>

<style scoped>
.item {
  transition: background-color 0.3s;
  background-color: rgba(255, 255, 255, 0);
}

.item:hover {
  background-color: rgba(255, 255, 255, 0.126);
}
</style>
