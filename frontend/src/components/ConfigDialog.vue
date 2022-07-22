<script lang="ts" setup>
import CodeHighlight from "@/components/CodeHighlight.vue";
import { useConfig } from "@/config";
import { format } from "@/interface/prettier";
import { Command, X } from "@vicons/tabler";
import { Icon } from "@vicons/utils";
import { Button, Dialog } from "tdesign-vue-next";
import { computed, ref } from "vue";
const visible = ref(false);

const handleSeeConfig = () => {
  visible.value = true;
};

const config = useConfig();
const configText = computed(() => {
  return format(JSON.stringify(config.config), "json");
});
</script>

<template>
  <Button variant="text" @click="handleSeeConfig">
    <span> 配置 </span>
    <template #icon>
      <Icon size="1.2rem" class="mr-1">
        <Command />
      </Icon>
    </template>
    <Dialog
      :footer="false"
      :header="false"
      width="80%"
      destroyOnClose
      :closeBtn="false"
      attach="body"
      v-model:visible="visible"
    >
      <Button
        @click="visible = false"
        variant="text"
        class="w-[1.5rem] h-[1.5rem] absolute top-1 right-1"
      >
        <template #icon>
          <Icon size="1.3rem">
            <X />
          </Icon>
        </template>
      </Button>
      <CodeHighlight type="json" :text="configText" />
    </Dialog>
  </Button>
</template>

<style scoped>
:deep(.t-dialog) {
  padding: 1rem;
}
</style>
