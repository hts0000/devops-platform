import {
  BizLine,
  BizLineEntity,
  CreateBizLineResponse,
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

  // TODO: handle error
  if (!res.ok) {
  }

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

  // TODO: handle error
  if (!res.ok) {
  }

  const data = await res.json();

  return GetBizLinesResponse.fromJSON(data);
};

export const CreateBizLine = async (
  biz: BizLine
): Promise<CreateBizLineResponse> => {
  console.log("###############", biz);
  const res = await fetch(`http://localhost:18080/v1/manager/bizline`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(biz),
  });
  console.log("*****************");

  // TODO: handle error
  if (!res.ok) {
    // console.log("&&&&&&&&&&&&&&");
  }

  const data = await res.json();

  return CreateBizLineResponse.fromJSON(data);
};
