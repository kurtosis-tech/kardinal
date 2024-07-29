"use client";

import { ReactNode, useEffect, useState } from "react";

import analytics from "@/lib/analytics";

import VideoStepperV2 from "./VideoStepperV2";

interface Props {
  children?: ReactNode;
  heading?: ReactNode;
  variation?: "default" | "v1" | "v2" | "v5";
}

export interface VideoStep {
  title: string;
  content: string;
  duration: number;
}

const VideoStepper = ({ children, heading }: Props) => {
  const [activeStep, setActiveStep] = useState(0);
  useEffect(() => {
    const handleFullscreenChange = () => {
      const isFullscreen = document.fullscreenElement !== null;
      if (isFullscreen) {
        analytics.track("BUTTON_CLICK", { analyticsId: "video_fullscreen" });
      }
    };

    document.addEventListener("fullscreenchange", handleFullscreenChange);

    // Clean up by removing the event listener when the component unmounts
    return () => {
      document.removeEventListener("fullscreenchange", handleFullscreenChange);
    };
  }, []);

  const onStepChange = (step: number) => {
    setActiveStep(step);
    analytics.track("BUTTON_CLICK", {
      analyticsId: `video_section_step_${step + 1}`,
    });
  };

  const onPause = () => {
    analytics.track("BUTTON_CLICK", {
      analyticsId: `video_pause`,
    });
  };
  return (
    <VideoStepperV2
      heading={heading}
      activeStep={activeStep}
      onStepChange={onStepChange}
      onPause={onPause}
    >
      {children}
    </VideoStepperV2>
  );
};

export default VideoStepper;
