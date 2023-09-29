import { PodcastSwitcher } from "@/components/PodcastSwitcher";
import DashNav from "@/components/dashboard/DashNav";
import { usePodcastStore } from "@/hooks/stores/podcastStore";
import React from "react";

interface Props {
  rootPath: string;
}

const DashHeader = ({ rootPath }: Props) => {
  const podcastStore = usePodcastStore();

  return (
    <header className="bg-background-alt h-full  flex flex-col">
      <header className="p-4">
        <p className="text-background-alt-foreground">Resonate</p>
      </header>
      <DashNav rootPath={rootPath} />
      <footer className="p-4">
        <PodcastSwitcher activePodcast={podcastStore.activePodcast} />
      </footer>
    </header>
  );
};

export default DashHeader;
