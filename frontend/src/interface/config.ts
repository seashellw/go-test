import { get, post } from ".";

export interface Config {
  route?: string;
}

interface Response {
  Config: string;
}

export const fetchConfig = async (data?: Config) => {
  if (data && typeof data === "object") {
    await post<undefined, Response>("/api/configSet", {
      Config: JSON.stringify(data),
    });
  } else {
    data = await get<Response>("/api/config")
      .then((res) => res.Config)
      .then((res) => {
        try {
          return JSON.parse(res || "{}");
        } catch {
          return {};
        }
      });
  }
  if (!data) data = {};
  return data;
};
