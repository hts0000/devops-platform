import camelcaseKeys from "camelcase-keys";

export namespace DevOpsService {
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

  export const SendRequestWithAuthRetry = async <REQ, RES>(
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
      // throw res.status;
    }

    const data = await res.json();

    return o.respMarshaller(camelcaseKeys(data, { deep: true }));
  };
}
