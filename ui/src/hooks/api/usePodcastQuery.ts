import { PodcastSettings } from "@/lib/types";
import { useQuery } from "@tanstack/react-query";
import axios from "axios";

const GetPodcastByNameWithEpisodes = async (name: string) => {
  const data = await axios.get(`http://localhost:8080/api/podcast/${name}`, {
    withCredentials: true,
  });

  return data.data;
};

export const usePodcastQuery = (name: string) => {
  const query = useQuery<PodcastSettings>({
    queryKey: ["podcast", name],
    queryFn: () => GetPodcastByNameWithEpisodes(name),
    enabled: !!name,
  });

  return query;
};
