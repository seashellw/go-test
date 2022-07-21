<script setup lang="ts">
import CodeHighlight from "@/components/CodeHighlight.vue";
import { format } from "@/interface/prettier";
import { isStrJson } from "@/lib/json";
import { computed } from "vue";

const props = defineProps<{
  data?: string;
}>();

const bodyStr = computed(() => {
  let { data } = props;
  if (!data) return "";
  if (!isStrJson(data)) return data;
  return format(data, "json");
});
</script>

<template>
  <CodeHighlight v-if="bodyStr" :text="bodyStr" type="json" />
</template>

<style scoped></style>
