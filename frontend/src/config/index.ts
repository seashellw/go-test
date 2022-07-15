import { Config } from "@/interface/config";
import { defineStore } from "pinia";

const init: () => Config = () => ({});

export const useConfig = defineStore("config", {
  state: init,
});
