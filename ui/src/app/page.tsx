import EpisodesTable from "@/components/EpisodesTable";
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

const getCurrentUser = async () => {
  const data = await axios
    .get("http://localhost:8080/api/user/me")
    .then((res) => res.data)
    .catch((err) => {
      console.log(err);
    });

  return data;
};

export default async function Home() {
  const user = await getCurrentUser();
  const episodes = await getAllEpisodes(user.id);

  return (
    <main className="w-full">
      <Header />

      <Suspense fallback={<div>Loading...</div>}>
        <h1 className="h1">Episodes</h1>
        <EpisodesTable episodes={episodes} />
      </Suspense>
    </main>
  );
}
