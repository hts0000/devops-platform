"use clinet";

import React, { useState } from "react";

import { Table } from "@tanstack/react-table";
import {
  MoreHorizontal,
  Plus,
  Trash,
  RotateCcw,
  SearchIcon,
} from "lucide-react";
import { useToast } from "@/hooks/use-toast";

import { Input } from "@/components/ui/input";
import BizLineTableToolbarButton from "./bizline-table-toolbar-button";
import { BizLineEntity } from "@/lib/manager/api/gen/v1/manager";
import { DeleteBizLine } from "@/lib/manager/bizline";

type BizLineTableToolbarProps<TData> = {
  table: Table<TData>;
};

const BizLineTableToolbar = <TData,>({
  table,
}: BizLineTableToolbarProps<TData>) => {
  // 获取选中的数据
  const selectedBizLineEntitys = table
    .getSelectedRowModel()
    .rows.map((row) => BizLineEntity.fromJSON(row.original));

  const { toast } = useToast();

  const [deletePending, setDeletePeding] = useState(false);

  const deleteBizLine = async () => {
    setDeletePeding(true);
    const deleteBizLine: string[] = [];
    try {
      const promises = selectedBizLineEntitys.map(async (be) => {
        try {
          const req = await DeleteBizLine({ id: be.id });

          toast({
            title: "删除成功",
            description: `一级项目名称：${be.bizLine?.name}`,
          });

          deleteBizLine.push(req.id);
        } catch (error: unknown) {
          toast({
            variant: "destructive",
            title: `${be.bizLine?.name} 删除失败`,
            description: `${error}`,
          });
        }
      });

      await Promise.all(promises);
    } catch (error) {
    } finally {
      console.log(deleteBizLine);
      setDeletePeding(false);
    }
  };

  return (
    <div className="flex items-center justify-between">
      <div className="flex flex-1 items-center space-x-2">
        <BizLineTableToolbarButton icon={Plus} label="新建" />
        <BizLineTableToolbarButton
          icon={Trash}
          onClick={deleteBizLine}
          label="删除"
          isActive={deletePending}
        />
        <BizLineTableToolbarButton icon={RotateCcw} label="刷新" />
        <BizLineTableToolbarButton icon={MoreHorizontal} label="更多" />
      </div>
      <Input
        placeholder="过滤一级项目..."
        value={
          (table.getColumn("bizLine_name")?.getFilterValue() as string) ?? ""
        }
        onChange={(event) =>
          table.getColumn("bizLine_name")?.setFilterValue(event.target.value)
        }
        className="h-8 w-[150px] lg:w-[250px]"
      />
      <SearchIcon className="size-4 ml-2 font-semibold text-gray-500" />
    </div>
  );
};

export default BizLineTableToolbar;
