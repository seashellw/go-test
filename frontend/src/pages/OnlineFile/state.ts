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

export const useFileList = defineStore("fileList", {
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
      this.list = await fetchFileList(this.space);
      return this.list;
    },
  },
});
