"use client";

import React, { useRef } from "react";

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

const MAX_FILE_SIZE = 500000;
const ACCEPTED_IMAGE_TYPES = [
  "image/jpeg",
  "image/jpg",
  "image/png",
  "image/webp",
];

const Page = () => {
  const router = useRouter();
  const fileUploadRef = useRef<HTMLInputElement>(null);

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      file: "",
      title: "",
      author: "",
      keywords: "",
      episodeNumber: "0",
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
    const formData = new FormData();
    const file = fileUploadRef.current?.files?.[0];
    const description = JSON.stringify(editor?.getJSON()) ?? "";

    if (!file) {
      return;
    }

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

    formData.append("file", file);
    formData.append("title", data.title);
    formData.append("description", description);
    formData.append("author", data.author);
    formData.append("keywords", data.keywords);
    formData.append("episodeNumber", data.episodeNumber.toString());
    formData.append("publishDate", getDate().toString());

    await axios.post("http://localhost:8080/api/episode/create", formData, {
      headers: {
        "Content-Type": "multipart/form-data",
      },
      withCredentials: true,
    });
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
        />
      </section>
    </DashLayout>
  );
};

export default Page;
