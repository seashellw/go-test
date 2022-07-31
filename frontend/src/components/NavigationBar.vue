<script setup lang="ts">
import { Routes } from "@/router";
import { Icon } from "@vicons/utils";
import {
  Button,
  Dropdown,
  DropdownMenu,
  Menu,
  MenuItem,
  Tooltip,
} from "tdesign-vue-next";
import { computed } from "vue";
import { useRoute, useRouter } from "vue-router";
import { Compass } from "@vicons/tabler";
import ConfigDialog from "./ConfigDialog.vue";

const route = useRoute();
const router = useRouter();
const MainRoutes = Routes.filter((item) => item.meta?.show).map(
  (item) => ({
    path: item.path,
    title: item.meta?.title,
    icon: item.meta?.icon,
  })
);

const path = computed(() => {
  return route.matched[0]?.path || "/";
});

const handleChange = (value: string) => {
  if (path.value !== value) router.push(value);
};
</script>
<template>
  <Menu class="menu" theme="dark" :value="path" :collapsed="true">
    <Tooltip
      v-for="item in MainRoutes"
      :key="item.path"
      :content="item.title"
      placement="right"
    >
      <MenuItem :value="item.path" @click="handleChange(item.path)">
        <template #icon>
          <Icon size="1.5rem" tag="button" color="white">
            <component :is="item.icon" />
          </Icon>
        </template>
      </MenuItem>
    </Tooltip>
    <template #operations>
      <Dropdown trigger="click" placement="right-bottom">
        <Button variant="text" class="operation-button">
          <template #icon>
            <Icon size="1.5rem">
              <Compass />
            </Icon>
          </template>
        </Button>
        <template #dropdown>
          <DropdownMenu class="drop-down">
            <ConfigDialog />
          </DropdownMenu>
        </template>
      </Dropdown>
    </template>
  </Menu>
</template>
<style scoped>
.menu {
  background-color: transparent;
}

.menu :deep(.t-menu__operations) {
  margin: 0;
  border: 0;
  padding: 8px;
}

.drop-down button {
  width: 100%;
}

.operation-button {
  width: 100%;
  height: 36px;
  padding: 0;
  margin: 0;
}
</style>
