import { defineStore } from "pinia";
import { ConfigGet, ConfigSet } from "wails/go/main/App";

export interface Config {
  route?: {
    path: string;
  };
}

const init: () => Config = () => ({});

let configFetchTimer = 0;

export const useConfig = defineStore("config", {
  state: init,
  getters: {
    config: (state) => ({
      route: state.route,
    }),
  },
  actions: {
    fetchSetConfig() {
      if (configFetchTimer) clearTimeout(configFetchTimer);
      setTimeout(() => {
        ConfigSet(JSON.stringify(this.config));
      }, 1000);
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
