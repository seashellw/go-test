<script setup lang="ts">
import { fetchUpload } from "@/interface/cos";
import { CloudUpload } from "@vicons/tabler";
import { useThrottleFn } from "@vueuse/core";
import { NButton, NIcon, useMessage } from "naive-ui";
import { reactive } from "vue";

import { UploadItem, useFileList } from "./state";

const fileList = useFileList();

const reFetch = useThrottleFn(() => fileList.fetchList(), 200);
const message = useMessage();

const handleClickUpload = () => {
  if (fileList.isSpaceError) {
    message.warning("请输入合法空间代码");
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
          message.success(`${uploadItem.name}上传成功`);
          setTimeout(() => {
            fileList.uploadList = fileList.uploadList.filter(
              (item) => item != uploadItem
            );
          }, 2000);
        } else {
          message.error(`${uploadItem.name}上传失败`);
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
  <NButton
    @click="handleClickUpload"
    :disabled="fileList.isSpaceError"
    type="primary"
  >
    <span> 上传文件 </span>
    <template #icon>
      <NIcon size="1.2rem" class="mr-2">
        <CloudUpload />
      </NIcon>
    </template>
  </NButton>
</template>

<style scoped></style>
