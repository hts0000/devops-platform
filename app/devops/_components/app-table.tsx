"use client";

import React from "react";

import {
  CaretSortIcon,
  ChevronDownIcon,
  DotsHorizontalIcon,
} from "@radix-ui/react-icons";
import {
  ColumnDef,
  ColumnFiltersState,
  SortingState,
  VisibilityState,
  flexRender,
  getCoreRowModel,
  getFilteredRowModel,
  getPaginationRowModel,
  getSortedRowModel,
  useReactTable,
} from "@tanstack/react-table";

import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import {
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Input } from "@/components/ui/input";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { cn } from "@/lib/utils";

const data: AppInfo[] = [
  {
    id: "m5gr84i1",
    bizLine: "海外供应链",
    project: "CPS集采系统",
    app: "cps-1",
    shortcuts: {
      Gitlab: "",
      Jenkins: "",
      K8S: "",
      SLS: "",
    },
    status: "live",
  },
  {
    id: "m5gr84i2",
    bizLine: "国内供应链",
    project: "ISCM系统",
    app: "iscm-1",
    shortcuts: {
      Gitlab: "",
      Jenkins: "",
      K8S: "",
      SLS: "",
    },
    status: "live",
  },
  {
    id: "m5gr84i3",
    bizLine: "国内供应链",
    project: "ISCM系统",
    app: "iscm-1",
    shortcuts: {
      Gitlab: "",
      Jenkins: "",
      K8S: "",
      SLS: "",
    },
    status: "live",
  },
  {
    id: "m5gr84i4",
    bizLine: "国内供应链",
    project: "ISCM系统",
    app: "iscm-1",
    shortcuts: {
      Gitlab: "",
      Jenkins: "",
      K8S: "",
      SLS: "",
    },
    status: "live",
  },
  {
    id: "m5gr84i5",
    bizLine: "国内供应链",
    project: "ISCM系统",
    app: "iscm-1",
    shortcuts: {
      Gitlab: "",
      Jenkins: "",
      K8S: "",
      SLS: "",
    },
    status: "live",
  },
  {
    id: "m5gr84i6",
    bizLine: "国内供应链",
    project: "ISCM系统",
    app: "iscm-1",
    shortcuts: {
      Gitlab: "",
      Jenkins: "",
      K8S: "",
      SLS: "",
    },
    status: "live",
  },
  {
    id: "m5gr84i7",
    bizLine: "国内供应链",
    project: "ISCM系统",
    app: "iscm-1",
    shortcuts: {
      Gitlab: "",
      Jenkins: "",
      K8S: "",
      SLS: "",
    },
    status: "live",
  },
];

type Shortcut = {
  Gitlab?: string;
  Jenkins?: string;
  K8S?: string;
  SLS?: string;
};

export type AppInfo = {
  id: string;
  bizLine: string;
  project: string;
  environment: "TEST" | "PRE" | "PRD";
  app: string;
  shortcuts?: Shortcut;
  status?: "live" | "dead";
};

export const columns: ColumnDef<AppInfo>[] = [
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
    accessorKey: "bizLine",
    header: "业务",
    cell: ({ row }) => <div>{row.getValue("bizLine")}</div>,
    size: 400,
  },
  {
    accessorKey: "project",
    header: "项目",
    cell: ({ row }) => <div>{row.getValue("project")}</div>,
    size: 500,
  },
  {
    accessorKey: "app",
    header: ({ column }) => {
      return (
        <Button
          variant="ghost"
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
        >
          应用
          <CaretSortIcon className="ml-2 h-4 w-4" />
        </Button>
      );
    },
    cell: ({ row }) => <div className="lowercase">{row.getValue("app")}</div>,
    size: 400,
  },
  {
    accessorKey: "shortcut",
    header: "传送门",
    cell: ({ row }) => <div>{row.getValue("shortcuts")}</div>,
    size: 500,
  },
  {
    id: "actions",
    enableHiding: false,
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
  },
];

type AppTableProps = {};

