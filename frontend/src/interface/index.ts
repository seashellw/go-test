export type HTTP_METHODS_TYPE = "GET" | "POST" | "PUT" | "DELETE";

/**
 * 使用fetcher发送请求的参数类型定义
 */
interface FetchConfig {
  url: string;
  data?: any;
  method?: HTTP_METHODS_TYPE;
  headers?: {
    [key: string]: string;
  };
  init?: RequestInit;
}

/**
 * 发送一个HTTP请求
 */
export const fetcher = async (config: FetchConfig) => {
  let { url, data, method, init, headers } = config;
  method = method ?? "GET";
  init = init ?? {};
  data = data ?? {};
  headers = headers ?? {};
  if (method === "POST") {
    headers = { "Content-Type": "application/json", ...headers };
  }
  const res = await fetch(url, {
    method,
    body: method === "POST" ? JSON.stringify(data) : undefined,
    headers,
    ...init,
  });
  return await res.json();
};

/**
 * 发送post请求
 */
export const post: <Response = any, Request = any>(
  url: string,
  data?: Request
) => Promise<Response> = (url, data) =>
  fetcher({
    url,
    data,
    method: "POST",
  });

/**
 * 发送get请求
 */
export const get: <Response = any>(
  url: string
) => Promise<Response> = (url: string) =>
  fetcher({ url, method: "GET" });
