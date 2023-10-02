import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import EditPodcast from "@/forms/EditPodcast";
import DashHeader from "@/layouts/dashboard/DashHeader";
import DashLayout from "@/layouts/dashboard/DashLayout";
import ShowSettings from "@/layouts/tabPanels/ShowSettings";
import { useRouter } from "next/router";
import React from "react";

const Index = () => {
  const router = useRouter();

  return (
    <DashLayout leftCol={<DashHeader rootPath={router.query.name as string} />}>
      <h1 className="h1">Settings</h1>
      <Tabs defaultValue="account" className="my-10">
        <TabsList>
          <TabsTrigger value="show_settings">Show settings</TabsTrigger>
          <TabsTrigger value="rss_feed">RSS feed</TabsTrigger>
          <TabsTrigger value="embeds">Embeds</TabsTrigger>
          <TabsTrigger value="social_media">Social media</TabsTrigger>
        </TabsList>
        <TabsContent value="show_settings">
          <ShowSettings />
        </TabsContent>
        <TabsContent value="password">Change your password here.</TabsContent>
      </Tabs>
    </DashLayout>
  );
};

export default Index;
