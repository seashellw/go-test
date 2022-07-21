<script setup lang="ts">
import { BaseTableColumns, Table } from "tdesign-vue-next";
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

const columns: BaseTableColumns = [
  {
    colKey: "key",
    title: "响应头",
  },
  {
    colKey: "value",
    title: "值",
  },
];
</script>

<template>
  <div v-if="data?.length" class="rounded overflow-hidden mt-2">
    <Table
      rowKey="key"
      tableLayout="auto"
      :columns="columns"
      :data="data"
      hover
      size="small"
      resizable
    />
  </div>
</template>

<style scoped>
:deep(.t-table__body) tr:nth-last-of-type(1) td {
  border-bottom: none;
}
</style>
