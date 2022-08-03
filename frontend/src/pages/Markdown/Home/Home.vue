<script setup lang="ts">
import { useConfig } from "@/config/config";
import { DeviceDesktop } from "@vicons/tabler";
import {
  NBreadcrumb,
  NBreadcrumbItem,
  NButton,
  NIcon,
} from "naive-ui";
import { useEditorState } from "../editor";

import { readDir } from "@/interface/path";
import { computed, ref, watchEffect } from "vue";

const config = useConfig();

const editor = useEditorState();

const options = ref<string[]>([]);

const handleClickQuick = (path: string | undefined) => {
  console.log(path);
  if (!path) return;
  editor.dir = path;
};

watchEffect(async () => {
  if (!editor.dir) {
    options.value = [];
    return;
  }
  let fileList = readDir(editor.dir);
});

const pathItemList = computed(() => {
  let dir = editor.dir;
  if (!dir) return [];
  return dir.split(/[/\\]/).filter((item) => item);
});
</script>

<template>
  <div>
    <p class="my-1 mr-3 flex justify-end gap-1">
      <NButton
        secondary
        @click="handleClickQuick(config.path?.desktopPath)"
        size="tiny"
      >
        <template #icon>
          <NIcon><DeviceDesktop /></NIcon>
        </template>
        桌面
      </NButton>
    </p>
    <NBreadcrumb>
      <NBreadcrumbItem
        clickable
        v-for="(item, index) in pathItemList"
        :key="index"
      >
        {{ item }}
      </NBreadcrumbItem>
    </NBreadcrumb>
  </div>
</template>

<style scoped>
div ::v-deep(.n-breadcrumb-item .n-breadcrumb-item__separator) {
  margin: 0 2px;
}
</style>
