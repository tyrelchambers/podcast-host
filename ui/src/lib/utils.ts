import { type ClassValue, clsx } from "clsx";
import { twMerge } from "tailwind-merge";
import { PodcastCategoryOption } from "./types";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export function formatBytes(bytes: number, decimals = 2) {
  if (!+bytes) return "0 Bytes";

  const k = 1024;
  const dm = decimals < 0 ? 0 : decimals;
  const sizes = [
    "Bytes",
    "KiB",
    "MiB",
    "GiB",
    "TiB",
    "PiB",
    "EiB",
    "ZiB",
    "YiB",
  ];

  const i = Math.floor(Math.log(bytes) / Math.log(k));

  return `${parseFloat((bytes / Math.pow(k, i)).toFixed(dm))} ${sizes[i]}`;
}

export const formatCategoryOptions = (options: PodcastCategoryOption[]) => {
  const arr = [];

  for (let index = 0; index < options.length; index++) {
    const element = options[index];

    arr.push(element);

    if (element.subCategories) {
      for (let j = 0; j < element.subCategories.length; j++) {
        const subEl = element.subCategories[j];
        arr.push({ ...subEl, value: subEl.value.replace("|", " :: ") });
      }
    }
  }

  return arr;
};
