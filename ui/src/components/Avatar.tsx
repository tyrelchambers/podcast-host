import React from "react";
import {
  Avatar as CNAvatar,
  AvatarFallback,
  AvatarImage,
} from "@/components/ui/avatar";

const Avatar = ({ src, fallback }: { src: string; fallback: string }) => {
  return (
    <CNAvatar>
      <AvatarImage src={src} alt="" />
      <AvatarFallback>{fallback}</AvatarFallback>
    </CNAvatar>
  );
};

export default Avatar;
