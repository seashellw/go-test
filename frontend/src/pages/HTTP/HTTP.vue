<script setup lang="ts">
import CodeHighlight from "@/components/CodeHighlight.vue";
import { useConfig } from "@/config";
import { fetchHTTP, Request, Response } from "@/interface/fetch";
import { format } from "@/interface/prettier";
import { Button, MessagePlugin, Space } from "tdesign-vue-next";
import { computed, ref } from "vue";
import Method from "./Method.vue";
import URL from "./URL.vue";

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
  let res = await fetchHTTP(request.value);
  console.log(res);
  response.value = res;
};
</script>

<template>
  <div class="mr-2 mt-2">
    <Space size="large" break-line>
      <URL v-model="request.Url" />
      <Method v-model="request.Method" />
      <Button @click="fetch"> 发送 </Button>
    </Space>
    <CodeHighlight :text="responseText" type="json" />
  </div>
</template>

<style scoped></style>