const AppTable = (props: AppTableProps) => {
  // TODO: label
  // format: brand-appname-program-environment
  // brand: miniso/toptoy/wow
  // program: java/web/go/php
  // environment: test/pre/prd
  // example1: miniso-app1-java-test
  // example2: toptoy-app2-web-prd

  const [sorting, setSorting] = React.useState<SortingState>([]);
  const [columnFilters, setColumnFilters] = React.useState<ColumnFiltersState>(
    []
  );
  const [columnVisibility, setColumnVisibility] =
    React.useState<VisibilityState>({});
  const [rowSelection, setRowSelection] = React.useState({});

  const table = useReactTable({
    data,
    columns,
    onSortingChange: setSorting,
    onColumnFiltersChange: setColumnFilters,
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    getSortedRowModel: getSortedRowModel(),
    getFilteredRowModel: getFilteredRowModel(),
    onColumnVisibilityChange: setColumnVisibility,
    onRowSelectionChange: setRowSelection,
    enableColumnResizing: true,
    columnResizeMode: "onEnd",
    columnResizeDirection: "rtl",
    state: {
      sorting,
      columnFilters,
      columnVisibility,
      rowSelection,
    },
  });

  return (
    <div className="w-full">
      <div className="flex items-center py-4">
        <Input
          placeholder="Filter emails..."
          value={(table.getColumn("email")?.getFilterValue() as string) ?? ""}
          onChange={(event) =>
            table.getColumn("email")?.setFilterValue(event.target.value)
          }
          className="max-w-sm"
        />
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant="outline" className="ml-auto">
              Columns <ChevronDownIcon className="ml-2 h-4 w-4" />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent align="end">
            {table
              .getAllColumns()
              .filter((column) => column.getCanHide())
              .map((column) => {
                return (
                  <DropdownMenuCheckboxItem
                    key={column.id}
                    className="capitalize"
                    checked={column.getIsVisible()}
                    onCheckedChange={(value) =>
                      column.toggleVisibility(!!value)
                    }
                  >
                    {column.id}
                  </DropdownMenuCheckboxItem>
                );
              })}
          </DropdownMenuContent>
        </DropdownMenu>
      </div>
      <div className="rounded-md border">
        <Table>
          <TableHeader>
            {table.getHeaderGroups().map((headerGroup) => (
              <TableRow key={headerGroup.id}>
                {headerGroup.headers.map((header) => {
                  return (
                    <TableHead
                      key={header.id}
                      className={`w-[${header.column.getSize()}px]`}
                    >
                      {header.isPlaceholder
                        ? null
                        : flexRender(
                            header.column.columnDef.header,
                            header.getContext()
                          )}
                    </TableHead>
                  );
                })}
              </TableRow>
            ))}
          </TableHeader>
          <TableBody>
            {table.getRowModel().rows?.length ? (
              table.getRowModel().rows.map((row) => (
                <TableRow
                  key={row.id}
                  data-state={row.getIsSelected() && "selected"}
                >
                  {row.getVisibleCells().map((cell) => (
                    <TableCell key={cell.id}>
                      {flexRender(
                        cell.column.columnDef.cell,
                        cell.getContext()
                      )}
                    </TableCell>
                  ))}
                </TableRow>
              ))
            ) : (
              <TableRow>
                <TableCell
                  colSpan={columns.length}
                  className="h-24 text-center"
                >
                  No results.
                </TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
      </div>
      <div className="flex items-center justify-end space-x-2 py-4">
        <div className="flex-1 text-sm text-muted-foreground">
          {table.getFilteredSelectedRowModel().rows.length} of{" "}
          {table.getFilteredRowModel().rows.length} row(s) selected.
        </div>
        <div className="space-x-2">
          <Button
            variant="outline"
            size="sm"
            onClick={() => table.previousPage()}
            disabled={!table.getCanPreviousPage()}
          >
            Previous
          </Button>
          <Button
            variant="outline"
            size="sm"
            onClick={() => table.nextPage()}
            disabled={!table.getCanNextPage()}
          >
            Next
          </Button>
        </div>
      </div>
    </div>
  );
};

export default AppTable;
