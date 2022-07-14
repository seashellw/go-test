<script setup lang="ts">
import { getConfig, setConfig } from "@/interface/config";
import { useDebounceFn } from "@vueuse/core";
import { watchEffect } from "vue";
import { useRoute, useRouter } from "vue-router";
import { useConfig } from ".";

const config = useConfig();
const router = useRouter();
const route = useRoute();

(async () => {
  const res = await getConfig();
  config.$state = res;
  if (res.route && res.route !== route.path) {
    router.push(res.route);
  }

  watchEffect(() => {
    config.route = route.path;
  });

  const set = useDebounceFn(setConfig, 1000);
  watchEffect(() => {
    set({
      route: config.route,
    });
  });
})();
</script>
<template></template>
