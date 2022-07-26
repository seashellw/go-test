<script setup lang="ts">
import { fetchUpload } from "@/interface/cos";
import { CloudUpload } from "@vicons/tabler";
import { Icon } from "@vicons/utils";
import { useThrottleFn } from "@vueuse/core";
import { Button, MessagePlugin } from "tdesign-vue-next";
import { reactive } from "vue";

import { UploadItem, useFileList } from "./state";

const fileList = useFileList();

const reFetch = useThrottleFn(() => fileList.fetchList(), 200);

const handleClickUpload = () => {
  if (fileList.isSpaceError) {
    MessagePlugin.warning("请输入合法空间代码");
    return;
  }
  let inputEl = document.createElement("input");
  inputEl.setAttribute("type", "file");
  inputEl.setAttribute("style", "visibility:hidden");
  inputEl.setAttribute("multiple", "multiple");
  inputEl.click();
  const listener = async () => {
    inputEl.removeEventListener("change", listener);
    if (!inputEl.files) return;
    const list = [...inputEl.files].map((item) => {
      const uploadItem: UploadItem = reactive({
        name: item.name,
        percent: 0,
        status: undefined,
      });
      fetchUpload({
        space: fileList.space,
        name: uploadItem.name,
        body: item,
        onProgress: (p) => {
          uploadItem.percent = p.percent * 100;
        },
      }).then((res) => {
        if (res.ok) {
          uploadItem.status = "success";
          MessagePlugin.success(`${uploadItem.name}上传成功`);
          setTimeout(() => {
            fileList.uploadList = fileList.uploadList.filter(
              (item) => item != uploadItem
            );
          }, 2000);
        } else {
          MessagePlugin.error(`${uploadItem.name}上传失败`);
          uploadItem.status = "error";
        }
        reFetch();
      });
      return uploadItem;
    });
    fileList.uploadList = [...fileList.uploadList, ...list];
  };
  inputEl.addEventListener("change", listener);
};
</script>
<template>
  <Button
    @click="handleClickUpload"
    :disabled="fileList.isSpaceError"
  >
    <span> 上传文件 </span>
    <template #icon>
      <Icon size="1.2rem" class="mr-2">
        <CloudUpload />
      </Icon>
    </template>
  </Button>
</template>

<style scoped></style>
