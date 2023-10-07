import { type ClassValue, clsx } from "clsx";
import { twMerge } from "tailwind-merge";
import { SelectOptions } from "./types";

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

export const formatCategoryOptions = (options: SelectOptions[]) => {
  const arr = [];

  for (let index = 0; index < options.length; index++) {
    const { value, display, sub_categories } = options[index];

    arr.push({ value, display });

    if (sub_categories) {
      for (let j = 0; j < sub_categories.length; j++) {
        const subEl = sub_categories[j];
        arr.push({ ...subEl, value: subEl.value.replace("|", " :: ") });
      }
    }
  }

  return arr;
};

export const formatUrlFromTitle = (title: string) => {
  return title.toLowerCase().replace(/ /g, "-");
};

export const formatTitleFromUrl = (title: string) => {
  return title.replace(/-/g, " ");
};
