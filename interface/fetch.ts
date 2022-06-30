import { post } from "interface";

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
  Data: string;
  Cookie: {
    [key: string]: string;
  };
  Header: {
    [key: string]: string;
  };
}

export const fetchFetch = (data: Request) =>
  post<Response, Request>("/api/fetch", data);
