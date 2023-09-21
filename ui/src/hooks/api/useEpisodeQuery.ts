import { Episode } from "@/lib/types";
import { useQuery } from "@tanstack/react-query";
import axios from "axios";

export const getEpisodeById = async (id: string): Promise<Episode> => {
  return await axios
    .get("http://localhost:8080/api/episode/" + id, {
      withCredentials: true,
    })
    .then((res) => res.data ?? {})
    .catch((err) => console.log(err));
};

export const useEpisodeQuery = (id: string) => {
  const query = useQuery<Episode>({
    queryKey: ["episode"],
    queryFn: () => getEpisodeById(id),
    enabled: !!id,
  });

  return query;
};
