import DashNav from "@/components/dashboard/DashNav";
import { faArrowLeft } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import Link from "next/link";
import React from "react";

interface Props {
  rootPath: string;
}

const DashHeader = ({ rootPath }: Props) => {
  return (
    <header className="bg-background-alt h-full  flex flex-col">
      <header className="p-4">
        <p className="text-background-alt-foreground">Resonate</p>
      </header>
      <DashNav rootPath={rootPath} />
      <Link
        href="/"
        className="text-background-alt-foreground text-sm hover:bg-white/10 p-4 w-full"
      >
        <FontAwesomeIcon className="mr-2" icon={faArrowLeft} />
        Back to podcasts
      </Link>
    </header>
  );
};

export default DashHeader;
