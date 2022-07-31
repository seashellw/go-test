<script setup lang="ts">
import { Copy } from "@vicons/tabler";
import { useClipboard } from "@vueuse/core";
import hljs from "highlight.js";
import {
  NButton,
  NCode,
  NConfigProvider,
  NIcon,
  NTooltip,
} from "naive-ui";
import { toRefs } from "vue";

const props = defineProps<{
  type: string;
  text: string;
}>();

const { copy } = useClipboard({
  source: toRefs(props).text,
});
</script>
<template>
  <NConfigProvider :hljs="hljs">
    <div class="rounded overflow-hidden relative code" v-if="text">
      <NCode
        :code="props.text"
        :language="props.type"
        show-line-numbers
      />
      <NTooltip trigger="hover" placement="left">
        <span> 复制 </span>
        <template #trigger>
          <NButton
            circle
            secondary
            class="copy-button"
            @click="copy()"
          >
            <template #icon>
              <NIcon size="1.2rem">
                <Copy />
              </NIcon>
            </template>
          </NButton>
        </template>
      </NTooltip>
    </div>
  </NConfigProvider>
</template>

<style scoped>
.code {
  min-height: 4rem;
}

.copy-button {
  position: absolute;
  right: 0.2rem;
  top: 0.2rem;
}
</style>
