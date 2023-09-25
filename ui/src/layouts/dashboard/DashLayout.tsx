import { useRouter } from "next/router";
import React, { useEffect } from "react";

interface Props {
  leftCol: React.ReactNode;
  children: React.ReactNode | React.ReactNode[];
  rightCol: React.ReactNode;
}
const DashLayout = ({ leftCol, children, rightCol }: Props) => {
  const router = useRouter();
  const nameParam = router.query.name;

  return !nameParam ? null : (
    <div className="dash-layout-grid ">
      <div className=" w-[250px] h-screen sticky top-0 bottom-0">{leftCol}</div>
      <section className="flex-1 h-full p-8 overflow-y-auto">
        {children}
      </section>
      <div className=" w-[250px] h-screen sticky top-0 bottom-0">
        {rightCol}
      </div>
    </div>
  );
};

export default DashLayout;
