<script setup lang="ts">
import Page from "@/components/Page.vue";
import { fetchHTTP, Request, Response } from "@/interface/fetch";
import Error from "@/pages/HTTP/Error.vue";
import HeaderTable from "@/pages/HTTP/HeaderTable.vue";
import { Send } from "@vicons/tabler";
import { NButton, NIcon, useMessage } from "naive-ui";
import { ref } from "vue";
import BodyInput from "./BodyInput.vue";
import BodyJsonPreview from "./BodyJsonPreview.vue";
import CookieTable from "./CookieTable.vue";
import Method from "./Method.vue";
import URL from "./URL.vue";

const req = ref<Request>({
  Url: "http://152.136.121.30:8080/api/user/logIn",
  Cookie: {},
  Data: "",
  Header: {},
  Method: "GET",
});

const res = ref<Response | undefined>(undefined);

const message = useMessage();

const fetch = async () => {
  req.value.Url = req.value.Url.trim();
  if (!req.value.Url.startsWith("http")) {
    message.warning("URL不正确");
    return;
  }
  if (!req.value.Url) {
    message.warning("请输入URL");
    return;
  }
  res.value = await fetchHTTP(req.value);
  if (res.value.Error) {
    message.error("请求失败");
    return;
  }
};
</script>

<template>
  <Page class="space-y-2">
    <URL v-model="req.Url" class="url" />
    <BodyInput v-model="req.Data" />
    <div class="flex flex-wrap gap-2">
      <Method v-model="req.Method" />
      <NButton @click="fetch" type="primary">
        发送
        <template #icon>
          <NIcon size="1.1rem" class="mr-1">
            <Send />
          </NIcon>
        </template>
      </NButton>
    </div>
    <Error :data="res?.Error" />
    <BodyJsonPreview :data="res?.Data" />
    <HeaderTable :header="res?.Header" />
    <CookieTable :cookie="res?.Cookie" />
  </Page>
</template>

<style scoped>
.url {
  width: 100%;
}
</style>
