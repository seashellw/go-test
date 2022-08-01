<script setup lang="ts">
import CodeHighlight from "@/components/CodeHighlight";
import TextReplace from "@/pages/Text/TextReplace.vue";
import { format } from "@/interface/prettier";
import { ControlType } from "@/pages/Text/Control";
import ControlBar from "@/pages/Text/ControlBar.vue";
import { NInput } from "naive-ui";
import { computed, ref } from "vue";
import Page from "@/components/Page.vue";

const inputText = ref("");

const Action: {
  [key in ControlType]: {
    fn: (text: string) => string;
    highlightType: string;
  };
} = {
  [ControlType.JsonFormat]: {
    fn: (text) => {
      return format(text, "json");
    },
    highlightType: "json",
  },
};

const type = ref(ControlType.JsonFormat);

const transformRes = computed(() => {
  return {
    text: Action[type.value]?.fn(inputText.value),
    type: Action[type.value]?.highlightType,
  };
});
</script>
<template>
  <Page>
    <div class="space-y-2">
      <TextReplace v-model="inputText" />
      <ControlBar v-model="type" />
      <NInput
        v-model:value="inputText"
        type="textarea"
        :autosize="{
          minRows: 3,
        }"
        placeholder="请输入源文本"
      />
      <CodeHighlight
        :text="transformRes.text"
        :type="transformRes.type"
      />
    </div>
  </Page>
</template>

<style scoped></style>
