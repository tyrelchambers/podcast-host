import React from "react";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "./ui/select";

interface Props {
  time: string | undefined;
  setTime: (time: string) => void;
}

const ScheduleMeridiemSelector = ({ time, setTime }: Props) => {
  return (
    <Select defaultValue={time} onValueChange={setTime}>
      <SelectTrigger className="w-[180px]">
        <SelectValue />
      </SelectTrigger>
      <SelectContent>
        <SelectGroup>
          {["AM", "PM"].map((time) => (
            <SelectItem value={String(time)} key={time}>
              {time}
            </SelectItem>
          ))}
        </SelectGroup>
      </SelectContent>
    </Select>
  );
};

export default ScheduleMeridiemSelector;
