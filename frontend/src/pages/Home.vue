<script setup lang="ts">
import CodeHighlight from "@/components/CodeHighlight.vue";
import Replace from "@/components/Replace.vue";
import { format } from "@/interface/prettier";
import { Textarea } from "tdesign-vue-next";
import { computed, ref } from "vue";

const inputText = ref("");

let type = ref("json");

const formatRes = computed(() => {
  return format(inputText.value, type.value);
});
</script>
<template>
  <div class="component">
    <h1>主页</h1>
    <div class="text">
      <Replace v-model="inputText" />
      <Textarea v-model="inputText" />
      <CodeHighlight :text="formatRes" :type="type" />
    </div>
  </div>
</template>

<style scoped>
.component {
  width: 100%;
}

h1 {
  margin: 1rem;
}

.text {
  margin: 0.5rem;
}
</style>
