import React from "react";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "./ui/select";
import { PodcastCategoryOption } from "@/lib/types";
import { formatCategoryOptions } from "@/lib/utils";

interface Props {
  placeholder: string;
  options: PodcastCategoryOption[];
}

const SelectCategory = ({ placeholder, options }: Props) => {
  const opts = formatCategoryOptions(options);

  return (
    <Select>
      <SelectTrigger className="flex-1">
        <SelectValue placeholder={placeholder} />
      </SelectTrigger>
      <SelectContent>
        {opts?.map((option, id) => (
          <SelectItem key={id} value={option.value}>
            {option.value}
          </SelectItem>
        ))}
      </SelectContent>
    </Select>
  );
};

export default SelectCategory;
