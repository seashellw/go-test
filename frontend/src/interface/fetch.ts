import { HttpFetch } from "wails/go/app/App";

export interface Request {
  Url: string;
  Method: string;
  Data: string;
  Cookie: {
    [key: string]: string;
  };
  Header: {
    [key: string]: string;
  };
}

export interface Response {
  Data?: string;
  Cookie: {
    [key: string]: string;
  };
  Header: {
    [key: string]: string;
  };
  Error?: {};
}

export const fetchHTTP: (
  request: Request
) => Promise<Response> = async (request) => {
  return (await HttpFetch(request as any)) as any;
};
