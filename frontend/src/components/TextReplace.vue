<script setup lang="ts">
import { ChevronsRight, Replace } from "@vicons/tabler";
import { NIcon } from "naive-ui";
import { useVModel } from "@vueuse/core";
import { Button, Input } from "tdesign-vue-next";
import { computed, ref } from "vue";

const props = defineProps<{
  modelValue: string;
}>();
const emit = defineEmits(["update:modelValue"]);

const text = useVModel(props, "modelValue", emit);

const regText = ref("");
const replaceText = ref("");

const regExp = computed(() => {
  return new RegExp(regText.value, "g");
});

const handleReplace = () => {
  text.value = text.value.replace(
    regExp.value,
    replaceText.value || ""
  );
};
</script>
<template>
  <div class="flex items-center gap-2">
    <Input v-model="regText" placeholder="正则表达式" />
    <NIcon size="1.3rem">
      <ChevronsRight />
    </NIcon>
    <Input v-model="replaceText" placeholder="替换文本" />
    <Button @click="handleReplace" class="button">
      替换
      <template #icon>
        <NIcon size="1rem" class="mr-1">
          <Replace />
        </NIcon>
      </template>
    </Button>
  </div>
</template>

<style scoped>
.button {
  width: 8rem;
}
</style>
