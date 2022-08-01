<script setup lang="ts">
import { ControlType } from "@/pages/Text/Control";
import { useVModel } from "@vueuse/core";
import { NSelect } from "naive-ui";

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
  <div class="flex items-center gap-1">
    <span> 处理模式： </span>
    <NSelect class="select" v-model:value="type" :options="options" />
  </div>
</template>

<style scoped>
.select {
  width: 15rem;
}
</style>
