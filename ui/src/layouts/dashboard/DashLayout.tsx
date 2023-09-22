import React from "react";

interface Props {
  leftCol: React.ReactNode;
  children: React.ReactNode | React.ReactNode[];
  rightCol: React.ReactNode;
}
const DashLayout = ({ leftCol, children, rightCol }: Props) => {
  return (
    <div className="dash-layout-grid h-screen">
      <div className=" w-[250px] h-full">{leftCol}</div>
      <section className="flex-1 h-full p-8">{children}</section>
      <div className=" w-[250px] h-full">{rightCol}</div>
    </div>
  );
};

export default DashLayout;
