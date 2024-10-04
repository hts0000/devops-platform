"use client";

import React from "react";
import Sidebar from "./_components/sidebar";

type DevopsLayoutProps = {
  children: React.ReactNode;
};

const DevopsLayout = ({ children }: DevopsLayoutProps) => {
  return (
    <div className="h-full flex">
      <Sidebar />
      {children}
    </div>
  );
};

export default DevopsLayout;
