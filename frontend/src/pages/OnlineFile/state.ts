import { checkStr, fetchFileList, FileItem } from "@/interface/cos";
import { defineStore } from "pinia";
import { MessagePlugin } from "tdesign-vue-next";

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
        MessagePlugin.error("获取文件列表失败");
        return;
      }
      this.list = res;
      return this.list;
    },
  },
});
