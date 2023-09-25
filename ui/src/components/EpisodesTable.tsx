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
}

const EpisodesTable = ({ episodes }: Props) => {
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
            .sort((a, b) => Number(b.episodeNumber) - Number(a.episodeNumber))
            .map((episode) => (
              <TableRow key={episode.id}>
                <TableCell className="text-muted-foreground font-light">
                  {episode.episodeNumber}
                </TableCell>
                <TableCell></TableCell>
                <TableCell className="text-blue-500 underline font-medium">
                  <Link href={`/episode/${episode.id}/edit`}>
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
                  {format(
                    fromUnixTime(Number(episode.publishDate) ?? 0),
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
