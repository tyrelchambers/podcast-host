import { Podcast } from "@/lib/types";
import React from "react";
import ThumbnailPlaceholder from "./ThumbnailPlaceholder";
import Image from "next/image";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faListMusic } from "@fortawesome/pro-regular-svg-icons";
interface Props {
  podcasts: Podcast[];
}

const PodcastList = ({ podcasts }: Props) => {
  if (!podcasts?.length) {
    return null;
  }

  return (
    <div className="flex flex-col">
      {podcasts.map((podcast) => (
        <div key={podcast.id} className="flex gap-6">
          {podcast.thumbnail ? (
            <Image src={podcast.thumbnail} alt="" />
          ) : (
            <ThumbnailPlaceholder />
          )}

          <div className="flex flex-col gap-3">
            <p className="font-medium text-2xl">{podcast.title}</p>
            <p className="font-light  text-muted-foreground">
              {podcast.author}
            </p>

            <footer className="flex gap-3">
              <p>
                <FontAwesomeIcon icon={faListMusic} />
                {podcast.episodes?.length}
              </p>
            </footer>
          </div>
        </div>
      ))}
    </div>
  );
};

export default PodcastList;
