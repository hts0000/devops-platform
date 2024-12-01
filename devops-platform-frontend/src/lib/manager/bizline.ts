import {
  BizLine,
  BizLineEntity,
  CreateBizLineResponse,
  DeleteBizLineRequest,
  DeleteBizLineResponse,
  GetBizLineRequest,
  GetBizLinesRequest,
  GetBizLinesResponse,
} from "@/lib/manager/api/gen/v1/manager";
import { SendRequestWithAuthRetry } from "../request";

export const CreateBizLine = async (
  biz: BizLine
): Promise<CreateBizLineResponse> => {
  return SendRequestWithAuthRetry({
    method: "POST",
    path: "/v1/manager/bizline",
    data: biz,
    respMarshaller: CreateBizLineResponse.fromJSON,
  });
};

export const DeleteBizLine = async ({
  id,
}: DeleteBizLineRequest): Promise<DeleteBizLineResponse> => {
  return SendRequestWithAuthRetry({
    method: "DELETE",
    path: `/v1/manager/bizline/${id}`,
    respMarshaller: DeleteBizLineResponse.fromJSON,
  });
};

export const GetBizLine = ({
  id,
}: GetBizLineRequest): Promise<BizLineEntity> => {
  return SendRequestWithAuthRetry({
    method: "POST",
    path: `/v1/manager/bizline/${id}`,
    respMarshaller: BizLineEntity.fromJSON,
  });
};

export const GetBizLines = ({
  page,
  pageSize,
}: GetBizLinesRequest): Promise<GetBizLinesResponse> => {
  return SendRequestWithAuthRetry({
    method: "GET",
    path: `/v1/manager/bizlines?page=${page}&page_size=${pageSize}`,
    respMarshaller: GetBizLinesResponse.fromJSON,
  });
};
