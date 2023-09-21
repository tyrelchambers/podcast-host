import { User } from "@/lib/types";
import { create } from "zustand";

interface IUser {
  user: User | undefined;
  setUser: (user: User | undefined) => void;
}

export const useUserStore = create<IUser>((set) => ({
  user: undefined,
  setUser: (user) => set({ user }),
}));
