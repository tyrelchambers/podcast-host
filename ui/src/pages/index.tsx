import { Button } from "@/components/ui/button";
import { useUserStore } from "@/hooks/stores/userStore";
import Header from "@/layouts/Header";
import Link from "next/link";

export default function Home() {
  const user = useUserStore((state) => state.user);

  return (
    <main className="w-full">
      <Header />

      <section className="max-w-screen-2xl mx-auto my-20 p-8 rounded-xl bg-card shadow-sm">
        <header className="flex items-center justify-between gap-3">
          <h1 className="h1">Shows</h1>
          <Link href="/podcast/add">
            <Button>Create show</Button>
          </Link>
        </header>
      </section>
    </main>
  );
}
