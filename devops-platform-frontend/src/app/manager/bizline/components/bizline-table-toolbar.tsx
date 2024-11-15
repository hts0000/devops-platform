import React from "react";

import { Table } from "@tanstack/react-table";

type BizLineTableToolbarProps<TData> = {
  table: Table<TData>;
};

const BizLineTableToolbar = <TData,>({
  table,
}: BizLineTableToolbarProps<TData>) => {
  return <div>BizLineTableToolbar</div>;
};

export default BizLineTableToolbar;
