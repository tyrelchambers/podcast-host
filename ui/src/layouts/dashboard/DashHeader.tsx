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
    <header className="bg-card h-full p-4 flex flex-col">
      <p>Resonate</p>
      <DashNav rootPath={rootPath} />
      <PodcastSwitcher activePodcast={podcastStore.activePodcast} />
    </header>
  );
};

export default DashHeader;
