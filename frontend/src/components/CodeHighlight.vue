<script setup lang="ts">
import { highlightAll } from "@/interface/prettier";
import { ClipboardPlus } from "@vicons/tabler";
import { Icon } from "@vicons/utils";
import { useClipboard } from "@vueuse/core";
import { Button, Tooltip } from "tdesign-vue-next";
import { toRefs, watch } from "vue";

const props = defineProps<{
  type: string;
  text: string;
}>();

watch(
  props,
  () => {
    highlightAll();
  },
  {
    immediate: true,
    flush: "post",
  }
);

const { copy } = useClipboard({
  source: toRefs(props).text,
});
</script>
<template>
  <div class="rounded my-2 overflow-hidden relative" v-show="text">
    <pre class="code"><code
      class="code"
      :class="`language-${type}`"
    >{{
        props.text
      }}</code></pre>
    <Tooltip content="复制">
      <Button
        shape="circle"
        class="copy-button"
        variant="text"
        @click="copy()"
      >
        <template #icon>
          <Icon size="1.2rem">
            <ClipboardPlus />
          </Icon>
        </template>
      </Button>
    </Tooltip>
  </div>
</template>

<style scoped>
.code {
  margin: 0;
  min-height: 4rem;
}

.copy-button {
  position: absolute;
  right: 0.3rem;
  top: 0.3rem;
}
</style>
