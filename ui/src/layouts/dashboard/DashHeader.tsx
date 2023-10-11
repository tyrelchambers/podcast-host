import DashNav from "@/components/dashboard/DashNav";
import { usePodcastStore } from "@/hooks/stores/podcastStore";
import { faArrowLeft } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import Link from "next/link";
import React from "react";

interface Props {
  rootPath: string;
}

const DashHeader = ({ rootPath }: Props) => {
  const { activePodcast } = usePodcastStore();

  return (
    <header className="bg-card h-full  flex flex-col">
      <header className="p-4">
        <p className="text-card-foreground">Resonate</p>
      </header>
      <DashNav rootPath={rootPath} />
      <div className="bg-primary/10 p-4">
        <p className="text-card-foreground/70 text-sm italic">
          Viewing podcast:
        </p>
        <p className="text-card-foreground font-bold text-xl">
          {activePodcast?.title}
        </p>
      </div>
      <Link
        href="/"
        className="text-card-foreground text-sm hover:bg-primary/10 p-4 w-full"
      >
        <FontAwesomeIcon className="mr-2" icon={faArrowLeft} />
        Back to podcasts
      </Link>
    </header>
  );
};

export default DashHeader;
