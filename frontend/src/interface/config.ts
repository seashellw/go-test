import { ConfigGet, ConfigSet } from "wails/go/main/App";

export interface Config {
  route?: string;
}

export const setConfig = (config: Config) => {
  ConfigSet(JSON.stringify(config));
};

export const getConfig: () => Promise<Config> = async () => {
  let res = await ConfigGet();
  try {
    return JSON.parse(res || "{}");
  } catch {
    return {};
  }
};
