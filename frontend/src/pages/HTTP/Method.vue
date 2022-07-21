<script setup lang="ts">
import { useVModel } from "@vueuse/core";
import { Space, Select } from "tdesign-vue-next";

const props = defineProps<{
  modelValue: string;
}>();
const emit = defineEmits(["update:modelValue"]);

const type = useVModel(props, "modelValue", emit);

const options: { label: string; value: string }[] = [
  "GET",
  "POST",
  "PUT",
  "DELETE",
].map((item) => ({
  label: item,
  value: item,
}));
</script>

<template>
  <div class="flex items-center gap-1">
    <span> 请求方法： </span>
    <Select class="select" v-model="type" :options="options" />
  </div>
</template>

<style scoped>
.select {
  width: 7rem;
}
</style>
