<script lang="ts" setup>
import CodeHighlight from "@/components/CodeHighlight";
import Page from "@/components/Page.vue";
import { useConfig } from "@/config";
import { format } from "@/interface/prettier";
import {
NButton
} from "naive-ui";
import { computed } from "vue";
import { SysRebuild } from "wails/go/app/App.js";

const config = useConfig();
const configText = computed(() => {
  return format(JSON.stringify(config.config), "json");
});
</script>

<template>
  <Page>
    <div class="my-3">
      <p class="text-center text-lg"> 配置文件 </p>
      <CodeHighlight type="json" :text="configText" />
    </div>
    <div class="my-3 flex justify-center items-center gap-3">
      <NButton @click="SysRebuild"> 重构建 </NButton>
    </div>
  </Page>
</template>

<style scoped></style>
