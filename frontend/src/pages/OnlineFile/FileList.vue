<script setup lang="ts">
import CollapseTransition from "@/components/CollapseTransition.vue";
import { downloadFile, fetchDeleteFile } from "@/interface/cos";
import {
  NTag,
  useDialog,
  useMessage,
  NButton,
  NTooltip,
  NIcon,
} from "naive-ui";
import { computed } from "vue";
import { useFileList } from "./state";
import { TrashX, CloudDownload } from "@vicons/tabler";

const fileList = useFileList();

interface Row {
  name: string;
  space: string;
}

const handleDownload = (item: Row) => {
  downloadFile(item);
};

const handleClickSpace = (item: { space: string }) => {
  fileList.space = item.space;
};

const message = useMessage();
const dialog = useDialog();

const handleDelete = (item: Row) => {
  dialog.warning({
    title: `确定删除吗？`,
    content: `${item.name}`,
    positiveText: "确定",
    negativeText: "不确定",
    onPositiveClick: async () => {
      let res = await fetchDeleteFile(item);
      if (res) {
        message.success("删除成功");
      } else {
        message.error("删除失败");
      }
      fileList.fetchList();
    },
  });
};

const data = computed<Row[]>(() => fileList.list);
</script>
<template>
  <CollapseTransition :show="data.length">
    <ul
      class="list divide-y divide-slate-500 divide-dashed rounded overflow-hidden"
    >
      <li
        class="item flex items-center px-1 py-2 gap-2"
        v-for="row in data"
        :key="row.name"
      >
        <NTag
          class="tag flex-shrink-0"
          type="info"
          :bordered="false"
          @click="handleClickSpace(row)"
        >
          {{ row.space }}
        </NTag>
        <button
          class="flex-grow text-left truncate"
          :title="row.name"
        >
          {{ row.name }}
        </button>
        <div
          class="flex items-center gap-1 action-group flex-shrink-0"
        >
          <NTooltip trigger="hover">
            <template #trigger>
              <NButton
                secondary
                @click="handleDownload(row)"
                circle
                size="small"
              >
                <template #icon>
                  <NIcon>
                    <CloudDownload />
                  </NIcon>
                </template>
              </NButton>
            </template>
            下载
          </NTooltip>
          <NTooltip trigger="hover">
            <template #trigger>
              <NButton
                secondary
                @click="handleDelete(row)"
                circle
                size="small"
              >
                <template #icon>
                  <NIcon>
                    <TrashX />
                  </NIcon>
                </template>
              </NButton>
            </template>
            删除
          </NTooltip>
        </div>
      </li>
    </ul>
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

.tag {
  cursor: pointer;
}

.action-group {
  transition: opacity 0.3s;
  opacity: 0;
}

.item:hover .action-group {
  opacity: 1;
}
</style>
