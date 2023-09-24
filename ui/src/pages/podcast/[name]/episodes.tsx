import EpisodesTable from "@/components/EpisodesTable";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { dashboardRoot } from "@/constants";
import { useEpisodesQuery } from "@/hooks/api/useEpisodesQuery";
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
  const episodes = useEpisodesQuery(nameParam as string);

  return (
    <DashLayout
      leftCol={<DashHeader rootPath={router.query.name as string} />}
      rightCol={<p>hey over here</p>}
    >
      <h1 className="h1">Episodes</h1>

      <section className="flex mt-10 justify-between gap-4">
        <div className="flex items-center gap-2 flex-1">
          <FontAwesomeIcon icon={faSearch} />
          <Input
            type="search"
            placeholder="Search for episodes"
            className="flex-1"
          />
        </div>
        <Link href={`/podcast/${nameParam}/episode/create`}>
          <Button>Create episode</Button>
        </Link>
      </section>

      <section className="bg-card p-4 rounded-xl shadow-sm mt-4">
        <EpisodesTable episodes={episodes.data} />
      </section>
    </DashLayout>
  );
};

export default Episodes;
