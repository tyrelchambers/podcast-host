import { IconProp } from "@fortawesome/fontawesome-svg-core";
import {
  faChartMixed,
  faCog,
  faGrid2,
  faListMusic,
  faTruckFast,
} from "@fortawesome/pro-duotone-svg-icons";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import Link from "next/link";
import React from "react";

interface Route {
  path: string;
  icon: IconProp;
  label: string;
}

const routes = (rootPath: string): Route[] => {
  return [
    {
      path: `${rootPath}/`,
      icon: faGrid2,
      label: "Overview",
    },
    {
      path: `${rootPath}/episodes`,
      icon: faListMusic,
      label: "Episodes",
    },
    {
      path: `${rootPath}/distribution`,
      icon: faTruckFast,
      label: "Distribution",
    },
    {
      path: `${rootPath}/analytics`,
      icon: faChartMixed,
      label: "Analytics",
    },
    {
      path: `${rootPath}/settings`,
      icon: faCog,
      label: "Settings",
    },
  ];
};

interface Props {
  rootPath: string;
}
const DashNav = ({ rootPath }: Props) => {
  return (
    <nav className="my-10 flex-1">
      <ul className="flex flex-col gap-4">
        {routes(rootPath).map((route) => (
          <li key={route.path} className="hover:text-accent transition-all">
            <Link href={route.path} className="flex items-center gap-3">
              <FontAwesomeIcon icon={route.icon} />
              <span className="font-light ">{route.label}</span>
            </Link>
          </li>
        ))}
      </ul>
    </nav>
  );
};

export default DashNav;
