"use client";
import { Episode, formSchema } from "@/lib/types";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import EpisodeForm, { SubmitHandlerProps } from "../forms/EpisodeForm";
import { Suspense, useRef, useState } from "react";
import { z } from "zod";
import { getUnixTime } from "date-fns";
import axios from "axios";
import { useToast } from "./ui/use-toast";
import { useRouter } from "next/navigation";

const EpisodeEditBody = ({ episode }: { episode: Episode | undefined }) => {
  const fileUploadRef = useRef<HTMLInputElement>(null);
  const [uploadProgress, setUploadProgress] = useState(0);

  const { toast } = useToast();
  const router = useRouter();
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
            Number(data.schedule_hour),
            Number(data.schedule_minute),
            0
          )
        );

        return date;
      }

      return getUnixTime(publishDate);
    };

    await axios
      .postForm(
        `http://localhost:8080/api/episode/${episode?.id}/edit`,
        {
          id: episode?.id,
          title: data.title,
          description,
          author: data.author,
          keywords: data.keywords,
          episodeNumber: data.episode_number,
          url: episode?.url,
          publishDate: episode?.publish_date,
          podcastId: episode?.podcast_id,
          file,
        },
        {
          headers: {
            "Content-Type": "multipart/form-data",
          },
          withCredentials: true,
          onUploadProgress: (progressEvent) => {
            if (file && progressEvent.total) {
              setUploadProgress(
                Math.round((progressEvent.loaded * 100) / progressEvent.total)
              );
            }
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

  const deleteHandler = (id: string | undefined) => {
    if (!id) return;

    axios
      .delete(`http://localhost:8080/api/episode/${id}/delete`)
      .then(() => {
        toast({
          description: "Episode deleted",
        });
        router.push("/");
      })
      .catch((err) => {
        toast({
          title: "Awe, man! Something went wrong.",
          description: err.message,
        });
      });
  };

  return (
    <section className="section-card">
      <EpisodeForm
        form={form}
        episode={episode}
        submitHandler={submitHandler}
        fileUploadRef={fileUploadRef}
        ctaText="Edit episode"
        deleteHandler={deleteHandler}
        isEditing
        uploadProgress={uploadProgress}
      />
    </section>
  );
};

export default EpisodeEditBody;
