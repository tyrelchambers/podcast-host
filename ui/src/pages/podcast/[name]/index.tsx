import DashCard from "@/components/dashboard/DashCard";
import { usePodcastQuery } from "@/hooks/api/usePodcastQuery";
import DashHeader from "@/layouts/dashboard/DashHeader";
import DashLayout from "@/layouts/dashboard/DashLayout";
import { useRouter } from "next/router";
import React from "react";

const Podcast = () => {
  const router = useRouter();
  const nameParam = router.query.name;

  const podcast = usePodcastQuery(nameParam as string).data;

  console.log(podcast);

  // const latestUpload = podcast?.episodes[podcast.episodes?.length - 1];

  return (
    <DashLayout
      leftCol={<DashHeader rootPath={router.asPath} />}
      rightCol={<p>hey over here</p>}
    >
      <h1 className="h1">{podcast?.title}</h1>

      <section>
        <DashCard />
      </section>
    </DashLayout>
  );
};

export default Podcast;
