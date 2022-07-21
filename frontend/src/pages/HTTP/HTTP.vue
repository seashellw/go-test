<script setup lang="ts">
import { useConfig } from "@/config";
import { fetchHTTP, Request, Response } from "@/interface/fetch";
import { format } from "@/interface/prettier";
import HeaderTable from "@/pages/HTTP/HeaderTable.vue";
import { Send } from "@vicons/tabler";
import { Icon } from "@vicons/utils";
import { Button, MessagePlugin } from "tdesign-vue-next";
import { computed, ref } from "vue";
import BodyInput from "./BodyInput.vue";
import BodyJsonPreview from "./BodyJsonPreview.vue";
import CookieTable from "./CookieTable.vue";
import Method from "./Method.vue";
import URL from "./URL.vue";

const config = useConfig();

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
  let data = await fetchHTTP(req.value);
  console.log(data);
  res.value = data;
};

const responseText = computed(() => {
  return format(JSON.stringify(res.value || {}), "json");
});
</script>

<template>
  <div class="my-2 mr-2">
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
    <div>
      <BodyJsonPreview :data="res?.Data" />
      <HeaderTable :header="res?.Header" />
      <CookieTable :cookie="res?.Cookie" />
    </div>
  </div>
</template>

<style scoped>
.url {
  width: 100%;
}
</style>
