import { Podcast } from "@/lib/types";
import { formatTitleFromUrl } from "@/lib/utils";
import { create } from "zustand";

interface Props {
  podcasts: Map<string, Podcast>;
  setPodcasts: (podcasts: Podcast[]) => void;
  findPodcast: (title: string | undefined) => Podcast | undefined;
  activePodcast: Podcast | undefined;
  setActivePodcast: (podcast: string) => void;
}

export const usePodcastStore = create<Props>((set, get) => ({
  podcasts: new Map(),
  setPodcasts: (podcasts: Podcast[]) => {
    const obj = new Map();

    for (const podcast of podcasts) {
      obj.set(podcast.title, podcast);
    }

    set({ podcasts: obj });
  },
  findPodcast: (title: string | undefined) => {
    return title ? get().podcasts.get(formatTitleFromUrl(title)) : undefined;
  },
  activePodcast: undefined,
  setActivePodcast: (title: string) => {
    const findPodcast = get().findPodcast(title);
    set({ activePodcast: findPodcast });
  },
}));
