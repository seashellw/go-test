<script setup lang="ts">
import CodeHighlight from "@/components/CodeHighlight";
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
  <div class="py-4">
    <CodeHighlight :text="bodyStr" v-if="bodyStr" type="json" />
  </div>
</template>

<style scoped></style>
