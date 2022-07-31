<script setup lang="ts">
import CodeHighlight from "@/components/CodeHighlight";
import CollapseTransition from "@/components/CollapseTransition.vue";
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
  <CollapseTransition :show="bodyStr">
    <div class="py-4">
      <CodeHighlight :text="bodyStr" type="json" />
    </div>
  </CollapseTransition>
</template>

<style scoped></style>
