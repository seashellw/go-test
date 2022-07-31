<script setup lang="ts">
import { useVModel } from "@vueuse/core";
import { NSelect } from "naive-ui";

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
  <NSelect
    placeholder="请选择请求方法"
    class="select"
    v-model:value="type"
    :options="options"
  />
</template>

<style scoped>
.select {
  width: 10rem;
}
</style>
