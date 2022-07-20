<script setup lang="ts">
import { watch, watchEffect } from "vue";
import { useRoute, useRouter } from "vue-router";
import { useConfig } from ".";

const config = useConfig();
const router = useRouter();
const route = useRoute();

const { route: oldRoute } = await config.fetchGetConfig();

if (oldRoute?.path && oldRoute?.path !== route.path) {
  await router.push(oldRoute?.path);
}

watchEffect(() => {
  config.route = {
    path: route.path,
  };
});

watch(config, () => {
  config.fetchSetConfig();
});
</script>
<template></template>
