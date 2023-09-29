import EpisodeEditBody from "@/components/EpisodeEditBody";
import { useEpisodeQuery } from "@/hooks/api/useEpisodeQuery";
import DashHeader from "@/layouts/dashboard/DashHeader";
import DashLayout from "@/layouts/dashboard/DashLayout";
import { Episode } from "@/lib/types";
import axios from "axios";
import { useRouter } from "next/router";
import React from "react";

const Page = () => {
  const router = useRouter();

  const episode = useEpisodeQuery(router.query.id as string);

  return (
    <DashLayout leftCol={<DashHeader rootPath={router.query.name as string} />}>
      <h1 className="h1">Edit episode</h1>
      <EpisodeEditBody episode={episode.data} />
    </DashLayout>
  );
};

export default Page;
