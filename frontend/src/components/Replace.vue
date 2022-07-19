<script setup lang="ts">
import { useVModel } from "@vueuse/core";
import { Input, Button, Icon } from "tdesign-vue-next";
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
  <div class="component">
    <Input class="input" v-model="regText" />
    <Button variant="text" @click="handleReplace">
      <span> 替换为 </span>
      <template #icon>
        <Icon name="chevron-right-double" />
      </template>
    </Button>
    <Input class="input" v-model="replaceText" />
  </div>
</template>

<style scoped>
.component {
  display: flex;
  justify-content: flex-start;
  align-items: center;
  gap: 0.5rem;
  margin: 0.5rem 0;
}

.input {
  width: auto;
}

.component span {
  white-space: nowrap;
}
</style>
