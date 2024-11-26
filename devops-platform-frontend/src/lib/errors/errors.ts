const Code2Message: Record<number, string> = {
  0: "未知错误",
  1: "1错误",
  2: "2错误",
  3: "3错误",
  4: "4错误",
  5: "5错误",
  6: "资源已存在",
};

interface MessageFns<T> {
  fromJSON(object: any): T;
  toJSON(message: T): unknown;
}

export interface ApiError {
  code: number;
  message: string;
}

export const ApiError: MessageFns<ApiError> = {
  fromJSON(object: any): ApiError {
    if (!object.code) {
      return {
        code: 0,
        message: Code2Message[0],
      };
    }
    const code = object.code as number;
    return {
      code: code,
      message: Code2Message[code],
    };
  },

  toJSON(message: ApiError): unknown {
    const obj: any = {};
    if (message.code !== 0) {
      obj.code = message.code;
    }
    if (message.message !== "") {
      obj.message = message.message;
    }
    return obj;
  },
};
