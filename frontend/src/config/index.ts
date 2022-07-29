import { useThrottleFn } from "@vueuse/core";
import { defineStore } from "pinia";
import { watch, watchEffect } from "vue";
import { useRoute, useRouter } from "vue-router";
import { ConfigGet, ConfigSet } from "wails/go/app/App";

const setConfig = useThrottleFn(ConfigSet, 1000);

export interface Config {
  route?: {
    path: string;
  };
}

const init: () => Config = () => ({});

export const useConfig = defineStore("globalConfig", {
  state: init,
  getters: {
    config: (state) => ({
      route: state.route,
    }),
  },
  actions: {
    fetchSetConfig() {
      setConfig(JSON.stringify(this.config));
    },
    async fetchGetConfig() {
      let res = await ConfigGet();
      let config: Config;
      try {
        config = JSON.parse(res || "{}");
      } catch {
        config = {};
      }
      this.$state = config;
      return config;
    },
  },
});

export const useConfigInit = async () => {
  const config = useConfig();
  const router = useRouter();
  const route = useRoute();

  const { route: oldRoute } = await config.fetchGetConfig();

  if (oldRoute?.path && oldRoute?.path !== route.path) {
    await router.push(oldRoute?.path);
  }

  watch(
    () => config.config,
    () => {
      config.fetchSetConfig();
    }
  );

  watchEffect(() => {
    config.route = {
      path: route.path,
    };
  });
};
