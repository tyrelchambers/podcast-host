import EpisodeEditBody from "@/components/EpisodeEditBody";
import { Episode } from "@/lib/types";
import axios from "axios";
import React from "react";

const getEpisodeById = async (id: string): Promise<Episode> => {
  return await axios
    .get("http://localhost:8080/api/episode/" + id)
    .then((res) => res.data)
    .catch((err) => console.log(err));
};

const Page = async ({ params }: { params: { id: string } }) => {
  const episode = await getEpisodeById(params.id);

  return (
    <main className="max-w-screen-lg mx-auto py-10">
      <h1 className="h1">Edit episode</h1>
      <EpisodeEditBody episode={episode} />
    </main>
  );
};

export default Page;
