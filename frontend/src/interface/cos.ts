import COS from "cos-js-sdk-v5";
import dayjs from "dayjs";
import { BrowserOpenURL } from "wails/runtime";

export const Bucket = "cache-1259243245";
export const Region = "ap-beijing";

export const AuthorizationUrl =
  "http://152.136.121.30:8080/api/file/authorization";

export const DownloadUrl =
  "https://cache-1259243245.cos-website.ap-beijing.myqcloud.com";

// 初始化实例
const cos = new COS({
  getAuthorization: function (_, callback) {
    fetch(AuthorizationUrl)
      .then((value) => value.json())
      .then((res) => {
        if (!res.ok) {
          console.error(res);
          return;
        }
        callback({
          TmpSecretId: res.tempKeys.credentials.tmpSecretId,
          TmpSecretKey: res.tempKeys.credentials.tmpSecretKey,
          SecurityToken: res.tempKeys.credentials.sessionToken,
          StartTime: res.tempKeys.startTime,
          ExpiredTime: res.tempKeys.expiredTime,
        });
      });
  },
});

export const checkStr = (name: string) => {
  return /^[^\t\r\n\v/]+$/.test(name);
};

export const generateKey = (param: { space: string; name: string }) =>
  `${param.space}/${param.name}`;

export const parseKey = (key: string) => {
  const index = key.indexOf("/");
  if (index < 0) return null;
  return {
    space: key.slice(0, index),
    name: key.slice(index + 1),
  };
};

export const fetchUpload = (param: {
  space: string;
  name: string;
  body: File;
  onProgress: (progressData: {
    // 已经上传的文件部分大小，以字节（Bytes）为单位
    loaded: number;
    // 整个文件的大小，以字节（Bytes）为单位
    total: number;
    // 文件的上传速度，以字节/秒（Bytes/s）为单位
    speed: number;
    // 文件的上传百分比，以小数形式呈现，例如，上传50%即为0.5
    percent: number;
  }) => void;
}) =>
  new Promise<{
    ok: boolean;
  }>((resolve) => {
    cos.uploadFile(
      {
        Bucket,
        Region,
        Key: generateKey(param),
        Body: param.body,
        onProgress: param.onProgress,
      },
      function (err) {
        if (err) {
          console.error(err);
          resolve({ ok: false });
          return;
        }
        resolve({ ok: true });
      }
    );
  });

export interface FileItem {
  space: string;
  // 文件名
  name: string;
  // 上传时间
  time: number;
  // 文件大小：以字节为单位
  size: number;
}

export const fetchFileList = (space: string) =>
  new Promise<FileItem[]>((resolve) => {
    cos.getBucket(
      {
        Bucket,
        Region,
        Prefix: space,
      },
      function (err, data) {
        if (err) {
          console.error(err);
          resolve([]);
          return;
        }
        let resList: FileItem[] = [];
        for (const item of data.Contents) {
          let keyObj = parseKey(item.Key);
          if (!keyObj) {
            console.log("文件key解析错误", item.Key);
            resolve([]);
            return;
          }
          resList.push({
            time: dayjs(item.LastModified).valueOf(),
            size: parseInt(item.Size),
            ...keyObj,
          });
        }
        resolve(resList);
      }
    );
  });

export const fetchDeleteFile = (item: {
  space: string;
  name: string;
}) =>
  new Promise<boolean>((resolve) => {
    cos.deleteObject(
      {
        Bucket,
        Region,
        Key: generateKey(item),
      },
      function (err) {
        if (err) {
          console.error(err);
          resolve(false);
          return;
        }
        resolve(true);
      }
    );
  });

export const fetchFileExist = (item: {
  space: string;
  name: string;
}) =>
  new Promise<boolean>((resolve) => {
    cos.headObject(
      {
        Bucket,
        Region,
        Key: generateKey(item),
      },
      function (err, data) {
        if (data || !err) {
          resolve(true);
          return;
        }
        resolve(false);
      }
    );
  });

export const getFileUrl = (item: { space: string; name: string }) => {
  return `${DownloadUrl}/${generateKey(item)}`;
};

export const downloadFile = (item: {
  space: string;
  name: string;
}) => {
  BrowserOpenURL(getFileUrl(item));
};
