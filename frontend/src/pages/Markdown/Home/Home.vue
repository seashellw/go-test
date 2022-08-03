<script setup lang="ts">
import { useConfig } from "@/config/config";
import {
  NButton,
  NIcon,
  NTooltip,
  NAutoComplete,
  NBreadcrumb,
  NBreadcrumbItem,
} from "naive-ui";
import { ArrowBigLeftLine, DeviceDesktop } from "@vicons/tabler";
import { useEditorState } from "../editor";

import { computed, ref, watchEffect, watch } from "vue";
import { readDir } from "@/interface/path";

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

watch(
  () => config.config,
  (val) => {
    console.log(val);
  }
);

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
        v-for="(item, index) in pathItemList"
        :key="index"
      >
        {{ item }}
      </NBreadcrumbItem>
    </NBreadcrumb>
  </div>
</template>

<style scoped></style>
