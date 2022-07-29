<script lang="ts" setup>
import { getDirPath, openDirDialog } from "@/interface/path";
import { ArrowBigUpLine } from "@vicons/tabler";
import { Icon } from "@vicons/utils";
import {
  Button,
  Input,
  MessagePlugin,
  Tooltip,
} from "tdesign-vue-next";
import { useFileList } from "./state";

const fileList = useFileList();

const handleOpen = async () => {
  fileList.dir = await openDirDialog();
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
  <div class="flex items-center gap-2">
    <Input
      readonly
      v-model="fileList.dir"
      @click="handleOpen"
      placeholder="打开本地目录"
    />
    <Tooltip content="回到上层目录">
      <Button @click="handleUp" ghost>
        <template #icon>
          <Icon size="1.2rem">
            <ArrowBigUpLine />
          </Icon>
        </template>
      </Button>
    </Tooltip>
  </div>
</template>
