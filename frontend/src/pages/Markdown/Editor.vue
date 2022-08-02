<script setup lang="ts">
import "@fontsource/roboto/100.css";
import "@fontsource/roboto/300.css";
import "@fontsource/roboto/400.css";
import "@fontsource/roboto/500.css";
import "@fontsource/roboto/700.css";
import "@fontsource/roboto/900.css";
import "@material-design-icons/font";
import { defaultValueCtx, Editor, rootCtx } from "@milkdown/core";
import { block } from "@milkdown/plugin-block";
import { clipboard } from "@milkdown/plugin-clipboard";
import { cursor } from "@milkdown/plugin-cursor";
import { history } from "@milkdown/plugin-history";
import { listener, listenerCtx } from "@milkdown/plugin-listener";
import { menu } from "@milkdown/plugin-menu";
import { tooltip } from "@milkdown/plugin-tooltip";
import { commonmark } from "@milkdown/preset-commonmark";
import { nord } from "@milkdown/theme-nord";
import { useEditor, VueEditor } from "@milkdown/vue";
import { watchThrottled } from "@vueuse/core";
import { ref } from "vue";

const props = defineProps<{
  default?: string;
}>();

const emits = defineEmits<{
  (e: "change", value: string): void;
}>();

const value = ref("");

const { editor } = useEditor((root) =>
  Editor.make()
    .config((ctx) => {
      ctx.set(rootCtx, root);
      ctx.get(listenerCtx).markdownUpdated((_ctx, markdown) => {
        value.value = markdown;
      });
      let defaultValue = props.default || "# Hello World";
      ctx.set(defaultValueCtx, defaultValue);
    })
    .use(nord)
    .use(commonmark)
    .use(tooltip)
    .use(history)
    .use(cursor)
    .use(clipboard)
    .use(menu)
    .use(block)
    .use(listener)
);

watchThrottled(
  value,
  (value) => {
    emits("change", value);
  },
  { throttle: 1000 }
);
</script>
<template>
  <div class="p-2 pr-3 box">
    <VueEditor :editor="editor" />
  </div>
</template>

<style scoped>
.box {
  background-color: rgba(46, 52, 64, 1);
}
</style>
