import { User } from "@/lib/types";
import { useQuery } from "@tanstack/react-query";
import axios, { AxiosError, AxiosResponse } from "axios";

const getCurrentUser = async () => {
  const data = await axios
    .get("http://localhost:8080/api/user/me", {
      withCredentials: true,
    })
    .then((res) => res.data)
    .catch((err) => {
      console.log(err);
    });

  return data;
};

export const useUserQuery = () => {
  const user = useQuery<User>({
    queryKey: ["currentUser"],
    queryFn: getCurrentUser,
  });

  return {
    ...user,
  };
};
