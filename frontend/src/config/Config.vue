<script setup lang="ts">
import { fetchConfig } from "@/interface/config";
import { useDebounceFn } from "@vueuse/core";
import { watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import { useConfig } from ".";

const config = useConfig();
const router = useRouter();
const route = useRoute();
const res = await fetchConfig();
config.$state = res;

if (res.route && res.route !== route.path) {
  router.push(res.route);
}

watch(
  () => route.path,
  () => {
    config.route = route.path;
  }
);

const setConfig = useDebounceFn(fetchConfig, 3000);
watch(config, () => {
  const { route } = config;
  setConfig({ route });
});
</script>
<template></template>
