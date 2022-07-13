<script setup lang="ts">
import { fetchConfig } from "@/interface/config";
import { useDebounceFn } from "@vueuse/core";
import { watchEffect } from "vue";
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

watchEffect(() => {
  config.route = route.path;
});

const setConfig = useDebounceFn(fetchConfig, 3000);
watchEffect(() => {
  const { route } = config;
  setConfig({ route });
});
</script>
<template></template>
