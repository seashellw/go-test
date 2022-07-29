<script setup lang="ts">
import Page from "@/components/Page.vue";
import { fetchHTTP, Request, Response } from "@/interface/fetch";
import Error from "@/pages/HTTP/Error.vue";
import HeaderTable from "@/pages/HTTP/HeaderTable.vue";
import { Send } from "@vicons/tabler";
import { Icon } from "@vicons/utils";
import { Button, MessagePlugin } from "tdesign-vue-next";
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

const fetch = async () => {
  req.value.Url = req.value.Url.trim();
  if (!req.value.Url.startsWith("http")) {
    await MessagePlugin.warning("URL不正确");
    return;
  }
  if (!req.value.Url) {
    await MessagePlugin.warning("请输入URL");
    return;
  }
  res.value = await fetchHTTP(req.value);
  if (res.value.Error) {
    await MessagePlugin.error("请求失败");
    return;
  }
};
</script>

<template>
  <Page class="py-2 pr-2 space-y-2">
    <div class="flex flex-wrap gap-2">
      <URL v-model="req.Url" class="url" />
      <BodyInput v-model="req.Data" />
      <Method v-model="req.Method" />
      <Button @click="fetch">
        发送
        <template #icon>
          <Icon size="1.1rem" class="mr-1">
            <Send />
          </Icon>
        </template>
      </Button>
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
