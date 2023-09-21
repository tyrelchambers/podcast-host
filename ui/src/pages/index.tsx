import EpisodesTable from "@/components/EpisodesTable";
import { useUserStore } from "@/hooks/stores/userStore";
import { useUserQuery } from "@/hooks/api/useUserQuery";
import Header from "@/layouts/Header";
import axios from "axios";
import { Suspense } from "react";
import { useEpisodesQuery } from "@/hooks/api/useEpisodesQuery";

export default function Home() {
  const user = useUserStore((state) => state.user);
  const episodes = useEpisodesQuery(user?.id);

  return (
    <main className="w-full">
      <Header />

      <h1 className="h1">Episodes</h1>
      <EpisodesTable episodes={episodes.data} />
    </main>
  );
}
