<script setup lang="ts">
import { useVModel } from "@vueuse/core";
import { NInput } from "naive-ui";

const props = defineProps<{
  modelValue: string;
}>();
const emit = defineEmits(["update:modelValue"]);

const text = useVModel(props, "modelValue", emit);
</script>

<template>
  <NInput clearable v-model:value="text" placeholder="请输入URL" />
</template>

<style scoped></style>
