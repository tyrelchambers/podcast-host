"use client";

import * as React from "react";
import { Check, ChevronsUpDown } from "lucide-react";

import { cn, formatUrlFromTitle } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
} from "@/components/ui/command";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { Podcast } from "@/lib/types";
import { usePodcastStore } from "@/hooks/stores/podcastStore";
import { useRouter } from "next/router";

export function PodcastSwitcher({
  activePodcast,
}: {
  activePodcast: Podcast | undefined;
}) {
  const router = useRouter();
  const podcastState = usePodcastStore((state) => state.podcasts);
  const [open, setOpen] = React.useState(false);
  const [value, setValue] = React.useState("");

  const podcasts = Array.from(podcastState.values());

  const clickHandler = (currentValue: string) => {
    if (currentValue !== value) {
      setValue(currentValue);
      router.push({
        query: {
          name: formatUrlFromTitle(currentValue),
        },
      });
    }
    setOpen(false);
  };

  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button
          variant="outline"
          role="combobox"
          aria-expanded={open}
          className="w-[200px] justify-between text-foreground"
        >
          {value ? podcastState.get(value)?.title : activePodcast?.title}
          <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-[200px] p-0">
        <Command>
          <CommandInput placeholder="Search podcasts..." />
          <CommandEmpty>No podcast found.</CommandEmpty>
          <CommandGroup>
            {podcasts.map((p) => (
              <CommandItem key={p.id} onSelect={clickHandler}>
                <Check
                  className={cn(
                    "mr-2 h-4 w-4",
                    value === p.title ? "opacity-100" : "opacity-0"
                  )}
                />
                {p.title}
              </CommandItem>
            ))}
          </CommandGroup>
        </Command>
      </PopoverContent>
    </Popover>
  );
}
