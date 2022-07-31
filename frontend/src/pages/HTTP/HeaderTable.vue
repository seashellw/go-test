<script setup lang="ts">
import type { DataTableColumns } from "naive-ui";
import { NDataTable } from "naive-ui";
import CollapseTransition from "@/components/CollapseTransition.vue";
import { computed } from "vue";

const props = defineProps<{
  header?: {
    [key: string]: string;
  };
}>();

interface Row {
  key: string;
  value: string | undefined;
}

const data = computed<Row[]>(() => {
  if (!props.header) return [];
  return Object.keys(props.header).map((key) => ({
    key,
    value: props.header?.[key],
  }));
});

const columns: DataTableColumns<Row> = [
  {
    key: "key",
    title: "响应头",
  },
  {
    key: "value",
    title: "值",
  },
];
</script>

<template>
  <CollapseTransition :show="data?.length">
    <div class="rounded overflow-hidden">
      <NDataTable
        :bordered="false"
        :columns="columns"
        :data="data"
        size="small"
      />
    </div>
  </CollapseTransition>
</template>

<style scoped></style>
