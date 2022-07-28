<script lang="ts" setup>
import Page from "@/components/Page.vue";
import { getDirPath, openDirDialog } from "@/interface/path";
import { Button, MessagePlugin, Input } from "tdesign-vue-next";
import FileList from "./FileList.vue";
import { useFileList } from "./state";

const fileList = useFileList();

const handleOpen = async () => {
  let res = await openDirDialog();
  fileList.dir = res;
};

const handleUp = () => {
  let dir = fileList.dir;
  if (!dir) {
    MessagePlugin.info("未打开目录");
    return;
  }
  dir = getDirPath(dir);
  if (dir) fileList.dir = dir;
  else MessagePlugin.info("已到达根目录");
};
</script>

<template>
  <Page>
    <div class="flex items-center gap-2">
      <Input
        readonly
        v-model="fileList.dir"
        @click="handleOpen"
        placeholder="打开本地目录"
      />
      <Button @click="handleUp" > 上层 </Button>
    </div>
    <FileList />
  </Page>
</template>
