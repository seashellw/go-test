import { FileItem, readDir } from "@/interface/path";
import { defineStore } from "pinia";

const init: () => {
  dir: string | undefined;
  list: FileItem[];
} = () => ({
  dir: undefined,
  list: [],
});

export const useFileList = defineStore("localFileList", {
  state: init,
  getters: {},
  actions: {
    async getList() {
      if (!this.dir) {
        this.list = [];
        return;
      }
      this.list = await readDir(this.dir);
      return this.list;
    },
  },
});
