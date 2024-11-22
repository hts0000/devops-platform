"use client";

import React from "react";

import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation, useQueryClient } from "@tanstack/react-query";

import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { BizLineFormType, BizLineFormSchema } from "../lib/bizline-form-schema";
import { ManagerService } from "@/lib/manager/bizline";
import { BizLine } from "@/lib/manager/api/gen/v1/manager";

type BizLineFormProps = {};

const BizLineForm = ({}: BizLineFormProps) => {
  // const queryClient = useQueryClient();

  const form = useForm<BizLineFormType>({
    resolver: zodResolver(BizLineFormSchema),
    defaultValues: {
      name: "b-end",
      responsibleId: "1",
      description: "B端",
    },
  });

  const mutation = useMutation({
    mutationFn: ManagerService.CreateBizLine,
    onSuccess: (data, variables, context) => {
      // 表单提交成功后，刷新相关数据
      // queryClient.invalidateQueries(["items"]);

      // TODO: toast通知

      form.reset(); // 清空表单
    },
    onError: (error, variables, context) => {
      alert(`Error: ${error.message}`);
    },
  });

  const onSubmit = (data: BizLineFormType) => {
    // console.log(data);
    // const biz = BizLine.fromJSON(data);
    // console.log("biz", biz);
    mutation.mutate(BizLine.fromJSON(data));
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <FormField
          control={form.control}
          name="name"
          render={({ field }) => (
            <FormItem>
              <FormLabel>一级项目名称</FormLabel>
              <FormControl>
                <Input placeholder="请输入" {...field} />
              </FormControl>
              <FormDescription>一级项目名称</FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="responsibleId"
          render={({ field }) => (
            <FormItem>
              <FormLabel>负责人</FormLabel>
              <FormControl>
                <Input placeholder="请输入" {...field} />
              </FormControl>
              <FormDescription>负责人</FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="description"
          render={({ field }) => (
            <FormItem>
              <FormLabel>描述</FormLabel>
              <FormControl>
                <Input placeholder="请输入" {...field} />
              </FormControl>
              <FormDescription>描述</FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <Button type="submit" disabled={mutation.isPending}>
          创建
        </Button>
      </form>
    </Form>
  );
};

export default BizLineForm;
