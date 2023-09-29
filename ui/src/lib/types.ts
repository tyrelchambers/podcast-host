import { z } from "zod";

export interface IFile {
  lastModified: number;
  lastModifiedDate: Date;
  name: string;
  size: number;
  type: string;
}

export interface User {
  id: string;
  email: string;
}

export const formSchema = z.object({
  id: z.string().optional(),
  file: z.unknown(),
  title: z.string(),
  description: z.string(),
  author: z.string(),
  keywords: z.string(),
  episode_number: z.number(),
  scheduleHour: z.string().optional(),
  scheduleMinute: z.string().optional(),
  scheduleMeridiem: z.string().optional(),
  publish_date: z.number().optional(),
  url: z.string().optional(),
  podcastId: z.string().optional(),
  explicitContent: z.boolean().optional(),
  draft: z.boolean(),
});

export type Episode = z.infer<typeof formSchema>;

export type SelectOptions = {
  value: string;
  display: string;
  subCategories?: SelectOptions[];
};

export interface Podcast {
  id: string;
  title: string;
  description: string;
  thumbnail: string | undefined;
  explicitContent: boolean | undefined;
  primaryCategory: string | undefined;
  secondaryCategory: string | undefined;
  author: string;
  copyright: string | undefined;
  keywords: string | undefined;
  website: string | undefined;
  language: string | undefined;
  timezone: string | undefined;
  showOwner: string | undefined;
  ownerEmail: string;
  displayEmailInRSS: boolean | undefined;
  userID: string;
  episodes: Episode[];
}

export interface PodcastSettings {
  podcast: Podcast;
  latestEpisode: Episode;
}

export interface MiscInfo {
  nextEpisodeNumber: number;
  rssFeed: string;
}
