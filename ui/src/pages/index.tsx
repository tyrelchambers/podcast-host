import PodcastList from "@/components/PodcastList";
import { Button } from "@/components/ui/button";
import { usePodcastsQuery } from "@/hooks/api/usePodcastsQuery";
import { usePodcastStore } from "@/hooks/stores/podcastStore";
import { useUserStore } from "@/hooks/stores/userStore";
import Header from "@/layouts/Header";
import Link from "next/link";
import { useEffect } from "react";

export default function Home() {
  const user = useUserStore((state) => state.user);

  return (
    <main className="w-full">
      <Header />

      <section className="max-w-screen-2xl mx-auto my-20 ">
        <header className="flex items-center justify-between gap-3 mb-8">
          <h1 className="h1">Your podcasts</h1>
          <Link href="/podcast/add">
            <Button>Create podcast</Button>
          </Link>
        </header>
        <section className="p-8 rounded-xl bg-card shadow-sm flex flex-col">
          {/* <PodcastList podcasts={podcasts} /> */}
        </section>
      </section>
    </main>
  );
}
