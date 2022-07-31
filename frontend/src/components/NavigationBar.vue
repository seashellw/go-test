<script setup lang="ts">
import { Routes } from "@/router";
import { useVModel } from "@vueuse/core";
import { MenuOption, NIcon, NMenu } from "naive-ui";
import { computed, h } from "vue";
import { RouterLink, useRoute } from "vue-router";

const props = defineProps<{
  collapsed: boolean;
}>();
const emit = defineEmits(["update:collapsed"]);

const collapsed = useVModel(props, "collapsed", emit);

const route = useRoute();
const path = computed(() => {
  return route.matched[0]?.path || "/";
});

const menuOptions: MenuOption[] = Routes.filter(
  (item) => item.meta?.show
).map((item) => ({
  key: item.path,
  label: () =>
    h(
      RouterLink,
      { to: item.path },
      { default: () => item.meta?.title }
    ),
  icon: () =>
    h(NIcon, null, { default: () => h(item.meta?.icon, {}) }),
}));
</script>
<template>
  <NMenu
    :value="path"
    :collapsed="collapsed"
    :collapsed-width="58"
    :collapsed-icon-size="28"
    :options="menuOptions"
  />
</template>
<style scoped></style>
