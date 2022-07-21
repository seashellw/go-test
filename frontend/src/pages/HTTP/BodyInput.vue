<script setup lang="ts">
import { format } from "@/interface/prettier";
import { Eraser } from "@vicons/tabler";
import { Icon } from "@vicons/utils";
import { useVModel } from "@vueuse/core";
import { Button, MessagePlugin, Textarea } from "tdesign-vue-next";
import { isStrJson } from "./util";
const props = defineProps<{
  modelValue: string;
}>();
const emit = defineEmits(["update:modelValue"]);

const text = useVModel(props, "modelValue", emit);

const handleFormat = () => {
  if (!isStrJson(text.value)) {
    MessagePlugin.warning("文本不符合JSON格式，无法格式化");
    return;
  }
  text.value = format(text.value, "json");
};
</script>

<template>
  <div class="w-full">
    <div class="flex items-center mb-1 gap-2">
      <span> 请求体： </span>
      <Button size="small" ghost @click="handleFormat">
        <span> 格式化 </span>
        <template #icon>
          <Icon size="1rem" class="mr-1">
            <Eraser />
          </Icon>
        </template>
      </Button>
    </div>
    <Textarea
      autosize
      class="input"
      v-model="text"
      placeholder="请输入请求体"
    />
  </div>
</template>

<style scoped></style>
