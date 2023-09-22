import { useQuery } from "@tanstack/react-query";
import axios from "axios";

const getPodcastByName = async (name: string) => {
  const data = await axios.get(`http://localhost:8080/api/podcast/${name}`, {
    withCredentials: true,
  });

  return data.data;
};

export const usePodcastQuery = (name: string) => {
  const query = useQuery({
    queryKey: ["podcast", name],
    queryFn: () => getPodcastByName(name),
    enabled: !!name,
  });
};
