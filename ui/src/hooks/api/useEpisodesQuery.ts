import { Episode } from "@/lib/types";
import { useQuery } from "@tanstack/react-query";
import axios from "axios";

export const getAllEpisodes = async (id: string | undefined) => {
  const data = await axios
    .get(`http://localhost:8080/api/user/${id}/episodes`)
    .then((res) => res.data)
    .catch((err) => {
      console.log(err);
    });

  return data;
};

export const useEpisodesQuery = (userId: string | undefined) => {
  const query = useQuery<Episode[]>({
    queryKey: ["episodes"],
    queryFn: () => getAllEpisodes(userId),
    enabled: !!userId,
  });

  return query;
};
