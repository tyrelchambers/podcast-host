import { Button } from "@/components/ui/button";
import Link from "next/link";
import React from "react";

const Header = () => {
  return (
    <header className="w-full flex justify-between p-4">
      Header{" "}
      <Link href="/episode/create">
        <Button>Create episode</Button>
      </Link>
    </header>
  );
};

export default Header;
