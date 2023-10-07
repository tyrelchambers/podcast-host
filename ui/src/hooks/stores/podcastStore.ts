import { Podcast } from "@/lib/types";
import { formatTitleFromUrl } from "@/lib/utils";
import { create } from "zustand";

interface Props {
  podcasts: Podcast[];
  setPodcasts: (podcasts: Podcast[]) => void;
  activePodcast: Podcast | undefined;
  setActivePodcast: (podcast: string) => void;
}

export const usePodcastStore = create<Props>((set, get) => ({
  podcasts: [],
  setPodcasts: (podcasts: Podcast[]) => {
    set({ podcasts: podcasts });
  },
  activePodcast: undefined,
  setActivePodcast: (title: string) => {
    const newTitle = formatTitleFromUrl(title).toLowerCase();
    const podcast = get().podcasts.find((podcast) => {
      return podcast.title.toLowerCase() === newTitle;
    });

    set({ activePodcast: podcast });
  },
}));
