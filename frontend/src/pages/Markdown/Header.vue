<script setup lang="ts">
import { ArrowBigLeftLine, NewSection } from "@vicons/tabler";
import { NButton, NIcon, NTooltip } from "naive-ui";
import { computed } from "vue";
import { useRoute, useRouter } from "vue-router";
import { useEditorState } from "./Editor.js";

const route = useRoute();
const router = useRouter();

const handleBack = () => {
  router.push("/markdown");
};
const title = computed(() => {
  return `${route.matched[1]?.meta.title}`;
});

const editor = useEditorState();
</script>
<template>
  <div class="mt-1 pr-3 flex items-center gap-1">
    <NTooltip trigger="hover">
      <template #trigger>
        <NButton quaternary @click="handleBack" size="small">
          <template #icon>
            <NIcon size="1.2rem"><ArrowBigLeftLine /></NIcon>
          </template>
        </NButton>
      </template>
      返回
    </NTooltip>

    <h3 class="flex-grow text-center">
      {{ title }}
    </h3>
    <NButton
      secondary
      @click="editor.header.handleClickNewButton"
      size="small"
      v-if="editor.header.showNewButton"
      type="primary"
    >
      <template #icon>
        <NIcon size="1.2rem"><NewSection /></NIcon>
      </template>
      新增
    </NButton>
  </div>
</template>

<style scoped></style>
