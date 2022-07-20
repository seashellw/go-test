<script setup lang="ts">
import { useVModel } from "@vueuse/core";
import { Button, Icon, Input } from "tdesign-vue-next";
import { computed, ref } from "vue";
import { Space } from "tdesign-vue-next";

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
  <Space size="small" align="center">
    <Input v-model="regText" placeholder="正则表达式" />
    <Icon name="chevron-right-double" />
    <Input v-model="replaceText" placeholder="替换文本" />
    <Button @click="handleReplace"> 替换</Button>
  </Space>
</template>

<style scoped>
.input {
  width: auto;
}
</style>
