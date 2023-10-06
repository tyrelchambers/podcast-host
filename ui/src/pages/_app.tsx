import { type AppType } from "next/app";

import { config } from "@fortawesome/fontawesome-svg-core";
import "@fortawesome/fontawesome-svg-core/styles.css";
config.autoAddCss = false;

import { Poppins } from "next/font/google";

import "../styles/globals.css";
import "@mantine/core/styles.css";
import "@mantine/tiptap/styles.css";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { getCurrentUser } from "@/hooks/api/useUserQuery";
import { useEffect } from "react";
import { useUserStore } from "@/hooks/stores/userStore";
import { useRouter } from "next/router";
import { usePodcastStore } from "@/hooks/stores/podcastStore";
import { getPodcasts } from "@/hooks/api/usePodcastsQuery";
import { MantineProvider } from "@mantine/core";

const queryClient = new QueryClient();

const font = Poppins({
  weight: ["300", "500", "700"],
  subsets: ["latin"],
});

const MyApp: AppType = ({ Component, pageProps: { ...pageProps } }) => {
  const userStore = useUserStore();
  const router = useRouter();
  const podcastStore = usePodcastStore();

  useEffect(() => {
    const fn = async () => {
      const currentUser = await getCurrentUser();
      const podcasts = await getPodcasts(currentUser?.uuid);

      if (currentUser) {
        userStore.setUser(currentUser);
      }

      if (podcasts) {
        podcastStore.setPodcasts(podcasts);
      }

      podcastStore.setActivePodcast(router.query.name as string);
    };

    fn();
  }, [router.query.name]);

  return (
    <QueryClientProvider client={queryClient}>
      <MantineProvider>
        <main className={font.className}>
          <Component {...pageProps} />
        </main>
      </MantineProvider>
    </QueryClientProvider>
  );
};

export default MyApp;
