<script setup lang="ts">
import { highlightAll } from "@/interface/prettier";
import { useClipboard } from "@vueuse/core";
import { toRef, watch } from "vue";
import { Button, Icon } from "tdesign-vue-next";

const props = defineProps<{
  type: string;
  text: string;
}>();

watch(
  () => props.type,
  () => {
    highlightAll();
  },
  {
    immediate: true,
  }
);

const { copy } = useClipboard({
  source: toRef(props, "text"),
});
</script>
<template>
  <div class="component" v-show="text">
    <pre class="code"><code class="code"
                            :class="`language-${type}`">{{
        props.text
      }}</code></pre>
    <Button shape="square" class="copy-button" variant="text">
      <template #icon>
        <Icon name="attach" />
      </template>
    </Button>
  </div>
</template>

<style scoped>
.component {
  margin: 0.5rem 0;
  border-radius: 0.2rem;
  overflow: hidden;
  position: relative;
}

.code {
  margin: 0;
}

.copy-button {
  position: absolute;
  right: 0.5rem;
  top: 0.5rem;
}
</style>
