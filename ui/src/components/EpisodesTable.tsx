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

interface Props {
  episodes: Episode[];
}

const EpisodesTable = ({ episodes }: Props) => {
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
          episodes.map((episode) => (
            <TableRow key={episode.id}>
              <TableCell>{episode.episodeNumber}</TableCell>
              <TableCell></TableCell>
              <TableCell>
                <Link href={`/episode/${episode.id}/edit`}>
                  {episode.title}
                </Link>
              </TableCell>
              <TableCell>Published</TableCell>
              <TableCell className="text-right">
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
