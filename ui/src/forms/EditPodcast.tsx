import Select from "@/components/Select";
import ThumbnailPlaceholder from "@/components/ThumbnailPlaceholder";
import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import { Form, FormField, FormItem } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import {
  podcastCategoryOptions,
  spokenLanguages,
  timeZones,
} from "@/constants";
import { usePodcastStore } from "@/hooks/stores/podcastStore";
import { Podcast, podcastSchema } from "@/lib/types";
import { formatBytes, formatCategoryOptions } from "@/lib/utils";
import { faCloudArrowUp, faImage } from "@fortawesome/pro-regular-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { zodResolver } from "@hookform/resolvers/zod";
import axios from "axios";
import React, { useRef } from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";

const EditPodcast = () => {
  const fileUploadRef = useRef<HTMLInputElement>(null);
  const activePodcast = usePodcastStore((store) => store.activePodcast);

  const form = useForm({
    resolver: zodResolver(podcastSchema),
    values: activePodcast,
  });

  const fileName = fileUploadRef.current?.files?.[0]?.name;
  const fileSize = fileUploadRef.current?.files?.[0]?.size
    ? formatBytes(fileUploadRef.current?.files?.[0]?.size)
    : null;
  const fileType = fileUploadRef.current?.files?.[0]?.type;

  const watchFileValue = form.watch("thumbnail");

  console.log(activePodcast);

  const submit = async (data: Podcast) => {
    const file = fileUploadRef.current?.files?.[0];

    await axios.postForm(
      `http://localhost:8080/api/podcast/${activePodcast?.title}/edit`,
      {
        ...data,
        file,
      }
    );
  };

  return (
    <Form {...form}>
      <form
        className="flex flex-col gap-6 section-card"
        onSubmit={form.handleSubmit(submit, console.log)}
      >
        <FormField
          name="title"
          render={({ field }) => (
            <FormItem>
              <Label htmlFor={field.name} required>
                Podcast title
              </Label>
              <Input placeholder="The title of your podcast" {...field} />
            </FormItem>
          )}
        />

        <FormField
          name="description"
          render={({ field }) => (
            <FormItem>
              <Label htmlFor={field.name} required>
                Description
              </Label>
              <Textarea
                placeholder="What is your show about? Try to keep it short and sweet. Maybe a few sentences to be safe."
                {...field}
              />
            </FormItem>
          )}
        />

        <div className="flex gap-3">
          <ThumbnailPlaceholder />
          <FormField
            name="file"
            render={({ field }) => (
              <FormItem className="flex-1">
                <Label htmlFor={field.name}>
                  Thumbnail
                  <div className="w-full border-2 border-dashed border-border p-4 rounded-md h-[100px] flex items-center px-10 mt-2">
                    {!watchFileValue ? (
                      <div className="flex items-center gap-4">
                        <FontAwesomeIcon
                          icon={faCloudArrowUp}
                          className="text-2xl"
                        />
                        <div className="flex flex-col">
                          <p className="font-semibold mb-1">
                            Drop an image file or click to upload
                          </p>
                          <p className="text-muted-foreground text-sm font-light">
                            Accepted filetypes - .jpeg, .png up to 15MB in size.
                            3000 x 3000 pixels recommended.
                          </p>
                        </div>
                      </div>
                    ) : (
                      <div className="flex items-center gap-4">
                        <FontAwesomeIcon icon={faImage} className="text-3xl" />
                        <div className="flex flex-col">
                          <p className="font-semibold">{fileName}</p>
                          <p className="text-muted-foreground text-sm">
                            {fileSize} - {fileType}
                          </p>
                        </div>
                      </div>
                    )}
                    <Input
                      type="file"
                      className="hidden"
                      id={field.name}
                      {...field}
                      ref={fileUploadRef}
                    />
                  </div>
                </Label>
              </FormItem>
            )}
          />
        </div>

        <div className="flex gap-3">
          <FormField
            name="primary_category"
            render={({ field }) => (
              <FormItem className="flex-1">
                <Label htmlFor={field.name}>Primary category</Label>
                <Select
                  placeholder="Primary category"
                  options={formatCategoryOptions(podcastCategoryOptions)}
                  onValueChange={field.onChange}
                  {...field}
                />
              </FormItem>
            )}
          />

          <FormField
            name="secondary_category"
            render={({ field }) => (
              <FormItem className="flex-1">
                <Label htmlFor={field.name}>Secondary category</Label>
                <Select
                  placeholder="Secondary category"
                  options={formatCategoryOptions(podcastCategoryOptions)}
                  onValueChange={field.onChange}
                  {...field}
                />
              </FormItem>
            )}
          />
        </div>

        <FormField
          name="author"
          render={({ field }) => (
            <FormItem>
              <Label htmlFor={field.name} required>
                Author
              </Label>
              <Input placeholder="Your name" {...field} />
            </FormItem>
          )}
        />

        <FormField
          name="copyright"
          render={({ field }) => (
            <FormItem>
              <Label htmlFor={field.name}>Copyright</Label>
              <Input placeholder="e.g. © 2022 Your Podcast" {...field} />
            </FormItem>
          )}
        />

        <FormField
          name="keywords"
          render={({ field }) => (
            <FormItem>
              <Label htmlFor={field.name}>Keywords</Label>
              <Input placeholder="e.g. podcast, audio" {...field} />
            </FormItem>
          )}
        />

        <FormField
          name="website"
          render={({ field }) => (
            <FormItem>
              <Label htmlFor={field.name}>Website</Label>
              <Input placeholder="e.g. http://yourpodcast.com" {...field} />
            </FormItem>
          )}
        />

        <div className="grid grid-cols-2 gap-3">
          <FormField
            name="language"
            render={({ field }) => (
              <FormItem>
                <Label htmlFor={field.name}>Language</Label>
                <Select
                  options={spokenLanguages}
                  placeholder="Your language"
                  onValueChange={field.onChange}
                  {...field}
                />
              </FormItem>
            )}
          />

          <FormField
            name="timezone"
            render={({ field }) => (
              <FormItem>
                <Label htmlFor={field.name}>Timezone</Label>
                <Select
                  options={timeZones}
                  placeholder="Your timezone"
                  onValueChange={field.onChange}
                  {...field}
                />
              </FormItem>
            )}
          />

          <FormField
            name="show_owner"
            render={({ field }) => (
              <FormItem>
                <Label htmlFor={field.name}>Show owner</Label>
                <Input placeholder="Name of show owner" {...field} />
              </FormItem>
            )}
          />

          <FormField
            name="owner_email"
            render={({ field }) => (
              <FormItem>
                <Label htmlFor={field.name} required>
                  Owner email
                </Label>
                <Input placeholder="email@example.com" {...field} />
              </FormItem>
            )}
          />
        </div>

        <FormField
          name="explicit_content"
          render={({ field }) => (
            <FormItem>
              <div className="items-top flex space-x-2">
                <Checkbox id={field.name} />
                <div className="grid gap-1.5 leading-none">
                  <label
                    htmlFor={field.name}
                    className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                  >
                    Show contains explicit content
                  </label>
                  <p className="text-sm font-light text-muted-foreground">
                    This will mark your podcast as explicit in podcast
                    distribution platforms.
                  </p>
                </div>
              </div>
            </FormItem>
          )}
        />

        <FormField
          name="display_email_in_rss_feed"
          render={({ field }) => (
            <FormItem>
              <div className="items-top flex space-x-2">
                <Checkbox id={field.name} />
                <div className="grid gap-1.5 leading-none">
                  <label
                    htmlFor={field.name}
                    className="text-sm  font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                  >
                    Display this email in your RSS Feed
                  </label>
                  <p className="text-sm font-light text-muted-foreground">
                    To reduce potential spam sent to your email we hide it from
                    your RSS Feed by default, but you can enable it for
                    verification or transfer purposes.
                  </p>
                </div>
              </div>
            </FormItem>
          )}
        />

        <Button type="submit">Save changes</Button>
      </form>
    </Form>
  );
};

export default EditPodcast;
