import DashCard from "@/components/dashboard/DashCard";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { useMiscInfoQuery } from "@/hooks/api/useMiscInfoQuery";
import { usePodcastQuery } from "@/hooks/api/usePodcastQuery";
import { usePodcastStore } from "@/hooks/stores/podcastStore";
import DashHeader from "@/layouts/dashboard/DashHeader";
import DashLayout from "@/layouts/dashboard/DashLayout";
import Link from "next/link";
import { useRouter } from "next/router";
import React, { useEffect } from "react";

const Podcast = () => {
  const router = useRouter();
  const nameParam = router.query.name;
  const { data } = usePodcastQuery(nameParam as string);
  const podcastStore = usePodcastStore();
  const miscInfo = useMiscInfoQuery(podcastStore.activePodcast?.id ?? "");

  const podcast = data?.podcast;
  const latestEpisodes = data?.latestEpisode;

  useEffect(() => {
    if (router.query.name) {
      podcastStore.setActivePodcast(nameParam as string);
    }
  }, [nameParam, router.query.name]);

  return (
    <DashLayout
      leftCol={<DashHeader rootPath={router.query.name as string} />}
      rightCol={<p>hey over here</p>}
    >
      <h1 className="h1">{podcast?.title}</h1>

      <section className="section-card">
        <Link href={`/podcast/${nameParam}/episode/create`}>
          <Button>Create episode</Button>
        </Link>
      </section>

      <section className="my-6">
        <Card>
          <CardHeader>
            <CardTitle>Latest episode analytics</CardTitle>
            <CardDescription>{latestEpisodes?.title}</CardDescription>
          </CardHeader>
          <CardContent>
            <p>Card Content</p>
          </CardContent>
          <CardFooter>
            <p>Card Footer</p>
          </CardFooter>
        </Card>
      </section>
    </DashLayout>
  );
};

export default Podcast;
