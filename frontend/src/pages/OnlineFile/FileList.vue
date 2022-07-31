<script setup lang="ts">
import { downloadFile, fetchDeleteFile } from "@/interface/cos";
import { TrashX } from "@vicons/tabler";
import { NIcon } from "naive-ui";
import {
  Button,
  DialogPlugin,
  List,
  ListItem,
  MessagePlugin,
  Tag,
  Tooltip,
} from "tdesign-vue-next";
import { useFileList } from "./state";

const fileList = useFileList();

const handleDownload = (item: { space: string; name: string }) => {
  downloadFile(item);
};

const handleDelete = (item: { space: string; name: string }) => {
  const dialog = DialogPlugin.confirm({
    header: `确定删除吗？`,
    body: `${item.name}`,
    destroyOnClose: true,
    onConfirm: async () => {
      let res = await fetchDeleteFile(item);
      if (res) {
        MessagePlugin.success("删除成功");
      } else {
        MessagePlugin.error("删除失败");
      }
      fileList.fetchList();
      dialog.destroy?.();
    },
  });
};

const handleClickSpace = (item: { space: string }) => {
  fileList.space = item.space;
};
</script>
<template>
  <List
    :split="true"
    size="small"
    class="list"
    v-if="fileList.list.length"
  >
    <ListItem v-for="item in fileList.list" :key="item.name">
      <p>
        <Tooltip :content="`空间：${item.space}`">
          <Tag
            class="mr-2 cursor-pointer"
            @click="handleClickSpace(item)"
          >
            {{ item.space }}
          </Tag>
        </Tooltip>
        <Tooltip :content="`下载：${item.name}`">
          <button
            @click="handleDownload(item)"
            class="hover:underline"
          >
            {{ item.name }}
          </button>
        </Tooltip>
      </p>
      <template #action>
        <Tooltip content="删除">
          <Button
            variant="text"
            shape="circle"
            @click="handleDelete(item)"
          >
            <template #icon>
              <NIcon size="1.2rem">
                <TrashX />
              </NIcon>
            </template>
          </Button>
        </Tooltip>
      </template>
    </ListItem>
  </List>
</template>

<style scoped>
.list {
  border-radius: 0.3rem;
}
.list li {
  height: 2.5rem;
}
.list li:last-child::after {
  display: none;
}
</style>
