import {
  BizLineEntity,
  GetBizLineRequest,
  GetBizLinesRequest,
  GetBizLinesResponse,
} from "@/lib/manager/api/gen/v1/manager";

export const GetBizLine = async ({
  id,
}: GetBizLineRequest): Promise<BizLineEntity> => {
  const res = await fetch(
    `http://localhost:18080/v1/manager/bizline/${id.toString()}`
  );
  const data = await res.json();

  return BizLineEntity.fromJSON(data);
};

export const GetBizLines = async ({
  page,
  pageSize,
}: GetBizLinesRequest): Promise<GetBizLinesResponse> => {
  const res = await fetch(
    `http://localhost:18080/v1/manager/bizlines?page=${page}&page_size=${pageSize}`
  );
  const data = await res.json();

  return GetBizLinesResponse.fromJSON(data);
};
