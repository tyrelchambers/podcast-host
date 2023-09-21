import { type AppType } from "next/app";

import { config } from "@fortawesome/fontawesome-svg-core";
import "@fortawesome/fontawesome-svg-core/styles.css";
config.autoAddCss = false;

import { Poppins } from "next/font/google";

import "../styles/globals.css";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { getCurrentUser } from "@/hooks/api/useUserQuery";
import { useEffect } from "react";
import { useUserStore } from "@/hooks/stores/userStore";

const queryClient = new QueryClient();

const font = Poppins({
  weight: ["300", "500", "700"],
  subsets: ["latin"],
});

const MyApp: AppType = ({ Component, pageProps: { ...pageProps } }) => {
  const userStore = useUserStore();

  useEffect(() => {
    const fn = async () => {
      const currentUser = await getCurrentUser();

      if (currentUser) {
        userStore.setUser(currentUser);
      }
    };

    fn();
  }, []);

  return (
    <QueryClientProvider client={queryClient}>
      <main className={font.className}>
        <Component {...pageProps} />
      </main>
    </QueryClientProvider>
  );
};

export default MyApp;
