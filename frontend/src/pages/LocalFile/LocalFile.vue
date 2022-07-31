<script lang="ts" setup>
import Output from "@/components/Output.vue";
import Page from "@/components/Page.vue";
import { ref } from "vue";
import FileList from "./FileList.vue";
import SelectBar from "./SelectBar.vue";

const list = ref(["1"]);

const add = () => {
  list.value.push(
    `${parseInt(list.value[list.value.length - 1]) + 1}`
  );
};
</script>

<template>
  <Page class="space-y-2">
    <SelectBar />
    <FileList />

    <button @click="add">add</button>
    <Output :list="list" />
  </Page>
</template>
