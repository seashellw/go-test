import { defineStore } from "pinia";

export const useAppStore = defineStore("app", {
  state: () => ({
    count: 0,
  }),
  actions: {
    async add() {
      this.count++;
    },
  },
});
