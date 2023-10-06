import { z } from "zod";

export interface IFile {
  lastModified: number;
  lastModifiedDate: Date;
  name: string;
  size: number;
  type: string;
}

export interface User {
  uuid: string;
  email: string;
}

export const formSchema = z.object({
  id: z.string().optional(),
  uuid: z.string(),
  file: z.unknown(),
  title: z.string(),
  description: z.string(),
  author: z.string(),
  keywords: z.string(),
  episode_number: z.number(),
  schedule_hour: z.string().optional(),
  schedule_minute: z.string().optional(),
  schedule_meridiem: z.string().optional(),
  publish_date: z.number().optional(),
  url: z.string().optional(),
  podcast_id: z.string().optional(),
  explicit_content: z.boolean().optional(),
  draft: z.boolean(),
});

export type Episode = z.infer<typeof formSchema>;

export type SelectOptions = {
  value: string;
  display: string;
  sub_categories?: SelectOptions[];
};

export interface Podcast {
  id: string;
  title: string;
  description: string;
  thumbnail: string | undefined;
  explicit_content: boolean | undefined;
  primary_category: string | undefined;
  secondary_category: string | undefined;
  author: string;
  copyright: string | undefined;
  keywords: string | undefined;
  website: string | undefined;
  language: string | undefined;
  timezone: string | undefined;
  show_owner: string | undefined;
  owner_email: string;
  display_email_in_RSS: boolean | undefined;
  user_id: string;
  episodes: Episode[];
}

export interface PodcastSettings {
  podcast: Podcast;
  latest_episode: Episode;
}

export interface MiscInfo {
  next_episode_number: number;
  rss_feed: string;
}

export const podcastSchema = z.object({
  title: z.string(),
  description: z.string(),
  thumbnail: z.string().optional(),
  explicit_content: z.boolean(),
  primary_category: z.string().optional(),
  secondary_category: z.string().optional(),
  author: z.string(),
  copyright: z.string().optional(),
  keywords: z.string().optional(),
  website: z.string().optional(),
  language: z.string().optional(),
  timezone: z.string().optional(),
  show_owner: z.string(),
  owner_email: z.string().optional(),
  display_email_in_rss_feed: z.boolean().optional(),
});
