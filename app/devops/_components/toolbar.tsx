import React from "react";

import { Plus, Trash, MoreHorizontal, Search } from "lucide-react";

import { Input } from "@/components/ui/input";
import ToolbarButton from "./toolbar-button";

type ToolbarProps = {};

const Toolbar = (props: ToolbarProps) => {
  return (
    <div className="flex m-2 mb-7 items-center justify-between bg-white">
      <div className="flex">
        <ToolbarButton icon={Plus} label="新建" />
        <ToolbarButton icon={Trash} label="删除" />
        <ToolbarButton icon={MoreHorizontal} label="更多" />
      </div>
      <div className="flex items-center justify-center">
        <Search className="size-5 mr-3 font-semibold text-gray-500" />
        <Input className="w-[700px]" />
      </div>
    </div>
  );
};

export default Toolbar;
