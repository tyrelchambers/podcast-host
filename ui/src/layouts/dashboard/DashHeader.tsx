import DashNav from "@/components/dashboard/DashNav";
import React from "react";

interface Props {
  rootPath: string;
}

const DashHeader = ({ rootPath }: Props) => {
  return (
    <header className="bg-card h-full p-4">
      <p>Resonate</p>
      <DashNav rootPath={rootPath} />
    </header>
  );
};

export default DashHeader;
