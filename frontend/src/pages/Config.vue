<script setup lang="ts">
import { useConfig } from "@/config";
import { Button } from "tdesign-vue-next";
const config = useConfig();
</script>
<template>
  <div>
    <Button class="margin"> config </Button>
  </div>
</template>
<style scoped>
.margin {
  margin: 1rem;
}
</style>
