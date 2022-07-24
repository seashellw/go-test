<script setup lang="ts">
import Page from "@/components/Page.vue";
import { fetchUpload } from "@/lib/cos";
import {
  Button,
  Progress,
  StatusEnum,
  List,
  ListItem,
} from "tdesign-vue-next";
import { ref } from "vue";

const status = ref<StatusEnum>("active");
const percent = ref(0);

const handleClickUpload = () => {
  var inputObj = document.createElement("input");
  inputObj.setAttribute("type", "file");
  inputObj.setAttribute("style", "visibility:hidden");
  inputObj.click();
  const listener = async () => {
    inputObj.removeEventListener("change", listener);
    const file = inputObj.files?.[0];
    if (!file) return;
    let res = await fetchUpload({
      space: "test",
      name: file.name,
      body: file,
      onProgress: (p) => {
        percent.value = p.percent * 100;
      },
    });
    if (res.ok) {
      status.value = "success";
    } else {
      status.value = "error";
    }
  };
  inputObj.addEventListener("change", listener);
};
</script>
<template>
  <Page>
    <div class="pr-2 mt-2 flex flex-col gap-2">
      <Button @click="handleClickUpload">上传文件</Button>
      <Progress :percentage="percent" :status="status" />

      <List :split="true" size="small">
        <ListItem>列表内容的描述性文字</ListItem>
      </List>
    </div>
  </Page>
</template>

<style scoped></style>
