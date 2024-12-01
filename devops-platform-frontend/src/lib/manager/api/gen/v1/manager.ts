// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.4.2
//   protoc               v5.27.2
// source: manager.proto

/* eslint-disable */

export const protobufPackage = "manager.v1";

export interface BizLine {
  name: string;
  responsibleId: string;
  description: string;
  createdAt: number;
}

export interface BizLineEntity {
  id: string;
  bizLine: BizLine | undefined;
}

export interface CreateBizLineResponse {
  id: string;
}

export interface DeleteBizLineRequest {
  id: string;
}

export interface DeleteBizLineResponse {
  id: string;
}

export interface UpdateBizLineResponse {
}

export interface GetBizLineRequest {
  id: string;
}

export interface GetBizLinesRequest {
  page: string;
  pageSize: string;
}

export interface GetBizLinesResponse {
  bizlines: BizLineEntity[];
  totalCount: string;
}

export interface Project {
}

export interface ProjectEntity {
  id: string;
  project: Project | undefined;
}

export interface CreateProjectRequest {
}

export interface CreateProjectResponse {
}

export interface DeleteProjectRequest {
}

export interface DeleteProjectResponse {
}

export interface UpdateProjectRequest {
}

export interface UpdateProjectResponse {
}

export interface GetProjectRequest {
}

export interface GetProjectResponse {
}

export interface GetProjectsRequest {
}

export interface GetProjectsResponse {
}

export interface App {
}

export interface AppEntity {
  id: string;
  app: App | undefined;
}

export interface CreateAppRequest {
}

export interface CreateAppResponse {
}

export interface DeleteAppRequest {
}

export interface DeleteAppResponse {
}

export interface UpdateAppRequest {
}

export interface UpdateAppResponse {
}

export interface GetAppRequest {
}

export interface GetAppResponse {
}

export interface GetAppsRequest {
}

export interface GetAppsResponse {
}

export const BizLine: MessageFns<BizLine> = {
  fromJSON(object: any): BizLine {
    return {
      name: isSet(object.name) ? globalThis.String(object.name) : "",
      responsibleId: isSet(object.responsibleId) ? globalThis.String(object.responsibleId) : "0",
      description: isSet(object.description) ? globalThis.String(object.description) : "",
      createdAt: isSet(object.createdAt) ? globalThis.Number(object.createdAt) : 0,
    };
  },

  toJSON(message: BizLine): unknown {
    const obj: any = {};
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.responsibleId !== "0") {
      obj.responsibleId = message.responsibleId;
    }
    if (message.description !== "") {
      obj.description = message.description;
    }
    if (message.createdAt !== 0) {
      obj.createdAt = Math.round(message.createdAt);
    }
    return obj;
  },
};

export const BizLineEntity: MessageFns<BizLineEntity> = {
  fromJSON(object: any): BizLineEntity {
    return {
      id: isSet(object.id) ? globalThis.String(object.id) : "0",
      bizLine: isSet(object.bizLine) ? BizLine.fromJSON(object.bizLine) : undefined,
    };
  },

  toJSON(message: BizLineEntity): unknown {
    const obj: any = {};
    if (message.id !== "0") {
      obj.id = message.id;
    }
    if (message.bizLine !== undefined) {
      obj.bizLine = BizLine.toJSON(message.bizLine);
    }
    return obj;
  },
};

export const CreateBizLineResponse: MessageFns<CreateBizLineResponse> = {
  fromJSON(object: any): CreateBizLineResponse {
    return { id: isSet(object.id) ? globalThis.String(object.id) : "0" };
  },

  toJSON(message: CreateBizLineResponse): unknown {
    const obj: any = {};
    if (message.id !== "0") {
      obj.id = message.id;
    }
    return obj;
  },
};

export const DeleteBizLineRequest: MessageFns<DeleteBizLineRequest> = {
  fromJSON(object: any): DeleteBizLineRequest {
    return { id: isSet(object.id) ? globalThis.String(object.id) : "0" };
  },

  toJSON(message: DeleteBizLineRequest): unknown {
    const obj: any = {};
    if (message.id !== "0") {
      obj.id = message.id;
    }
    return obj;
  },
};

export const DeleteBizLineResponse: MessageFns<DeleteBizLineResponse> = {
  fromJSON(object: any): DeleteBizLineResponse {
    return { id: isSet(object.id) ? globalThis.String(object.id) : "0" };
  },

  toJSON(message: DeleteBizLineResponse): unknown {
    const obj: any = {};
    if (message.id !== "0") {
      obj.id = message.id;
    }
    return obj;
  },
};

