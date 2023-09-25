import { MiscInfo } from "@/lib/types";
import { useQuery } from "@tanstack/react-query";
import axios from "axios";

const getMiscInfo = (id: string) => {
  return axios
    .get(`http://localhost:8080/api/podcast/${id}/info`, {
      withCredentials: true,
    })
    .then((res) => res.data ?? {})
    .catch((err) => console.log(err));
};

export const useMiscInfoQuery = (podcastId: string) => {
  const query = useQuery<MiscInfo>({
    queryKey: ["miscInfo", podcastId],
    queryFn: () => getMiscInfo(podcastId),
    enabled: !!podcastId,
  });

  return query;
};
