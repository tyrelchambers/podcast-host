import DashHeader from "@/layouts/dashboard/DashHeader";
import DashLayout from "@/layouts/dashboard/DashLayout";
import { useRouter } from "next/router";
import React from "react";

const Rss = () => {
  const router = useRouter();

  return (
    <DashLayout leftCol={<DashHeader rootPath={router.query.name as string} />}>
      rss
    </DashLayout>
  );
};

export default Rss;
