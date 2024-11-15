import React from "react";

import { Table } from "@tanstack/react-table";

type BizLineTablePaginationProps<TData> = {
  table: Table<TData>;
};

const BizLineTablePagination = <TData,>({
  table,
}: BizLineTablePaginationProps<TData>) => {
  return <div>BizLineTablePagination</div>;
};

export default BizLineTablePagination;
