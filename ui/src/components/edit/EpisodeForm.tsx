"use client";
import { Episode, formSchema } from "@/lib/types";
import React, { useEffect, useState } from "react";
import { Form, FormField, FormItem } from "../ui/form";
import { Label } from "../ui/label";
import { Input } from "../ui/input";
import { UseFormReturn } from "react-hook-form";
import { Editor, useEditor } from "@tiptap/react";
import StarterKit from "@tiptap/starter-kit";
import Underline from "@tiptap/extension-underline";
import Link from "@tiptap/extension-link";
import Placeholder from "@tiptap/extension-placeholder";
import TextAlign from "@tiptap/extension-text-align";
import { TextEditor } from "../TextEditor";
import ScheduleHourSelector from "../ScheduleHourSelector";
import ScheduleMinuteSelector from "../ScheduleMinuteSelector";
import ScheduleMeridiemSelector from "../ScheduleMeridiemSelector";
import { formatBytes } from "@/lib/utils";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faCheckCircle,
  faCircleExclamation,
  faCloudArrowUp,
  faMusic,
  faSpinner,
} from "@fortawesome/free-solid-svg-icons";
import PublishSelector from "../PublishSelector";
import DatePicker from "../DatePicker";
import { Button } from "../ui/button";
import { z } from "zod";
import { format, fromUnixTime } from "date-fns";
import { Badge } from "../ui/badge";
import { Checkbox } from "../ui/checkbox";

// 1GB in bytes
const MAX_FILE_SIZE = 1073741824;
const ACCEPTED_FILE_TYPES = ["audio/mpeg", "audio/ogg", "audio/wav"];

export interface SubmitHandlerProps {
  data: z.infer<typeof formSchema> | Episode;
  editor: Editor | null;
  publishDate: Date;
  whenToPublish: "immediate" | "schedule";
}

interface Props {
  episode?: Episode;
  form: UseFormReturn<z.infer<typeof formSchema>, any, undefined>;
  submitHandler: (props: SubmitHandlerProps) => void;
  fileUploadRef: React.RefObject<HTMLInputElement>;
  ctaText: string;
  isEditing?: boolean;
  deleteHandler?: (id: string | undefined) => void;
  uploadProgress: number;
}

