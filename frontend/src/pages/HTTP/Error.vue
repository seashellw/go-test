<script setup lang="ts">
import CollapseTransition from "@/components/CollapseTransition.vue";

const props = defineProps<{
  data?: string;
}>();
</script>

<template>
  <CollapseTransition :show="props.data">
    <div class="border-red-400 border-2 p-2 rounded">
      {{ props.data }}
    </div>
  </CollapseTransition>
</template>

<style scoped></style>
