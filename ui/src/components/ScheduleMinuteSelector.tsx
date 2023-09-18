import React from "react";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { ScrollArea } from "./ui/scroll-area";

interface Props {
  time: string | undefined;
  setTime: (time: string) => void;
}

const getTime = () => {
  const time = [];
  for (let index = 0; index <= 60; index++) {
    let formattedNumber = index.toLocaleString("en-US", {
      minimumIntegerDigits: 2,
      useGrouping: false,
    });

    time.push(formattedNumber);
  }

  return time;
};
const ScheduleMinuteSelector = ({ setTime, time }: Props) => {
  return (
    <Select defaultValue={time} onValueChange={setTime}>
      <SelectTrigger className="w-[180px]">
        <SelectValue />
      </SelectTrigger>
      <SelectContent>
        <ScrollArea className="h-72 w-48">
          <SelectGroup>
            {getTime().map((time) => (
              <SelectItem value={time} key={time}>
                {time}
              </SelectItem>
            ))}
          </SelectGroup>
        </ScrollArea>
      </SelectContent>
    </Select>
  );
};

export default ScheduleMinuteSelector;