export const UpdateBizLineResponse: MessageFns<UpdateBizLineResponse> = {
  fromJSON(_: any): UpdateBizLineResponse {
    return {};
  },

  toJSON(_: UpdateBizLineResponse): unknown {
    const obj: any = {};
    return obj;
  },
};

export const GetBizLineRequest: MessageFns<GetBizLineRequest> = {
  fromJSON(object: any): GetBizLineRequest {
    return { id: isSet(object.id) ? globalThis.String(object.id) : "0" };
  },

  toJSON(message: GetBizLineRequest): unknown {
    const obj: any = {};
    if (message.id !== "0") {
      obj.id = message.id;
    }
    return obj;
  },
};

export const GetBizLinesRequest: MessageFns<GetBizLinesRequest> = {
  fromJSON(object: any): GetBizLinesRequest {
    return {
      page: isSet(object.page) ? globalThis.String(object.page) : "0",
      pageSize: isSet(object.pageSize) ? globalThis.String(object.pageSize) : "0",
    };
  },

  toJSON(message: GetBizLinesRequest): unknown {
    const obj: any = {};
    if (message.page !== "0") {
      obj.page = message.page;
    }
    if (message.pageSize !== "0") {
      obj.pageSize = message.pageSize;
    }
    return obj;
  },
};

export const GetBizLinesResponse: MessageFns<GetBizLinesResponse> = {
  fromJSON(object: any): GetBizLinesResponse {
    return {
      bizlines: globalThis.Array.isArray(object?.bizlines)
        ? object.bizlines.map((e: any) => BizLineEntity.fromJSON(e))
        : [],
      totalCount: isSet(object.totalCount) ? globalThis.String(object.totalCount) : "0",
    };
  },

  toJSON(message: GetBizLinesResponse): unknown {
    const obj: any = {};
    if (message.bizlines?.length) {
      obj.bizlines = message.bizlines.map((e) => BizLineEntity.toJSON(e));
    }
    if (message.totalCount !== "0") {
      obj.totalCount = message.totalCount;
    }
    return obj;
  },
};

export const Project: MessageFns<Project> = {
  fromJSON(_: any): Project {
    return {};
  },

  toJSON(_: Project): unknown {
    const obj: any = {};
    return obj;
  },
};

export const ProjectEntity: MessageFns<ProjectEntity> = {
  fromJSON(object: any): ProjectEntity {
    return {
      id: isSet(object.id) ? globalThis.String(object.id) : "0",
      project: isSet(object.project) ? Project.fromJSON(object.project) : undefined,
    };
  },

  toJSON(message: ProjectEntity): unknown {
    const obj: any = {};
    if (message.id !== "0") {
      obj.id = message.id;
    }
    if (message.project !== undefined) {
      obj.project = Project.toJSON(message.project);
    }
    return obj;
  },
};

export const CreateProjectRequest: MessageFns<CreateProjectRequest> = {
  fromJSON(_: any): CreateProjectRequest {
    return {};
  },

  toJSON(_: CreateProjectRequest): unknown {
    const obj: any = {};
    return obj;
  },
};

export const CreateProjectResponse: MessageFns<CreateProjectResponse> = {
  fromJSON(_: any): CreateProjectResponse {
    return {};
  },

  toJSON(_: CreateProjectResponse): unknown {
    const obj: any = {};
    return obj;
  },
};

export const DeleteProjectRequest: MessageFns<DeleteProjectRequest> = {
  fromJSON(_: any): DeleteProjectRequest {
    return {};
  },

  toJSON(_: DeleteProjectRequest): unknown {
    const obj: any = {};
    return obj;
  },
};

export const DeleteProjectResponse: MessageFns<DeleteProjectResponse> = {
  fromJSON(_: any): DeleteProjectResponse {
    return {};
  },

  toJSON(_: DeleteProjectResponse): unknown {
    const obj: any = {};
    return obj;
  },
};

export const UpdateProjectRequest: MessageFns<UpdateProjectRequest> = {
  fromJSON(_: any): UpdateProjectRequest {
    return {};
  },

  toJSON(_: UpdateProjectRequest): unknown {
    const obj: any = {};
    return obj;
  },
};

export const UpdateProjectResponse: MessageFns<UpdateProjectResponse> = {
  fromJSON(_: any): UpdateProjectResponse {
    return {};
  },

  toJSON(_: UpdateProjectResponse): unknown {
    const obj: any = {};
    return obj;
  },
};

