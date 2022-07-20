<script setup lang="ts">
import { useVModel } from "@vueuse/core";
import { Space, Input } from "tdesign-vue-next";

const props = defineProps<{
  modelValue: string;
}>();
const emit = defineEmits(["update:modelValue"]);

const text = useVModel(props, "modelValue", emit);
</script>

<template>
  <Space size="small" align="center">
    <span> URL </span>
    <Input autofocus clearable class="input" v-model="text" />
  </Space>
</template>

<style scoped>
.input {
  width: 25rem;
}
</style>
