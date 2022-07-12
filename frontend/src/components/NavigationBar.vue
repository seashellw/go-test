<script setup lang="ts">
import { Icon, Menu, MenuItem, Tooltip } from "tdesign-vue-next";
import { computed } from "vue";
import { useRoute, useRouter } from "vue-router";
import { Routes } from "../router";
const route = useRoute();
const router = useRouter();
const MainRoutes = Routes.filter((item) => item.meta?.show).map(
  (item) => ({
    path: item.path,
    title: item.meta?.title,
    icon: item.meta?.icon,
  })
);

const path = computed(
  () => `/${route.path.split("/").filter((v) => v)[0] || ""}`
);

const handleChange = (value: string) => {
  if (path.value !== value) router.push(value);
};
</script>
<template>
  <Menu theme="dark" :value="path" :collapsed="true">
    <Tooltip
      v-for="item in MainRoutes"
      :key="item.path"
      :content="item.title"
      placement="right"
    >
      <MenuItem :value="item.path" @click="handleChange(item.path)">
        <template #icon>
          <Icon :name="item.icon" />
        </template>
      </MenuItem>
    </Tooltip>
  </Menu>
</template>
