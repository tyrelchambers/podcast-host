import { usePodcastsQuery } from "@/hooks/api/usePodcastsQuery";
import { usePodcastStore } from "@/hooks/stores/podcastStore";
import { useUserStore } from "@/hooks/stores/userStore";
import React, { useEffect } from "react";

interface Props {
  leftCol: React.ReactNode;
  children: React.ReactNode | React.ReactNode[];
  rightCol: React.ReactNode;
}
const DashLayout = ({ leftCol, children, rightCol }: Props) => {
  const user = useUserStore((state) => state.user);

  const podcastsQuery = usePodcastsQuery(user?.id);
  const podcastStore = usePodcastStore();
  const podcasts = podcastsQuery.data;

  useEffect(() => {
    if (podcasts && podcasts.length) {
      podcastStore.setPodcasts(podcasts);
    }
  }, [podcasts]);

  return (
    <div className="dash-layout-grid h-screen">
      <div className=" w-[250px] h-full">{leftCol}</div>
      <section className="flex-1 h-full p-8">{children}</section>
      <div className=" w-[250px] h-full">{rightCol}</div>
    </div>
  );
};

export default DashLayout;
