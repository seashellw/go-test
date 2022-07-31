<script setup lang="ts">
import { NCollapseTransition } from "naive-ui";

const props = defineProps<{
  show?: any;
}>();
</script>

<template>
  <NCollapseTransition :show="!!props.show" :appear="true">
    <slot />
  </NCollapseTransition>
</template>

<style scoped></style>
