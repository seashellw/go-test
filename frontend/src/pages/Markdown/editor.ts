import { defineStore } from "pinia";

const init = () => ({
  header: {
    showNewButton: true,
    handleClickNewButton: () => {},
  },
  dir: "",
});

export const useEditorState = defineStore("editor", {
  state: init,
});
