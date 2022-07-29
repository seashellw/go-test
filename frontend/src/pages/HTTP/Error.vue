<script setup lang="ts">

const props = defineProps<{
  data?: string;
}>();
</script>

<template>
  <div class="border-red-500 border-2 p-2 rounded" v-if="data">
    {{ data }}
  </div>
</template>

<style scoped></style>
