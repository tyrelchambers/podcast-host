import { Button } from "@/components/ui/button";
import { useUserStore } from "@/hooks/stores/userStore";
import Link from "next/link";
import React from "react";

const Header = () => {
  const user = useUserStore((state) => state.user);

  return (
    <header className="w-full flex justify-between p-4">
      Header{" "}
      <div className="flex items-center gap-4">
        <p>{user?.email}</p>
        <Link href="/episode/create">
          <Button>Create episode</Button>
        </Link>
      </div>
    </header>
  );
};

export default Header;
