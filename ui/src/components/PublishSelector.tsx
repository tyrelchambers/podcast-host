import React from "react";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "./ui/select";

type WhenPublish = "immediate" | "schedule";

interface Props {
  setWhenToPublish: (whenToPublish: WhenPublish) => void;
}

const PublishSelector = ({ setWhenToPublish }: Props) => {
  return (
    <Select onValueChange={(v) => setWhenToPublish(v as WhenPublish)}>
      <SelectTrigger className="w-[180px]">
        <SelectValue placeholder="Publish now" />
      </SelectTrigger>
      <SelectContent>
        <SelectItem value="immediate">Publish immediately</SelectItem>
        <SelectItem value="schedule">Schedule for later</SelectItem>
      </SelectContent>
    </Select>
  );
};

export default PublishSelector;