const EpisodeForm = ({
  episode,
  form,
  submitHandler,
  fileUploadRef,
  ctaText,
  isEditing = false,
  deleteHandler,
  uploadProgress,
}: Props) => {
  if (isEditing && !deleteHandler) {
    throw new Error("EpisodeForm must have a deleteHandler");
  }

  const [uploadError, setUploadError] = useState<string>("");
  const [isSubmitting, setIsSubmitting] = useState(false);
  const editor = useEditor({
    extensions: [
      StarterKit,
      Underline,
      Link,
      Placeholder.configure({
        placeholder:
          "Add in your shownotes here. A portion of your notes will show up as a summary.",
      }),
      TextAlign.configure({ types: ["heading", "paragraph"] }),
    ],
    content: JSON.parse(episode?.description || "{}"),
  });

  const [whenToPublish, setWhenToPublish] = useState<"immediate" | "schedule">(
    "immediate"
  );
  const [publishDate, setPublishDate] = useState<Date>(new Date(Date.now()));

  // const [changePublishDate, setChangePublishDate] = useState(false);

  useEffect(() => {
    if (fileUploadRef.current && fileUploadRef.current?.files) {
      if (fileUploadRef.current?.files[0]?.size > MAX_FILE_SIZE) {
        setUploadError("File size must be less than 1GB");
      } else {
        setUploadError("");
      }
    }
  }, [fileUploadRef.current]);

  const fileName = fileUploadRef.current?.files?.[0]?.name;
  const fileSize = fileUploadRef.current?.files?.[0]?.size
    ? formatBytes(fileUploadRef.current?.files?.[0]?.size)
    : null;
  const fileType = fileUploadRef.current?.files?.[0]?.type;

  const watchFileValue = form.watch("file");

  const errors = form.formState.errors;

  const submit = async (data: z.infer<typeof formSchema>) => {
    setIsSubmitting(true);
    await submitHandler({ data, editor, publishDate, whenToPublish });
    setIsSubmitting(false);
  };

  return (
    <Form {...form}>
      <form className="flex flex-col gap-6 ">
        <FormField
          name="title"
          render={({ field }) => (
            <FormItem>
              <Label htmlFor="title">Episode title</Label>
              <Input placeholder="Your episode title" {...field} />
              {errors.title && (
                <p className="text-red-500">{errors.title.message}</p>
              )}
            </FormItem>
          )}
        />

        {episode?.url && (
          <FormItem>
            <Label>Current audio file</Label>
            <div className="audio-player">
              <audio controls className="w-full">
                <source src={episode?.url} type="audio/mpeg" />
                Your browser does not support the audio tag.
              </audio>
            </div>
          </FormItem>
        )}

        <FormField
          name="file"
          render={({ field }) => (
            <FormItem>
              <Label htmlFor={field.name}>
                Upload new audio file
                <div className="relative">
                  <div className="w-full z-10 relative border-2 border-dashed border-border p-4 rounded-md h-[100px] flex items-center px-10 mt-2 hover:border-primary transition-all">
                    {!watchFileValue ? (
                      <NoFileSelected />
                    ) : (
                      <div className="flex items-center gap-4 z-0 w-full">
                        <FontAwesomeIcon icon={faMusic} className="text-3xl" />
                        <div className="flex flex-col flex-1">
                          <p className="font-medium mb-2">{fileName}</p>
                          <p className="text-muted-foreground text-sm font-light">
                            {fileSize} - {fileType}
                          </p>
                          {uploadError && (
                            <Badge
                              variant="destructive"
                              className="font-light w-fit flex gap-2 mt-2"
                            >
                              <FontAwesomeIcon icon={faCircleExclamation} />
                              {uploadError}
                            </Badge>
                          )}
                        </div>
                        {uploadProgress > 0 && (
                          <Badge variant="default">{uploadProgress}%</Badge>
                        )}
                      </div>
                    )}
                    <Input
                      type="file"
                      className="hidden"
                      accept={ACCEPTED_FILE_TYPES.join(",")}
                      id={field.name}
                      {...field}
                      ref={fileUploadRef}
                    />
                  </div>
                  <div
                    className="absolute top-0 z-10 h-full bg-green-600 mix-blend-screen rounded-md transition-all"
                    style={{ width: `${uploadProgress}%` }}
                  ></div>
                  <div
                    className="absolute top-0 z-0 h-full bg-green-100  rounded-md transition-all"
                    style={{ width: `${uploadProgress}%` }}
                  ></div>
                </div>
              </Label>
            </FormItem>
          )}
        />

        <FormField
          name="description"
          render={() => (
            <FormItem className="w-full">
              <Label htmlFor="description">Description</Label>
              <TextEditor editor={editor} />

              {errors.description && (
                <p className="text-red-500">{errors.description.message}</p>
              )}
            </FormItem>
          )}
        />

        <div className="flex gap-4 w-full">
          <FormField
            name="author"
            render={({ field }) => (
              <FormItem className="w-full">
                <Label htmlFor="author">Author</Label>
                <Input
                  placeholder="Defaults to: whatever your name is"
                  {...field}
                />

                {errors.author && (
                  <p className="text-red-500">{errors.author.message}</p>
                )}
              </FormItem>
            )}
          />

          <FormField
            name="keywords"
            render={({ field }) => (
              <FormItem className="w-full">
                <Label htmlFor="author">Keywords</Label>
                <Input
                  placeholder="eg. comedy, horror, technology, space"
                  {...field}
                />

                {errors.keywords && (
                  <p className="text-red-500">{errors.keywords.message}</p>
                )}
              </FormItem>
            )}
          />
        </div>
        <FormField
          name="episodeNumber"
          render={({ field }) => (
            <FormItem className="w-fit">
              <Label htmlFor="episodeNumber">Episode number</Label>
              <Input defaultValue={0} type="number" {...field} />
              {errors.episodeNumber && (
                <p className="text-red-500">{errors.episodeNumber.message}</p>
              )}
            </FormItem>
          )}
        />

        {form.getValues("publishDate") && (
          <div className="flex gap-4">
            <div className="bg-green-100 w-fit p-2 px-5 rounded-full text-green-700 flex items-center gap-2 text-sm">
              <FontAwesomeIcon icon={faCheckCircle} />
              <span className="font-bold">Published</span> on{" "}
              {format(
                fromUnixTime(Number(form.getValues("publishDate"))),
                "MMMM dd, yyyy hh:mm a"
              )}
            </div>

            {/* <Button
              variant="ghost"
              className="w-fit"
              type="button"
              onClick={() => setChangePublishDate(!changePublishDate)}
            >
              Change date
            </Button> */}
          </div>
        )}

        {!isEditing && (
          <>
            <FormItem>
              <Label>Publish date</Label>
              <PublishSelector setWhenToPublish={setWhenToPublish} />
            </FormItem>

            {whenToPublish === "schedule" && (
              <div className="flex items-center gap-2">
                <DatePicker date={publishDate} setDate={setPublishDate} />
                <span>at</span>
                <ScheduleHourSelector
                  time={form.getValues("scheduleHour")}
                  setTime={(v) => form.setValue("scheduleHour", v)}
                />
                <ScheduleMinuteSelector
                  time={form.getValues("scheduleMinute")}
                  setTime={(v) => form.setValue("scheduleMinute", v)}
                />
                <ScheduleMeridiemSelector
                  time={form.getValues("scheduleMeridiem")}
                  setTime={(v) => form.setValue("scheduleMeridiem", v)}
                />
              </div>
            )}
          </>
        )}
        <FormField
          name="explicitContent"
          render={({ field }) => (
            <FormItem>
              <div className="items-top flex space-x-2">
                <Checkbox id={field.name} />
                <div className="grid gap-1.5 leading-none">
                  <label
                    htmlFor={field.name}
                    className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                  >
                    Episode contains explicit content
                  </label>
                  <p className="text-sm font-light text-muted-foreground">
                    This will mark your episode as containing explicit content.
                  </p>
                </div>
              </div>
            </FormItem>
          )}
        />
        <Button
          type="submit"
          onClick={() => submit(form.getValues())}
          disabled={isSubmitting}
        >
          {isSubmitting ? <FontAwesomeIcon icon={faSpinner} spin /> : ctaText}
        </Button>
        {isEditing ? (
          <div className="flex justify-end">
            <Button
              type="button"
              variant="outlineDestructive"
              className="w-fit"
              onClick={() => deleteHandler?.(episode?.id)}
            >
              Delete episode
            </Button>
          </div>
        ) : (
          <div className="flex justify-end">
            <Button
              type="button"
              variant="secondary"
              className="w-fit"
              disabled={isSubmitting}
              onClick={() =>
                submit({
                  ...form.getValues(),
                  draft: true,
                })
              }
            >
              Save as draft
            </Button>
          </div>
        )}
      </form>
    </Form>
  );
};

const NoFileSelected = () => (
  <div className="flex items-center gap-4">
    <FontAwesomeIcon icon={faCloudArrowUp} className="text-2xl" />
    <div className="flex flex-col">
      <p className="font-medium mb-2">Drop an audio file or click to upload</p>
      <p className="text-muted-foreground text-sm font-light">
        Accepted file types - .mp3, .m4a, .aiff, .wav, .mp4 up to 1000MB in
        size.
      </p>
    </div>
  </div>
);

export default EpisodeForm;
