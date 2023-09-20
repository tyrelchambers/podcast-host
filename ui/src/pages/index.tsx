import EpisodesTable from "@/components/EpisodesTable";
import { useUserQuery } from "@/hooks/useUserQuery";
import Header from "@/layouts/Header";
import axios from "axios";
import { Suspense } from "react";

const getAllEpisodes = async (id: string) => {
  const data = await axios
    .get(`http://localhost:8080/api/user/${id}/episodes`)
    .then((res) => res.data)
    .catch((err) => {
      console.log(err);
    });

  return data;
};

export default function Home() {
  const user = useUserQuery();
  // const episodes = await getAllEpisodes(user.id);

  console.log(user);

  return (
    <main className="w-full">
      <Header />

      <h1 className="h1">Episodes</h1>
      {/* <EpisodesTable episodes={episodes} /> */}
    </main>
  );
}
