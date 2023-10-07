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
import { usePodcastQuery } from "@/hooks/api/usePodcastQuery";
import { usePodcastStore } from "@/hooks/stores/podcastStore";
import DashHeader from "@/layouts/dashboard/DashHeader";
import DashLayout from "@/layouts/dashboard/DashLayout";
import { faArrowRight, faRss } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import Link from "next/link";
import { useRouter } from "next/router";
import React, { useEffect } from "react";

const Podcast = () => {
  const router = useRouter();
  const nameParam = router.query.name;
  const podcastStore = usePodcastStore();
  const miscInfo = useMiscInfoQuery(podcastStore.activePodcast?.uuid ?? "");

  const podcast = podcastStore?.activePodcast;

  useEffect(() => {
    if (router.query.name) {
      podcastStore.setActivePodcast(nameParam as string);
    }
  }, [nameParam, router.query.name]);

  return (
    <DashLayout leftCol={<DashHeader rootPath={router.query.name as string} />}>
      <h1 className="h1">{podcast?.title}</h1>

      <section className="section-card flex gap-2">
        <Link href={`/podcast/${nameParam}/episode/create`}>
          <Button className="w-full">Create episode</Button>
        </Link>
        <Dialog>
          <DialogTrigger>
            <Button variant="secondary" className="w-full">
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
    </DashLayout>
  );
};

export default Podcast;
