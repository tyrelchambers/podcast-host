import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { useMiscInfoQuery } from "@/hooks/api/useMiscInfoQuery";
import { usePodcastStore } from "@/hooks/stores/podcastStore";
import DashHeader from "@/layouts/dashboard/DashHeader";
import DashLayout from "@/layouts/dashboard/DashLayout";
import { faArrowRight, faRss } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import Link from "next/link";
import { useRouter } from "next/router";
import React from "react";
import { format, fromUnixTime } from "date-fns";
import { Card, CardDescription, CardTitle } from "@/components/ui/card";
import {
  faClock,
  faDownload,
  faListMusic,
  faPenRuler,
} from "@fortawesome/pro-duotone-svg-icons";

const Podcast = () => {
  const router = useRouter();
  const nameParam = router.query.name;
  const podcastStore = usePodcastStore();
  const miscInfo = useMiscInfoQuery(podcastStore.activePodcast?.uuid ?? "");

  const podcast = podcastStore?.activePodcast;

  return (
    <DashLayout leftCol={<DashHeader rootPath={router.query.name as string} />}>
      <h1 className="font-bold text-foreground flex-1 text-3xl mb-8">
        {podcast?.title}
      </h1>

      <section className="bg-card p-4 rounded-xl w-full flex gap-3 mb-8">
        <Link href={`/podcast/${nameParam}/episode/create`}>
          <Button className="w-full">Create episode</Button>
        </Link>
        <Dialog>
          <DialogTrigger asChild>
            <Button variant="secondary">
              <FontAwesomeIcon icon={faRss} className="mr-3" />
              RSS feed
            </Button>
          </DialogTrigger>
          <DialogContent>
            <DialogHeader>
              <DialogTitle>
                <FontAwesomeIcon icon={faRss} className="mr-3" />
                Your RSS feed
              </DialogTitle>
              <DialogDescription>
                An RSS feed for a podcast is used to syndicate and distribute
                episodes of the podcast to various platforms and apps, allowing
                subscribers to automatically receive new episodes as they are
                released.
              </DialogDescription>
            </DialogHeader>
            <DialogFooter className="flex !flex-col gap-2">
              <div className="flex gap-2">
                <Input
                  value="https://feeds.transistor.fm/the-midnight-podcast"
                  readOnly
                />
                <Button variant="outline">Copy</Button>
              </div>
              <DialogDescription>
                <Link
                  href="https://feeds.transistor.fm/the-midnight-podcast"
                  className="flex items-center"
                >
                  Configure your RSS feed settings{" "}
                  <FontAwesomeIcon icon={faArrowRight} className="ml-2" />
                </Link>
              </DialogDescription>
            </DialogFooter>
          </DialogContent>
        </Dialog>
      </section>
      <section className="flex gap-4 mb-10 h-[350px] flex-1 w-full">
        <div className="flex flex-col gap-3 w-96 h-full">
          <Card className="flex-1 flex justify-between items-center p-4 h-full">
            <CardTitle className="text-lg">
              <FontAwesomeIcon icon={faDownload} className="mr-2" />
              Total downloads
            </CardTitle>
            <CardDescription>Card Description</CardDescription>
          </Card>

          <Card className="flex-1 flex justify-between items-center p-4 h-full">
            <CardTitle className="text-lg">
              <FontAwesomeIcon icon={faPenRuler} className="mr-2" /> Drafts
            </CardTitle>
            <CardDescription className="text-3xl">
              {miscInfo.data?.draft_count}
            </CardDescription>
          </Card>

          <Card className="flex-1 flex justify-between items-center p-4 h-full">
            <CardTitle className="text-lg">
              <FontAwesomeIcon
                icon={faClock}
                style={{ ["--fa-secondary-opacity" as string]: "0.1" }}
                className="mr-2"
              />{" "}
              Scheduled
            </CardTitle>
            <CardDescription>Card Description</CardDescription>
          </Card>

          <Card className="flex-1 flex justify-between items-center p-4 h-full">
            <CardTitle className="text-lg">
              <FontAwesomeIcon icon={faListMusic} className="mr-2" /> Published
              episodes
            </CardTitle>
            <CardDescription className="text-3xl">
              {miscInfo.data?.episode_count}
            </CardDescription>
          </Card>
        </div>

        <section className="relative bg-overlay flex h-full flex-1 shadow-lg rounded-3xl overflow-hidden">
          <div className="absolute inset-0 z-10 p-10 h-[245px] flex flex-col top-[50%] translate-y-[-50%] ">
            <div className="flex flex-col flex-1">
              <h3 className="text-4xl font-medium text-background-alt-foreground">
                {miscInfo.data?.latest_episode.title}
              </h3>
              <p className="text-background-alt-foreground/60">
                Published on{" "}
                {miscInfo.data?.latest_episode.publish_date &&
                  format(
                    fromUnixTime(miscInfo.data?.latest_episode.publish_date),
                    "MMM d, yyyy"
                  )}
              </p>
            </div>

            <Link
              href={`/podcast/${nameParam}/episode/${miscInfo.data?.latest_episode.uuid}`}
            >
              <Button variant="outline" className="w-fit">
                View episode
              </Button>
            </Link>
          </div>
        </section>
      </section>
    </DashLayout>
  );
};

export default Podcast;
