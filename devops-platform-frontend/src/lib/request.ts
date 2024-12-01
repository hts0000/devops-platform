import camelcaseKeys from "camelcase-keys";
import { ApiError } from "./errors/errors";

export const serverAddr = "http://localhost:18080";

export interface RequestOption<REQ, RES> {
  method: "GET" | "PUT" | "POST" | "DELETE";
  path: string;
  data?: REQ;
  respMarshaller: (r: object) => RES;
}

export interface AuthOption {
  attachAuthHeader: boolean;
  retryOnAuthError: boolean;
}

export const SendRequest = async <REQ, RES>(
  o: RequestOption<REQ, RES>,
  a?: AuthOption
): Promise<RES> => {
  const url = serverAddr + o.path;
  console.log("*****", url);

  const res = await fetch(url, {
    method: o.method,
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(o.data),
  });

  if (!res.ok) {
    const data = await res.json();

    const err = ApiError.fromJSON(camelcaseKeys(data, { deep: true }));

    // TODO: throw res.status
    console.log("request error:", res.status, err.code, err.message);
    throw err.message;
  }

  const data = await res.json();

  return o.respMarshaller(camelcaseKeys(data, { deep: true }));
};

export const SendRequestWithAuthRetry = async <REQ, RES>(
  o: RequestOption<REQ, RES>,
  a?: AuthOption
): Promise<RES> => {
  // TODO: with auth and retry
  return SendRequest(o, a);
};
