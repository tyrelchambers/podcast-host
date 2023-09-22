import { usePodcastQuery } from "@/hooks/api/usePodcastQuery";
import { useRouter } from "next/router";
import React from "react";

const Podcast = () => {
  const router = useRouter();
  const nameParam = router.query.name;

  const podcast = usePodcastQuery(nameParam as string);

  console.log(podcast);

  return <div>index</div>;
};

export default Podcast;
