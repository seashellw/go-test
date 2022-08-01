import { checkStr, fetchFileList, FileItem } from "@/interface/cos";
import { defineStore } from "pinia";

export interface UploadItem {
  name: string;
  percent: number;
  status: "success" | "error" | undefined;
}

const init: () => {
  space: string;
  list: FileItem[];
  uploadList: UploadItem[];
} = () => ({
  space: "",
  list: [],
  uploadList: [],
});

export const useFileList = defineStore("onlineFileList", {
  state: init,
  getters: {
    isSpaceError: (state) => {
      return !checkStr(state.space);
    },
  },
  actions: {
    async fetchList() {
      if (!this.space) {
        this.list = [];
        return;
      }
      if (this.isSpaceError) return;
      let res = await fetchFileList(this.space);
      if (!res) {
        return;
      }
      this.list = res;
      return this.list;
    },
  },
});
