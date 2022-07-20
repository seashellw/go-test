<script setup lang="ts">
import CodeHighlight from "@/components/CodeHighlight.vue";
import { useConfig } from "@/config";
import { format } from "@/interface/prettier";
import { computed } from "vue";

const config = useConfig();
const configText = computed(() => {
  return format(JSON.stringify(config.config), "json");
});
</script>

<template>
  <div class="mr-2">
    <CodeHighlight type="json" :text="configText" />
  </div>
</template>

<style scoped></style>
