import EpisodesTable from "@/components/EpisodesTable";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { dashboardRoot } from "@/constants";
import { useEpisodesQuery } from "@/hooks/api/useEpisodesQuery";
import { usePodcastStore } from "@/hooks/stores/podcastStore";
import DashHeader from "@/layouts/dashboard/DashHeader";
import DashLayout from "@/layouts/dashboard/DashLayout";
import { faSearch } from "@fortawesome/pro-regular-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import Link from "next/link";
import { useRouter } from "next/router";
import React from "react";

const Episodes = () => {
  const router = useRouter();
  const nameParam = router.query.name;
  const activePodcast = usePodcastStore((state) => state.activePodcast);
  const episodes = useEpisodesQuery(activePodcast?.uuid);

  return (
    <DashLayout leftCol={<DashHeader rootPath={router.query.name as string} />}>
      <h1 className="h1">Episodes</h1>

      <section className="flex mt-10 justify-between gap-4">
        <Input
          type="search"
          placeholder="Search for episodes"
          className="flex-1"
          icon={faSearch}
        />
        <Link href={`/podcast/${nameParam}/episode/create`}>
          <Button>Create episode</Button>
        </Link>
      </section>

      <section className="bg-card p-4 rounded-xl shadow-sm mt-4">
        <EpisodesTable
          episodes={episodes.data}
          podcastName={nameParam as string}
        />
      </section>
    </DashLayout>
  );
};

export default Episodes;
