import React from "react";
import {
  Select as CNSelect,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "./ui/select";
import * as SelectPrimitive from "@radix-ui/react-select";
import { ScrollArea } from "./ui/scroll-area";

interface Props extends SelectPrimitive.SelectProps {
  placeholder: string;
  options: string[];
}

const Select = ({ placeholder, options, ...props }: Props) => {
  return (
    <CNSelect onValueChange={props.onValueChange}>
      <SelectTrigger className="flex-1">
        <SelectValue placeholder={placeholder} />
      </SelectTrigger>
      <SelectContent>
        <ScrollArea className="h-72 w-full">
          {options?.map((option, id) => (
            <SelectItem key={id} value={option}>
              {option}
            </SelectItem>
          ))}
        </ScrollArea>
      </SelectContent>
    </CNSelect>
  );
};

export default Select;
