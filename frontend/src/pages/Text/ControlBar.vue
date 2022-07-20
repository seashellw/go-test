<script setup lang="ts">
import { ControlType } from "@/pages/Text/Control";
import { useVModel } from "@vueuse/core";

import { Select, Space } from "tdesign-vue-next";

const props = defineProps<{
  modelValue: ControlType;
}>();
const emit = defineEmits(["update:modelValue"]);

const type = useVModel(props, "modelValue", emit);

const options: { label: string; value: ControlType }[] = [
  {
    label: "JSON格式化",
    value: ControlType.JsonFormat,
  },
];
</script>

<template>
  <Space size="small" align="center">
    <span> 处理模式 </span>
    <Select v-model="type" :options="options" />
  </Space>
</template>

<style scoped></style>
