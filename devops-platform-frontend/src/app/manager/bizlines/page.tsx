import React from "react";

import { z } from "zod";

import BizLineTable from "./components/bizline-table";
import { BizLineEntitySchema } from "./lib/bizline-entity-schema";
import { bizLineColumns } from "./lib/bizline-columns";
import { GetBizLines } from "@/lib/manager/bizline";

const BizLinesPage = async () => {
  const page = "1";
  const pageSize = "20";
  const res = await GetBizLines({
    page,
    pageSize,
  });

  const { bizlines, totalCount } = res;
  console.log("get bizlines success", bizlines, totalCount);

  const result = z.array(BizLineEntitySchema).safeParse(bizlines);
  if (!result.success) {
    console.log("parse data failed, err:", result.error);
    // TODO: 直接返回出错组件
    // <Error />
    return;
  }

  const zodBizs = result.data;

  return (
    <div className="p-14">
      <div className="mb-5">展示一级项目数量等统计信息</div>
      <BizLineTable data={zodBizs} columns={bizLineColumns} />
    </div>
  );
};

export default BizLinesPage;
