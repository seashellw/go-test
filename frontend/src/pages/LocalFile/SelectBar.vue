<script lang="ts" setup>
import { getDirPath, openDirDialog } from "@/interface/path";
import { ArrowBigUpLine } from "@vicons/tabler";
import {
  NButton,
  NIcon,
  NInput,
  NTooltip,
  useMessage,
} from "naive-ui";
import { useFileList } from "./state";

const fileList = useFileList();

const message = useMessage();

const handleOpen = async () => {
  fileList.dir = await openDirDialog();
};

const handleUp = () => {
  let dir = fileList.dir;
  if (!dir) {
    message.info("未打开目录");
    return;
  }
  dir = getDirPath(dir);
  if (dir) fileList.dir = dir;
  else message.info("已到达根目录");
};
</script>

<template>
  <div class="flex items-center gap-2">
    <NInput
      readonly
      v-model:value="fileList.dir"
      @click="handleOpen"
      placeholder="打开本地目录"
      class="click-input"
    />
    <NTooltip trigger="hover">
      <template #trigger>
        <NButton @click="handleUp" type="info" secondary>
          <template #icon>
            <NIcon>
              <ArrowBigUpLine />
            </NIcon>
          </template>
        </NButton>
      </template>
      回到上层目录
    </NTooltip>
  </div>
</template>

<style scoped></style>
