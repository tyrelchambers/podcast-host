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
import { SelectOptions } from "@/lib/types";
import { formatCategoryOptions } from "@/lib/utils";
interface Props extends SelectPrimitive.SelectProps {
  placeholder: string;
  options: string[] | SelectOptions[];
}

const Select = ({ placeholder, options, ...props }: Props) => {
  const opts: SelectOptions[] | string[] = isSelectOptions(options)
    ? formatCategoryOptions(options)
    : options;

  return (
    <CNSelect onValueChange={props.onValueChange}>
      <SelectTrigger className="flex-1">
        <SelectValue placeholder={placeholder} />
      </SelectTrigger>
      <SelectContent>
        <ScrollArea className="h-72 w-full">
          {opts?.map((option, id) => (
            <SelectItem key={id + placeholder} value={getValue(option)}>
              {getDisplay(option)}
            </SelectItem>
          ))}
        </ScrollArea>
      </SelectContent>
    </CNSelect>
  );
};

const isSelectOptions = (options: any): options is SelectOptions[] => {
  if (typeof options === "string") {
    return false;
  }

  return true;
};

const getValue = (option: string | SelectOptions) => {
  return typeof option === "string" ? option : option.value;
};

const getDisplay = (option: string | { value: string }) => {
  return typeof option === "string" ? option : option.value;
};

export default Select;
