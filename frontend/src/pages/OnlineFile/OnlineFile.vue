<script setup lang="ts">
import Page from "@/components/Page.vue";
import ControlBar from "./ControlBar.vue";
import FileList from "./FileList.vue";
import UploadProgress from "./UploadProgress.vue";
</script>
<template>
  <Page>
    <div class="space-y-2">
      <ControlBar />
      <UploadProgress />
      <FileList />
    </div>
  </Page>
</template>

<style scoped></style>
