import React from "react";

import { Table } from "@tanstack/react-table";
import { MoreHorizontal, Plus, Trash } from "lucide-react";

import { Input } from "@/components/ui/input";
import BizLineTableToolbarButton from "./bizline-table-toolbar-button";

type BizLineTableToolbarProps<TData> = {
  table: Table<TData>;
};

const BizLineTableToolbar = <TData,>({
  table,
}: BizLineTableToolbarProps<TData>) => {
  return (
    <div className="flex items-center justify-between">
      <div className="flex flex-1 items-center space-x-2">
        <BizLineTableToolbarButton icon={Plus} label="新建" />
        <BizLineTableToolbarButton icon={Trash} label="删除" />
        <BizLineTableToolbarButton icon={MoreHorizontal} label="更多" />
      </div>
      <Input
        placeholder="Filter tasks..."
        value={(table.getColumn("bizLine")?.getFilterValue() as string) ?? ""}
        onChange={(event) =>
          table.getColumn("bizLine")?.setFilterValue(event.target.value)
        }
        className="h-8 w-[150px] lg:w-[250px]"
      />
    </div>
  );
};

export default BizLineTableToolbar;
