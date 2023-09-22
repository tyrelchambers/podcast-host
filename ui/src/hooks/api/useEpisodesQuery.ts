import { Episode } from "@/lib/types";
import { useQuery } from "@tanstack/react-query";
import axios from "axios";

export const getAllEpisodes = async (name: string | undefined) => {
  const data = await axios
    .get(`http://localhost:8080/api/podcast/${name}/episodes`)
    .then((res) => res.data)
    .catch((err) => {
      console.log(err);
    });

  return data;
};

export const useEpisodesQuery = (name: string | undefined) => {
  const query = useQuery<Episode[]>({
    queryKey: ["episodes"],
    queryFn: () => getAllEpisodes(name),
    enabled: !!name,
  });

  return query;
};
