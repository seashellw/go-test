import {
  DialogDirSelect,
  FileGetAllFileList,
  FileHash,
  FileReadDir,
} from "wails/go/main/App";

export const getExtensionName = (path: string) =>
  path.slice(path.lastIndexOf("."));

export const getDirPath = (path: string) => {
  let pathList = path.split("\\").filter((item) => item);
  if (pathList.length <= 1) return undefined;
  pathList = pathList.slice(0, -1);
  if (pathList.length <= 1) return pathList.join() + "\\";
  return pathList.join("\\");
};

export const getName = (path: string) => path.split("\\").pop();

export interface FileItem {
  path: string;
  name: string;
  isDir: boolean;
  size: number;
}

const sortingFileItem: (list: any[]) => FileItem[] = (list) => {
  return list.map((item) => ({
    path: item.Path,
    name: item.Name,
    isDir: item.IsDir,
    size: item.Size,
  }));
};

export const getAllFileList: (
  path: string
) => Promise<FileItem[]> = async (path: string) => {
  const res = await FileGetAllFileList(path);
  return sortingFileItem(res);
};

export const getFileHash = async (path: string) => {
  const res = await FileHash(path);
  return `${res}`;
};

export const readDir: (path: string) => Promise<FileItem[]> = async (
  path: string
) => {
  const res = await FileReadDir(path);
  return sortingFileItem(res);
};

export const openDirDialog = async () => {
  const res = await DialogDirSelect();
  if (!res) return undefined;
  return res;
};
