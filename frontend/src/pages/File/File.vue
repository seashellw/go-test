<script setup lang="ts">
import Page from "@/components/Page.vue";
import FileList from "./FileList.vue";
import SpaceInput from "./SpaceInput.vue";
import UploadButton from "./UploadButton.vue";
import UploadProgress from "./UploadProgress.vue";
</script>
<template>
  <Page>
    <div class="pr-2 mt-2 flex flex-col gap-2">
      <SpaceInput />
      <FileList />
      <UploadButton />
      <UploadProgress />
    </div>
  </Page>
</template>

<style scoped></style>
