import EpisodeEditBody from "@/components/EpisodeEditBody";
import { useEpisodeQuery } from "@/hooks/api/useEpisodeQuery";
import { Episode } from "@/lib/types";
import axios from "axios";
import { useRouter } from "next/router";
import React from "react";

const Page = () => {
  const router = useRouter();

  const episode = useEpisodeQuery(router.query.id as string);

  return (
    <main className="max-w-screen-lg mx-auto py-10">
      <h1 className="h1">Edit episode</h1>
      <EpisodeEditBody episode={episode.data} />
    </main>
  );
};

export default Page;
