<script setup lang="ts">
import CodeHighlight from "@/components/CodeHighlight.vue";
import { useConfig } from "@/config";
import { format } from "@/interface/prettier";
import { computed, ref } from "vue";
import { Input, Space, Button } from "tdesign-vue-next";
import { fetchHTTP, Request, Response } from "@/interface/fetch";
import { MessagePlugin } from "tdesign-vue-next";

const config = useConfig();

const request = ref<Request>({
  Url: "http://152.136.121.30:8080/api/tools/news",
  Cookie: {},
  Data: "",
  Header: {},
  Method: "GET",
});

const response = ref<Response>();

const responseText = computed(() => {
  return format(JSON.stringify(response.value || {}), "json");
});

const fetch = async () => {
  if (!request.value.Url) {
    await MessagePlugin.warning("请输入URL");
    return;
  }
  console.log(request.value);
  let res = await fetchHTTP(request.value);
  console.log(res);
  response.value = res;
};
</script>

<template>
  <div class="mr-2 mt-2">
    <Space size="small" align="center">
      <span> URL </span>
      <Input v-model="request.Url" />
    </Space>
    <Button @click="fetch"> 发送</Button>
    <CodeHighlight :text="responseText" type="json" />
  </div>
</template>

<style scoped></style>
