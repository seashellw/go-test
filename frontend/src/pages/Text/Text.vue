<script setup lang="ts">
import CodeHighlight from "@/components/CodeHighlight.vue";
import Replace from "@/components/TextReplace.vue";
import { format } from "@/interface/prettier";
import { ControlType } from "@/pages/Text/Control";
import ControlBar from "@/pages/Text/ControlBar.vue";
import { Space, Textarea } from "tdesign-vue-next";
import { computed, ref } from "vue";

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
  <div>
    <div class="mr-2 mt-2 flex flex-col gap-2">
      <Replace v-model="inputText" />
      <ControlBar v-model="type" />
      <Textarea v-model="inputText" placeholder="源文本" />
      <CodeHighlight
        :text="transformRes.text"
        :type="transformRes.type"
      />
    </div>
  </div>
</template>

<style scoped></style>