export const GetProjectRequest: MessageFns<GetProjectRequest> = {
  fromJSON(_: any): GetProjectRequest {
    return {};
  },

  toJSON(_: GetProjectRequest): unknown {
    const obj: any = {};
    return obj;
  },
};

export const GetProjectResponse: MessageFns<GetProjectResponse> = {
  fromJSON(_: any): GetProjectResponse {
    return {};
  },

  toJSON(_: GetProjectResponse): unknown {
    const obj: any = {};
    return obj;
  },
};

export const GetProjectsRequest: MessageFns<GetProjectsRequest> = {
  fromJSON(_: any): GetProjectsRequest {
    return {};
  },

  toJSON(_: GetProjectsRequest): unknown {
    const obj: any = {};
    return obj;
  },
};

export const GetProjectsResponse: MessageFns<GetProjectsResponse> = {
  fromJSON(_: any): GetProjectsResponse {
    return {};
  },

  toJSON(_: GetProjectsResponse): unknown {
    const obj: any = {};
    return obj;
  },
};

export const App: MessageFns<App> = {
  fromJSON(_: any): App {
    return {};
  },

  toJSON(_: App): unknown {
    const obj: any = {};
    return obj;
  },
};

export const AppEntity: MessageFns<AppEntity> = {
  fromJSON(object: any): AppEntity {
    return {
      id: isSet(object.id) ? globalThis.String(object.id) : "0",
      app: isSet(object.app) ? App.fromJSON(object.app) : undefined,
    };
  },

  toJSON(message: AppEntity): unknown {
    const obj: any = {};
    if (message.id !== "0") {
      obj.id = message.id;
    }
    if (message.app !== undefined) {
      obj.app = App.toJSON(message.app);
    }
    return obj;
  },
};

export const CreateAppRequest: MessageFns<CreateAppRequest> = {
  fromJSON(_: any): CreateAppRequest {
    return {};
  },

  toJSON(_: CreateAppRequest): unknown {
    const obj: any = {};
    return obj;
  },
};

export const CreateAppResponse: MessageFns<CreateAppResponse> = {
  fromJSON(_: any): CreateAppResponse {
    return {};
  },

  toJSON(_: CreateAppResponse): unknown {
    const obj: any = {};
    return obj;
  },
};

export const DeleteAppRequest: MessageFns<DeleteAppRequest> = {
  fromJSON(_: any): DeleteAppRequest {
    return {};
  },

  toJSON(_: DeleteAppRequest): unknown {
    const obj: any = {};
    return obj;
  },
};

export const DeleteAppResponse: MessageFns<DeleteAppResponse> = {
  fromJSON(_: any): DeleteAppResponse {
    return {};
  },

  toJSON(_: DeleteAppResponse): unknown {
    const obj: any = {};
    return obj;
  },
};

export const UpdateAppRequest: MessageFns<UpdateAppRequest> = {
  fromJSON(_: any): UpdateAppRequest {
    return {};
  },

  toJSON(_: UpdateAppRequest): unknown {
    const obj: any = {};
    return obj;
  },
};

export const UpdateAppResponse: MessageFns<UpdateAppResponse> = {
  fromJSON(_: any): UpdateAppResponse {
    return {};
  },

  toJSON(_: UpdateAppResponse): unknown {
    const obj: any = {};
    return obj;
  },
};

export const GetAppRequest: MessageFns<GetAppRequest> = {
  fromJSON(_: any): GetAppRequest {
    return {};
  },

  toJSON(_: GetAppRequest): unknown {
    const obj: any = {};
    return obj;
  },
};

export const GetAppResponse: MessageFns<GetAppResponse> = {
  fromJSON(_: any): GetAppResponse {
    return {};
  },

  toJSON(_: GetAppResponse): unknown {
    const obj: any = {};
    return obj;
  },
};

export const GetAppsRequest: MessageFns<GetAppsRequest> = {
  fromJSON(_: any): GetAppsRequest {
    return {};
  },

  toJSON(_: GetAppsRequest): unknown {
    const obj: any = {};
    return obj;
  },
};

export const GetAppsResponse: MessageFns<GetAppsResponse> = {
  fromJSON(_: any): GetAppsResponse {
    return {};
  },

  toJSON(_: GetAppsResponse): unknown {
    const obj: any = {};
    return obj;
  },
};

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

export interface MessageFns<T> {
  fromJSON(object: any): T;
  toJSON(message: T): unknown;
}
