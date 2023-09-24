import { usePodcastsQuery } from "@/hooks/api/usePodcastsQuery";
import { usePodcastStore } from "@/hooks/stores/podcastStore";
import { useUserStore } from "@/hooks/stores/userStore";
import { useRouter } from "next/router";
import React, { useEffect } from "react";

interface Props {
  leftCol: React.ReactNode;
  children: React.ReactNode | React.ReactNode[];
  rightCol: React.ReactNode;
}
const DashLayout = ({ leftCol, children, rightCol }: Props) => {
  const router = useRouter();
  const nameParam = router.query.name;

  const podcastStore = usePodcastStore();

  return !nameParam ? null : (
    <div className="dash-layout-grid h-screen ">
      <div className=" w-[250px] h-full">{leftCol}</div>
      <section className="flex-1 h-full p-8 overflow-y-auto">
        {children}
      </section>
      <div className=" w-[250px] h-full">{rightCol}</div>
    </div>
  );
};

export default DashLayout;
