import { Podcast, PodcastSettings } from "@/lib/types";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import axios from "axios";

const GetPodcastByNameWithEpisodes = async (name: string) => {
  const data = await axios.get(`http://localhost:8080/api/podcast/${name}`, {
    withCredentials: true,
  });

  return data.data;
};

export const usePodcast = (name?: string) => {
  const context = useQueryClient();

  const query = useQuery<PodcastSettings>({
    queryKey: ["podcast", name],
    queryFn: () => GetPodcastByNameWithEpisodes(name ?? ""),
    enabled: !!name,
  });

  const update = useMutation(
    async ({
      podcastId,
      data,
      file,
    }: {
      podcastId: string;
      data: Podcast;
      file: File | undefined;
    }) =>
      await axios.postForm(
        `http://localhost:8080/api/podcast/${podcastId}/edit`,
        {
          ...data,
          file,
        },
        {
          withCredentials: true,
        }
      ),
    {
      onSuccess: () => {
        context.invalidateQueries(["podcast", "podcasts"]);
      },
    }
  );

  return { query, update: update.mutate };
};
