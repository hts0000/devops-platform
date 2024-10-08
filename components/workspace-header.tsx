import React, { useState } from "react";

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Button } from "./ui/button";
import { ChevronDown, ListFilter, SquarePen } from "lucide-react";
import Hint from "./hint";
import PerferencesModal from "./perferences-modal";

const WorkspaceHeader = () => {
  const [perferencesOpen, setPerferencesOpen] = useState(false);

  return (
    <>
      <PerferencesModal
        open={perferencesOpen}
        setOpen={setPerferencesOpen}
        initialValue="workspace"
      />

      <div className="flex items-center justify-between px-4 h-[49px] gap-0.5">
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button
              variant={"transparent"}
              className="font-semibold text-lg w-auto p-1.5 overflow-hidden"
              size={"sm"}
            >
              <span className="truncate">{"Workspace1"}</span>
              <ChevronDown className="size-4 ml-1 shrink-0" />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent side="bottom" align="start" className="w-64">
            <DropdownMenuItem className="cursor-pointer capitalize">
              <div className="size-9 relative overflow-hidden bg-[#616061] text-white font-semibold text-xl rounded-md flex items-center justify-center mr-2">
                {"A"}
              </div>
              <div className="flex flex-col items-start">
                <p className="font-bold">{"worktest"}</p>
                <p className="text-xs text-muted-foreground">
                  {"Active workspace"}
                </p>
              </div>
            </DropdownMenuItem>
            <DropdownMenuSeparator />
            <DropdownMenuItem
              onClick={() => {}}
              className="cursor-pointer py-2"
            >
              Invite people to {"workspace"}
            </DropdownMenuItem>
            <DropdownMenuSeparator />
            <DropdownMenuItem
              onClick={() => setPerferencesOpen(true)}
              className="cursor-pointer py-2"
            >
              Preferences
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
        <div className="flex items-center gap-0.5">
          <Hint label="Filter conversations" side="bottom">
            <Button variant={"transparent"} size={"iconSm"}>
              <ListFilter className="size-4" />
            </Button>
          </Hint>
          <Hint label="New message" side="bottom">
            <Button variant={"transparent"} size={"iconSm"}>
              <SquarePen className="size-4" />
            </Button>
          </Hint>
        </div>
      </div>
    </>
  );
};

export default WorkspaceHeader;
