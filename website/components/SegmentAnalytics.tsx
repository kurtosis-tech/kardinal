"use client";

import { usePathname } from "next/navigation";
import { useEffect } from "react";

import analytics from "@/lib/analytics";

export default function Analytics() {
  const pathname = usePathname();

  useEffect(() => {
    requestAnimationFrame(() => analytics.page());
  }, [pathname]);

  // track when the user scrolls 1vh
  useEffect(() => {
    const scrollThreshold = window.innerHeight;
    const handleScroll = () => {
      const scrollPosition = window.pageYOffset;
      if (scrollPosition > scrollThreshold) {
        analytics.track("SCROLLED_1VH");

        // Remove the event listener once the event has been fired once
        window.removeEventListener("scroll", handleScroll);
      }
    };

    // Add the event listener when the component mounts
    window.addEventListener("scroll", handleScroll);

    // Clean up the event listener when the component unmounts
    return () => {
      window.removeEventListener("scroll", handleScroll);
    };
  }, []);

  // track user time on page
  useEffect(() => {
    let timeOnPage = 0;
    let interval: number | null = null;
    analytics.track("TIME_ON_PAGE", {
      duration_seconds: timeOnPage,
      page_path: pathname,
    });

    const startTracking = () => {
      if (interval !== null) {
        clearInterval(interval);
      }
      // @ts-ignore
      interval = setInterval(() => {
        timeOnPage += 5;
        analytics.track("TIME_ON_PAGE", {
          duration_seconds: timeOnPage,
          page_path: pathname,
        });
      }, 5000);
    };

    const stopTracking = () => {
      if (interval !== null) {
        clearInterval(interval);
        interval = null;
      }
    };

    // Start tracking when the page is loaded
    startTracking();

    // Listen for visibility changes, stop tracking when tab is unfocused
    const handleVisibilityChange = () => {
      if (document.hidden) {
        stopTracking();
      } else {
        startTracking();
      }
    };
    document.addEventListener("visibilitychange", handleVisibilityChange);

    // Clean up on unmount
    return () => {
      stopTracking();
      document.removeEventListener("visibilitychange", handleVisibilityChange);
    };
  }, [pathname]);

  return null;
}
