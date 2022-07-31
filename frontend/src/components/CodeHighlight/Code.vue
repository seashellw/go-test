<script setup lang="ts">
import { defineAsyncComponent } from "vue";

const CodeHighlight = defineAsyncComponent(
  () => import("./CodeHighlight.vue")
);
const props = defineProps<{
  type: string;
  text: string;
}>();
</script>
<template>
  <CodeHighlight v-bind="props" />
</template>
