import { type AppType } from "next/app";

import { config } from "@fortawesome/fontawesome-svg-core";
import "@fortawesome/fontawesome-svg-core/styles.css";
config.autoAddCss = false;

import { Poppins } from "next/font/google";

import "../styles/globals.css";
import "@mantine/core/styles.css";
import "@mantine/tiptap/styles.css";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { useUserQuery } from "@/hooks/api/useUserQuery";
import { useEffect } from "react";
import { useUserStore } from "@/hooks/stores/userStore";
import { useRouter } from "next/router";
import { usePodcastStore } from "@/hooks/stores/podcastStore";
import { usePodcastsQuery } from "@/hooks/api/usePodcastsQuery";
import { MantineProvider } from "@mantine/core";

const queryClient = new QueryClient();

const font = Poppins({
  weight: ["300", "500", "700"],
  subsets: ["latin"],
});

const InitialData = ({ children }: { children: React.ReactNode }) => {
  const user = useUserQuery();
  const podcasts = usePodcastsQuery(user.data?.uuid ?? "");

  const userStore = useUserStore();
  const router = useRouter();
  const podcastStore = usePodcastStore();

  useEffect(() => {
    if (user.data) {
      userStore.setUser(user.data);
    }
  }, [user.data]);

  useEffect(() => {
    if (podcasts.data) {
      podcastStore.setPodcasts(podcasts.data);
    }
  }, [podcasts.data]);

  useEffect(() => {
    if (router.query.name && router.isReady && podcastStore.podcasts) {
      podcastStore.setActivePodcast(router.query.name as string);
    }
  }, [router.query.name, router.isReady, podcastStore.podcasts]);

  return <>{children}</>;
};

const MyApp: AppType = ({ Component, pageProps: { ...pageProps } }) => {
  return (
    <QueryClientProvider client={queryClient}>
      <InitialData>
        <MantineProvider>
          <main className={font.className}>
            <Component {...pageProps} />
          </main>
        </MantineProvider>
      </InitialData>
    </QueryClientProvider>
  );
};

export default MyApp;
