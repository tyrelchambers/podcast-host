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
  episodeNumber: z.string(),
  scheduleHour: z.string().optional(),
  scheduleMinute: z.string().optional(),
  scheduleMeridiem: z.string().optional(),
  publishDate: z.string().optional(),
  url: z.string().optional(),
});

export type Episode = z.infer<typeof formSchema>;

export type SelectOptions = {
  value: string;
  display: string;
  subCategories?: SelectOptions[];
};
