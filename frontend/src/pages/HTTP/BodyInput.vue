<script setup lang="ts">
import { format } from "@/interface/prettier";
import { isStrJson } from "@/lib/json";
import { Eraser } from "@vicons/tabler";
import { useVModel } from "@vueuse/core";
import { NButton, NIcon, NInput, useMessage } from "naive-ui";

const props = defineProps<{
  modelValue: string;
}>();
const emit = defineEmits(["update:modelValue"]);

const text = useVModel(props, "modelValue", emit);

const message = useMessage();

const handleFormat = () => {
  if (!isStrJson(text.value)) {
    message.warning("文本不符合JSON格式，无法格式化");
    return;
  }
  text.value = format(text.value, "json");
};
</script>

<template>
  <div>
    <NInput
      v-model:value="text"
      type="textarea"
      placeholder="请输入请求体"
      autosize
      clearable
    />
    <div class="flex justify-end items-center mt-1 gap-2">
      <NButton secondary size="tiny" @click="handleFormat">
        <span> 格式化 </span>
        <template #icon>
          <NIcon size="1rem" class="mr-1">
            <Eraser />
          </NIcon>
        </template>
      </NButton>
    </div>
  </div>
</template>

<style scoped></style>
