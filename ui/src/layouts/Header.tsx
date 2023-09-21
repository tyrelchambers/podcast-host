import Avatar from "@/components/Avatar";
import { useUserStore } from "@/hooks/stores/userStore";
import React from "react";

const Header = () => {
  const user = useUserStore((state) => state.user);

  return (
    <header className="w-full flex justify-between p-4">
      <p>header</p>{" "}
      {user && (
        <div className="flex items-center gap-2 bg-card p-2 rounded-full">
          <Avatar src={user?.avatar} fallback="EM" />
          <p className="card-foreground">{user?.email}</p>
        </div>
      )}
    </header>
  );
};

export default Header;
