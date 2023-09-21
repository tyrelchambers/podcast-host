import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import Header from "@/layouts/Header";
import { faMicrophone } from "@fortawesome/free-solid-svg-icons";
import {
  FontAwesomeIcon,
  FontAwesomeIconProps,
} from "@fortawesome/react-fontawesome";
import Link from "next/link";
import React from "react";

interface Props {
  title: string;
  description: string;
  icon?: FontAwesomeIconProps["icon"];
}

const PodcastConfigureCard = ({ title, description, icon }: Props) => {
  return (
    <Card>
      <CardHeader>
        <CardTitle className="flex gap-4 font-medium">
          {icon && <FontAwesomeIcon icon={icon} />}
          {title}
        </CardTitle>
        <CardDescription className="font-thin">{description}</CardDescription>
      </CardHeader>

      <CardFooter>
        <Link href="/podcast/create">
          <Button>Configure podcast</Button>
        </Link>
      </CardFooter>
    </Card>
  );
};

const Add = () => {
  return (
    <>
      <Header />
      <div className="container">
        <h1 className="h1">Create a new podcast</h1>

        <section className="my-10">
          <PodcastConfigureCard
            title="Configure a new podcast"
            description="Setup a new podcast from scratch."
            icon={faMicrophone}
          />
        </section>
      </div>
    </>
  );
};

export default Add;
