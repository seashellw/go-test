<script lang="ts" setup>
import { useVirtualList } from "@vueuse/core";
import { defineProps, ref, toRefs, watch } from "vue";

const props = defineProps<{
  list: string[];
}>();
const { list } = toRefs(props);

const terminal = ref<HTMLElement | null>(null);

const { containerProps, wrapperProps } = useVirtualList(list, {
  itemHeight: 16,
});

watch(
  props,
  () => {
    if (!terminal.value) return;
    terminal.value.scrollTop = terminal.value.scrollHeight;
  },
  {
    flush: "post",
  }
);
</script>

<template>
  <div
    v-if="list.length"
    v-bind="(containerProps as any)"
    class="terminal"
    ref="terminal"
  >
    <div v-bind="wrapperProps">
      <span class="line" v-for="(item, index) in list" :key="index">
        {{ item }}
      </span>
    </div>
  </div>
</template>

<style scoped>
.terminal {
  font-size: 12px;
  line-height: 16px;
  overflow-x: hidden;
  height: 160px;
  color: #00a870;
  background-color: rgba(26, 79, 112, 0.13);
  padding: 5px;
  border-radius: 3px;
}

.terminal::-webkit-scrollbar {
  display: none;
}

.line {
  display: block;
  height: 16px;
}
</style>
