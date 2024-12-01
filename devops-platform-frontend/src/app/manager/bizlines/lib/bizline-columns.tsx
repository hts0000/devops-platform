"use client";

import { ColumnDef } from "@tanstack/react-table";
import { DotsHorizontalIcon } from "@radix-ui/react-icons";
import { format } from "date-fns";

import { BizLineEntity } from "./bizline-entity-schema";
import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";

export const bizLineColumns: ColumnDef<BizLineEntity>[] = [
  {
    id: "select",
    header: ({ table }) => (
      <Checkbox
        checked={
          table.getIsAllPageRowsSelected() ||
          (table.getIsSomePageRowsSelected() && "indeterminate")
        }
        onCheckedChange={(value) => table.toggleAllPageRowsSelected(!!value)}
        aria-label="Select all"
      />
    ),
    cell: ({ row }) => (
      <Checkbox
        checked={row.getIsSelected()}
        onCheckedChange={(value) => row.toggleSelected(!!value)}
        aria-label="Select row"
      />
    ),
    enableSorting: false,
    enableHiding: false,
  },
  {
    accessorKey: "bizLine.name",
    header: ({ column }) => <div>一级项目</div>,
    cell: ({ row }) => <div>{row.getValue("bizLine_name")}</div>,
  },
  {
    accessorKey: "bizLine.responsibleId",
    header: ({ column }) => <div>负责人</div>,
    cell: ({ row }) => <div>{row.getValue("bizLine_responsibleId")}</div>,
  },
  {
    accessorKey: "bizLine.description",
    header: ({ column }) => <div>描述</div>,
    cell: ({ row }) => <div>{row.getValue("bizLine_description")}</div>,
  },
  {
    accessorKey: "bizLine.createdAt",
    header: ({ column }) => <div>创建时间</div>,
    cell: ({ row }) => (
      <div>
        {format(
          (row.getValue("bizLine_createdAt") as number) * 1000,
          "yyyy-MM-dd HH:mm:ss"
        )}
      </div>
    ),
  },
  {
    id: "actions",
    header: ({ column }) => <div>更多</div>,
    cell: ({ row }) => {
      const payment = row.original;
      return (
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant="ghost" className="h-8 w-8 p-0">
              <span className="sr-only">Open menu</span>
              <DotsHorizontalIcon className="h-4 w-4" />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent align="end">
            <DropdownMenuLabel>Actions</DropdownMenuLabel>
            <DropdownMenuItem
              onClick={() => navigator.clipboard.writeText(payment.id)}
            >
              Copy payment ID
            </DropdownMenuItem>
            <DropdownMenuSeparator />
            <DropdownMenuItem>View customer</DropdownMenuItem>
            <DropdownMenuItem>View payment details</DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      );
    },
    enableHiding: false,
  },
];
