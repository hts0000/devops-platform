"use client";

import React from "react";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

import BizLineForm from "./components/bizline-form";

const BizLinePage = () => {
  const queryClient = new QueryClient();

  return (
    <div className="p-14">
      <p className="text-4xl mb-5 font-semibold">创建一级项目</p>
      <QueryClientProvider client={queryClient}>
        <BizLineForm />
      </QueryClientProvider>
    </div>
  );
};

export default BizLinePage;
