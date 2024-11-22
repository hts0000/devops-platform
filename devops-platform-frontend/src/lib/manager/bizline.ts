import {
  BizLine,
  BizLineEntity,
  CreateBizLineResponse,
  GetBizLineRequest,
  GetBizLinesRequest,
  GetBizLinesResponse,
} from "@/lib/manager/api/gen/v1/manager";
import { DevOpsService } from "../request";

export namespace ManagerService {
  export const GetBizLine = ({
    id,
  }: GetBizLineRequest): Promise<BizLineEntity> => {
    return DevOpsService.SendRequestWithAuthRetry({
      method: "POST",
      path: `/v1/manager/bizline/${id}`,
      respMarshaller: BizLineEntity.fromJSON,
    });
  };

  export const GetBizLines = ({
    page,
    pageSize,
  }: GetBizLinesRequest): Promise<GetBizLinesResponse> => {
    return DevOpsService.SendRequestWithAuthRetry({
      method: "GET",
      path: `/v1/manager/bizlines?page=${page}&page_size=${pageSize}`,
      respMarshaller: GetBizLinesResponse.fromJSON,
    });
  };

  export const CreateBizLine = async (
    biz: BizLine
  ): Promise<CreateBizLineResponse> => {
    return DevOpsService.SendRequestWithAuthRetry({
      method: "POST",
      path: "/v1/manager/bizline",
      data: biz,
      respMarshaller: CreateBizLineResponse.fromJSON,
    });
  };
}
