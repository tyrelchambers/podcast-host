"use client";

import React, { useRef, useState } from "react";

import EpisodeForm, { SubmitHandlerProps } from "@/forms/EpisodeForm";
import { getUnixTime } from "date-fns";
import { useForm } from "react-hook-form";
import { formSchema } from "@/lib/types";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import axios from "axios";
import DashLayout from "@/layouts/dashboard/DashLayout";
import DashHeader from "@/layouts/dashboard/DashHeader";
import { useRouter } from "next/router";
import { usePodcastStore } from "@/hooks/stores/podcastStore";
import { useMiscInfoQuery } from "@/hooks/api/useMiscInfoQuery";

const Page = () => {
  const router = useRouter();
  const [uploadProgress, setUploadProgress] = useState(0);
  const fileUploadRef = useRef<HTMLInputElement>(null);
  const podcastStore = usePodcastStore();
  const podcast = podcastStore.activePodcast;
  const miscInfo = useMiscInfoQuery(podcast?.uuid ?? "");

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    values: {
      episode_number: miscInfo.data?.next_episode_number
        ? miscInfo.data.next_episode_number
        : "0",
      schedule_hour: "12",
      schedule_minute: "00",
      schedule_meridiem: "PM",
      explicit_content: false,
      description: "",
      keywords: "",
      author: "",
      title: "",
      draft: false,
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
            Number(data.schedule_hour),
            Number(data.schedule_minute),
            0
          )
        );

        return date;
      }

      return getUnixTime(publishDate);
    };

    await axios.postForm(
      "http://localhost:8080/api/episode/create",
      {
        file: file ?? "",
        title: data.title,
        description: description,
        author: data.author,
        keywords: data.keywords,
        episodeNumber: data.episode_number,
        publishDate: getDate().toString(),
        podcastId: podcast?.uuid ?? "",
        draft: data.draft,
        isScheduled: whenToPublish === "schedule",
      },
      {
        withCredentials: true,
        onUploadProgress: (progressEvent) => {
          if (file && progressEvent.total) {
            setUploadProgress(
              Math.round((progressEvent.loaded * 100) / progressEvent.total)
            );
          }
        },
      }
    );
  };

  return (
    <DashLayout leftCol={<DashHeader rootPath={router.query.name as string} />}>
      <h1 className="h1">Create your episode</h1>
      <p className="font-light mt-3">
        Creating an episode for{" "}
        <span className="underline italic font-bold">
          {podcastStore.activePodcast?.title}
        </span>
      </p>
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
