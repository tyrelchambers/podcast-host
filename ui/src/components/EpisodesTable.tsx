import React from "react";
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "./ui/table";
import { Episode } from "@/lib/types";
import Link from "next/link";
import { format, fromUnixTime } from "date-fns";
import { Badge } from "./ui/badge";

interface Props {
  episodes: Episode[] | undefined;
  podcastName: string;
}

const EpisodesTable = ({ episodes, podcastName }: Props) => {
  if (!episodes?.length) {
    return null;
  }
  return (
    <Table>
      <TableHeader>
        <TableRow>
          <TableHead className="w-[100px]">#</TableHead>
          <TableHead></TableHead>
          <TableHead>Episode title</TableHead>
          <TableHead>Status</TableHead>
          <TableHead className="text-right">Publish date</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {episodes.length > 0 &&
          episodes
            .sort((a, b) => b.episode_number - a.episode_number)
            .map((episode) => (
              <TableRow key={episode.id}>
                <TableCell className="text-muted-foreground font-light">
                  {episode.episode_number}
                </TableCell>
                <TableCell></TableCell>
                <TableCell className="text-blue-500 underline font-medium">
                  <Link
                    href={`/podcast/${podcastName}/episode/${episode.uuid}/edit`}
                  >
                    {episode.title}
                  </Link>
                </TableCell>
                <TableCell>
                  {!episode.draft ? (
                    <Badge>Published</Badge>
                  ) : (
                    <Badge variant="secondary">Draft</Badge>
                  )}
                </TableCell>
                <TableCell className="text-right text-muted-foreground font-light">
                  {episode.publish_date &&
                    format(
                      fromUnixTime(episode.publish_date),
                      "MMM dd, yyyy hh:mm a"
                    )}
                </TableCell>
              </TableRow>
            ))}
      </TableBody>
    </Table>
  );
};

export default EpisodesTable;
