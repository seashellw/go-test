<script setup lang="ts">
import { useVModel } from "@vueuse/core";
import { Input } from "tdesign-vue-next";

const props = defineProps<{
  modelValue: string;
}>();
const emit = defineEmits(["update:modelValue"]);

const text = useVModel(props, "modelValue", emit);
</script>

<template>
  <div class="flex items-center gap-1">
    <span> URL： </span>
    <Input autofocus clearable class="input" v-model="text" />
  </div>
</template>

<style scoped>
.input {
  flex-grow: 1;
}
</style>
