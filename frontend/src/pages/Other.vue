<script lang="ts" setup>
import CodeHighlight from "@/components/CodeHighlight";
import { useConfig } from "@/config";
import { format } from "@/interface/prettier";
import { computed } from "vue";
import Page from "@/components/Page.vue";

const config = useConfig();
const configText = computed(() => {
  return format(JSON.stringify(config.config), "json");
});
</script>

<template>
  <Page>
    <div class="my-3">
      <span> 配置文件： </span>
      <CodeHighlight type="json" :text="configText" />
    </div>
  </Page>
</template>

<style scoped></style>
