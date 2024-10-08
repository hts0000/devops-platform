import React from "react";

import { LucideIcon } from "lucide-react";
import { IconType } from "react-icons/lib";
import { Button } from "@/components/ui/button";
import { cn } from "@/lib/utils";

type ToolbarButtonProps = {
  icon: LucideIcon | IconType;
  label: string;
  isActive?: boolean;
};

const ToolbarButton = ({ icon: Icon, label, isActive }: ToolbarButtonProps) => {
  return (
    <div className="flex items-center justify-center gap-y-0.5 cursor-pointer group">
      <Button
        variant="outline"
        className={cn(
          "p-3 mr-2 rounded-lg items-center justify-center group-hover:bg-accent/20 group-hover:scale-105 transition-all",
          isActive && "bg-accent/20"
        )}
      >
        <Icon className="size-5 mr-2 font-semibold text-gray-500" />
        <span className="text-lg text-gray-500">{label}</span>
      </Button>
    </div>
  );
};

export default ToolbarButton;
