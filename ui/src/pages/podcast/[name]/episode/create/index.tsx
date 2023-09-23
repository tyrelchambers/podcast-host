"use client";

import React, { useRef, useState } from "react";

import EpisodeForm, { SubmitHandlerProps } from "@/components/edit/EpisodeForm";
import { getUnixTime } from "date-fns";
import { useForm } from "react-hook-form";
import { formSchema } from "@/lib/types";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import axios from "axios";
import DashLayout from "@/layouts/dashboard/DashLayout";
import DashHeader from "@/layouts/dashboard/DashHeader";
import { dashboardRoot } from "@/constants";
import { useRouter } from "next/router";
import { usePodcastStore } from "@/hooks/stores/podcastStore";
import { useMiscInfoQuery } from "@/hooks/api/useMiscInfoQuery";

const MAX_FILE_SIZE = 500000;
const ACCEPTED_IMAGE_TYPES = [
  "image/jpeg",
  "image/jpg",
  "image/png",
  "image/webp",
];

const Page = () => {
  const router = useRouter();
  const [uploadProgress, setUploadProgress] = useState(0);
  const [isUploading, setIsUploading] = useState(false);
  const fileUploadRef = useRef<HTMLInputElement>(null);
  const podcastStore = usePodcastStore();
  const podcast = podcastStore.findPodcast(router.query.name as string);
  const miscInfo = useMiscInfoQuery(podcast?.id ?? "");

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      file: "",
      title: "",
      author: "",
      keywords: "",
      episodeNumber: String(miscInfo.data?.nextEpisodeNumber) ?? "1",
      description: "",
      scheduleHour: "12",
      scheduleMinute: "00",
      scheduleMeridiem: "PM",
    },
  });

  const submitHandler = async ({
    data,
    editor,
    publishDate,
    whenToPublish,
  }: SubmitHandlerProps) => {
    const file = fileUploadRef.current?.files?.[0];
    const description = JSON.stringify(editor?.getJSON()) ?? "";

    const getDate = () => {
      if (whenToPublish === "schedule") {
        const date = getUnixTime(
          new Date(
            publishDate.getUTCFullYear(),
            publishDate.getUTCMonth(),
            publishDate.getUTCDate(),
            Number(data.scheduleHour),
            Number(data.scheduleMinute),
            0
          )
        );

        return date;
      }

      return getUnixTime(publishDate);
    };

    setIsUploading(true);

    await axios.postForm(
      "http://localhost:8080/api/episode/create",
      {
        file: file ?? "",
        title: data.title,
        description: description,
        author: data.author,
        keywords: data.keywords,
        episodeNumber: data.episodeNumber,
        publishDate: getDate().toString(),
        podcastId: podcast?.id ?? "",
      },
      {
        withCredentials: true,
        onUploadProgress: (progressEvent) => {
          if (progressEvent.total) {
            setUploadProgress(
              Math.round((progressEvent.loaded * 100) / progressEvent.total)
            );
          }

          if (progressEvent.progress === 1) {
            setIsUploading(false);
          }
        },
      }
    );
  };

  return (
    <DashLayout
      leftCol={<DashHeader rootPath={dashboardRoot(router.asPath)} />}
      rightCol={<p>hey over here</p>}
    >
      <h1 className="h1">Create your episode</h1>
      <section className="section-card">
        <EpisodeForm
          form={form}
          fileUploadRef={fileUploadRef}
          submitHandler={submitHandler}
          ctaText="Create episode"
          uploadProgress={uploadProgress}
        />
      </section>
    </DashLayout>
  );
};

export default Page;
