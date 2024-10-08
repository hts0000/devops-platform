import React from "react";
import Toolbar from "./_components/toolbar";
import AppTable from "./_components/app-table";
import PaymentTable from "./_components/payment-table";
import { Table } from "@/components/ui/table";
import { TableDemo } from "./_components/demo-table";

type Props = {};

const DevopsPage = (props: Props) => {
  return (
    <div className="w-full px-8 py-5 bg-gray-100">
      <div className="h-full rounded-xl p-10 bg-white shadow-md">
        <p className="mb-9 text-black text-2xl font-semibold">应用部署</p>
        <Toolbar />
        {/* <TableDemo /> */}
        <AppTable />
        {/* <PaymentTable /> */}
      </div>
    </div>
  );
};

export default DevopsPage;
