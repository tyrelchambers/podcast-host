import { useQuery } from "@tanstack/react-query";
import axios from "axios";

export const getPodcasts = async (id: string | undefined) => {
  const data = await axios
    .get(`http://localhost:8080/api/user/${id}/podcasts`, {
      withCredentials: true,
    })
    .then((res) => res.data ?? [])
    .catch((err) => {
      console.log(err);
    });

  return data;
};

export const usePodcastsQuery = (id: string | undefined) => {
  const query = useQuery({
    queryKey: ["podcasts"],
    queryFn: () => getPodcasts(id),
    enabled: !!id,
  });

  return query;
};
