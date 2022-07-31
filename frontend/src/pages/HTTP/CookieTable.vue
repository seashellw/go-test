<script setup lang="ts">
import type { DataTableColumns } from "naive-ui";
import { NDataTable } from "naive-ui";
import { computed } from "vue";
import CollapseTransition from "@/components/CollapseTransition.vue";

const props = defineProps<{
  cookie?: {
    [key: string]: string;
  };
}>();

interface Row {
  key: string;
  value: string | undefined;
}

const data = computed<Row[]>(() => {
  if (!props.cookie) return [];
  return Object.keys(props.cookie).map((key) => ({
    key,
    value: props.cookie?.[key],
  }));
});

const columns: DataTableColumns<Row> = [
  {
    key: "key",
    title: "SetCookie",
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
