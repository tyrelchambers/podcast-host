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
import { usePodcastQuery } from "@/hooks/api/usePodcastQuery";
import DashHeader from "@/layouts/dashboard/DashHeader";
import DashLayout from "@/layouts/dashboard/DashLayout";
import Link from "next/link";
import { useRouter } from "next/router";
import React from "react";

const Podcast = () => {
  const router = useRouter();
  const nameParam = router.query.name;
  const { data } = usePodcastQuery(nameParam as string);

  const podcast = data?.podcast;
  const latestEpisodes = data?.latestEpisode;

  console.log(podcast);

  // const latestUpload = podcast?.episodes[podcast.episodes?.length - 1];

  return (
    <DashLayout
      leftCol={<DashHeader rootPath={router.asPath} />}
      rightCol={<p>hey over here</p>}
    >
      <h1 className="h1">{podcast?.title}</h1>

      <section className="border-border border-[1px] p-4 rounded-xl mt-6">
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
