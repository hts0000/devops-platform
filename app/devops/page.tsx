import React from "react";
import Toolbar from "./_components/toolbar";
import AppTable from "./_components/app-table";

type Props = {};

const DevopsPage = (props: Props) => {
  return (
    <div className="w-full px-8 py-5 bg-gray-100">
      <div className="h-full rounded-xl p-8 bg-white shadow-md">
        <p className="m-2 mb-9 text-black text-2xl font-semibold">应用部署</p>
        <Toolbar />
        <AppTable />
      </div>
    </div>
  );
};

export default DevopsPage;
