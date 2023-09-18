"use client";
import { Episode, formSchema } from "@/lib/types";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import EpisodeForm, { SubmitHandlerProps } from "./edit/EpisodeForm";
import { Suspense, useRef } from "react";
import { z } from "zod";
import { getUnixTime } from "date-fns";
import axios from "axios";
import { useToast } from "./ui/use-toast";

const EpisodeEditBody = ({ episode }: { episode: Episode }) => {
  const fileUploadRef = useRef<HTMLInputElement>(null);
  const { toast } = useToast();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    values: episode,
    shouldFocusError: true,
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

    await axios
      .putForm(
        `http://localhost:8080/api/episode/${data.id}/edit`,
        {
          id: episode.id,
          title: data.title,
          description,
          author: data.author,
          keywords: data.keywords,
          episodeNumber: data.episodeNumber,
          publishDate: getDate().toString(),
          url: episode.url,
          file,
        },
        {
          headers: {
            "Content-Type": "multipart/form-data",
          },
        }
      )
      .then(() => {
        toast({
          description: "Episode updated",
        });
      })
      .catch((err) => {
        toast({
          title: "Awe, man! Something went wrong.",
          description: err.message,
        });
      });
  };

  return (
    <Suspense fallback={<div>Loading...</div>}>
      <EpisodeForm
        form={form}
        episode={episode}
        submitHandler={submitHandler}
        fileUploadRef={fileUploadRef}
        ctaText="Edit episode"
      />
    </Suspense>
  );
};

export default EpisodeEditBody;
