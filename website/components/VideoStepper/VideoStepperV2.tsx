"use client";

import Image from "next/image";
import { ReactNode, SyntheticEvent, useRef, useState } from "react";
import styled from "styled-components";

import DottedLine from "@/components/DottedLine";
import Heading from "@/components/Heading";
import Section from "@/components/Section";
import Text from "@/components/Text";
import { mobile, tablet } from "@/constants/breakpoints";
// assets
import continuityImg from "@/public/illustrations/continuity-3.svg";
import continuityImgMobile from "@/public/illustrations/continuity-3-mobile.svg";
import video from "@/public/videos/animation.mp4";

import type { VideoStep } from ".";

const SECOND = 1000;
const TOTAL_VIDEO_DURATION = 45 * SECOND;

const steps: VideoStep[] = [
  {
    title: "Staging at a glance",
    content:
      "Imagine you're on the voting-app team. \
            Your shared staging environment looks like this.",
    duration: 8 * SECOND,
  },
  {
    title: "Spin up a dev flow",
    content:
      "Create an isolated, logical dev environment with just one command.",
    duration: 12 * SECOND,
  },
  {
    title: "Spin up a second dev flow",
    content:
      "Create another isolated dev environment, spinning up only necessary services.",
    duration: 10 * SECOND,
  },
  {
    title: "Develop with confidence",
    content:
      "Test new features on staging in an isolated traffic flow. Access to application state is isolated by a stateful service sidecar.",
    duration: 15 * SECOND,
  },
];

// throw an error if the total video duration does not match the sum of the
// video durations this check is just to ensure future updates to this
// component are consistent and result in a working animation
if (
  steps.reduce((acc, step) => acc + step.duration, 0) !== TOTAL_VIDEO_DURATION
) {
  throw new Error(
    "VideoStepperV2: Sum of step durations do not match the total video duration",
  );
}

interface Props {
  children: ReactNode;
  heading: ReactNode;
  // eslint-disable-next-line no-unused-vars
  onStepChange: (step: number) => void;
  onPause: () => void;
  activeStep: number;
}

const VideoStepper = ({
  children,
  heading,
  activeStep,
  onStepChange,
  onPause,
}: Props) => {
  const [progress, setProgress] = useState(0);
  const [shouldAnimateProgress, setShouldAnimateProgress] = useState(true);
  const videoRef = useRef<HTMLVideoElement>(null);

  const onTimeUpdate = (e: SyntheticEvent<HTMLVideoElement, Event>) => {
    const { currentTime, duration } = e.currentTarget;
    const percentComplete = Math.floor(100 * (currentTime / duration));
    // +8% lag adjustment as the css animation delays the progress indicator
    // slightly compared to the true video progress
    const lagAdjustment = 8;
    setProgress(percentComplete + lagAdjustment);

    // video segments are not evenly distributed, these timings are to make the
    // progress indicator visually align with the video transitions
    const part0end = Math.floor(
      (steps[0].duration / TOTAL_VIDEO_DURATION) * 100,
    );
    if (activeStep !== 0 && percentComplete < part0end) {
      onStepChange(0);
      setShouldAnimateProgress(false);

      setTimeout(() => {
        setShouldAnimateProgress(true);
      }, 100);
      return;
    }
    const part1end =
      part0end + (steps[1].duration / TOTAL_VIDEO_DURATION) * 100;
    if (
      activeStep !== 1 &&
      percentComplete > part0end &&
      percentComplete < part1end
    ) {
      onStepChange(1);
      return;
    }
    const part2end =
      part1end + (steps[2].duration / TOTAL_VIDEO_DURATION) * 100;
    if (
      activeStep !== 2 &&
      percentComplete > part1end &&
      percentComplete < part2end
    ) {
      onStepChange(2);
      return;
    }
    if (activeStep !== 3 && percentComplete > part2end) {
      onStepChange(3);
    }
  };

  const handleStepChange = (index: number) => {
    if (videoRef.current) {
      const jumpTo = videoRef.current.duration * (index / steps.length);
      videoRef.current.currentTime = jumpTo;
    }
    onStepChange(index);
  };

  return (
    <Section flexDirection="row">
      <S.VideoStepper>
        <S.ContinuityWrapper>
          <S.ContinuityImage
            src={continuityImg}
            width={83}
            height={382}
            alt="Diagram lines"
          />
          <S.ContinuityImageMobile
            src={continuityImgMobile}
            width={83}
            height={225}
            alt="Diagram lines"
          />
          <DottedLine
            $left
            $width={105}
            $height={1097}
            $offsetTop={-72}
            $offsetLeft={28}
          />
        </S.ContinuityWrapper>
        <S.ContentWrapper>
          <S.TextContent>
            <S.HowItWorksText>How it works</S.HowItWorksText>
            <Heading.H2>{heading}</Heading.H2>
            <Text.Base>{children}</Text.Base>
          </S.TextContent>
          <S.VideoWrapper>
            <div></div>
            <S.Video
              ref={videoRef}
              playsInline
              onPause={onPause}
              autoPlay
              muted
              loop
              controls
              onTimeUpdate={onTimeUpdate}
              preload="auto"
            >
              <source src={video} type="video/mp4" />
            </S.Video>
          </S.VideoWrapper>
          <S.Steps>
            <DottedLine
              $left
              $bottom
              $width={105}
              $height={48}
              $offsetTop={0}
              $offsetLeft={-77}
              style={{ position: "absolute" }}
            />

            <S.LineWrapper>
              <DottedLine $top $width={"100%"} />
              <S.ProgressOverflow>
                <S.ProgressIndicator
                  $progress={progress}
                  $animate={shouldAnimateProgress}
                />
              </S.ProgressOverflow>
            </S.LineWrapper>
            {steps.map((video, index) => (
              <S.Step
                key={video.title}
                $isActive={activeStep === index}
                onClick={() => {
                  handleStepChange(index);
                }}
                role="button"
              >
                <S.StepNumber>{index + 1}</S.StepNumber>
                <S.StepContentWrapper>
                  <S.StepHeading>{video.title}</S.StepHeading>
                  <S.StepContent>{video.content}</S.StepContent>
                </S.StepContentWrapper>
              </S.Step>
            ))}
          </S.Steps>
        </S.ContentWrapper>
      </S.VideoStepper>
    </Section>
  );
};

export namespace S {
  export const VideoStepper = styled.div`
    display: grid;
    grid-template-columns: 105px 1fr;
    width: 100%;
    @media ${tablet} {
      padding: 0;
      grid-template-columns: 1fr;
      ${DottedLine} {
        display: none;
      }
    }
  `;

  export const HowItWorksText = styled(Text.Small)`
    color: var(--brand-primary);
    font-weight: 500;
  `;

  export const TextContent = styled.div`
    max-width: 616px;
    display: flex;
    flex-direction: column;
    gap: 16px;
    margin-top: 192px;

    @media ${tablet} {
      margin-top: 0px;
      width: 100%;
      max-width: 420px;
      align-items: center;
      text-align: center;
    }
  `;

  export const VideoWrapper = styled.div`
    display: flex;
    margin-top: 64px;

    @media ${tablet} {
    }
  `;

  export const Steps = styled.div`
    margin-top: 20px;
    display: grid;
    grid-template-columns: 1fr 1fr 1fr 1fr;
    grid-column-gap: 16px;
    align-self: stretch;
    width: 100%;
    padding: 32px 0px 0px;
    position: relative;

    @media ${mobile} {
      padding: 32px 0;
      grid-template-columns: 1fr;
    }
  `;

  export const StepNumber = styled.div`
    align-items: center;
    background-color: var(--white);
    border-radius: 50%;
    display: flex;
    flex-shrink: 0;
    font-size: 16px;
    font-style: normal;
    font-weight: 700;
    height: 32px;
    justify-content: center;
    line-height: 24px; /* 150% */
    pointer-events: none;
    transition: all 0.2s ease-in-out;
    user-select: none;
    width: 32px;
    filter: drop-shadow(1px 1px 4px rgba(252, 160, 97, 0.1))
      drop-shadow(4px 6px 7px rgba(252, 160, 97, 0.09))
      drop-shadow(9px 12px 9px rgba(252, 160, 97, 0.05))
      drop-shadow(17px 22px 11px rgba(252, 160, 97, 0.01))
      drop-shadow(26px 35px 12px rgba(252, 160, 97, 0));
  `;

  export const StepContentWrapper = styled.div`
    display: flex;
    flex-direction: column;
    gap: 4px;
  `;

  export const StepHeading = styled.h4`
    font-size: 16px;
    font-weight: 500;
    line-height: 24px;
    transition: all 0.2s ease-in-out;
  `;

  export const StepContent = styled.p`
    font-size: 16px;
    font-weight: 400;
    line-height: normal;
    transition: all 0.2s ease-in-out;
  `;

  export const Step = styled.div<{ $isActive?: boolean }>`
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
    align-self: stretch;
    width: 100%;
    max-width: 100%;
    min-height: 90px;
    transition: all 0.2s ease-in-out;
    transform: translateY(0);

    &:hover {
      cursor: pointer;
      transform: translateY(-2px);
      color: red;
    }

    @media ${mobile} {
      flex-direction: row;
    }

    ${StepNumber} {
      color: ${(props) =>
        props.$isActive ? "var(--brand-primary)" : "var(--gray-dark)"};
    }

    ${StepHeading} {
      color: ${(props) =>
        props.$isActive ? "var(--brand-primary)" : "var(--gray-dark)"};
    }

    ${StepContent} {
      color: ${(props) =>
        props.$isActive ? "var(--gray-dark)" : "var(--foreground-light)"};
    }
  `;

  export const Video = styled.video`
    max-width: 100%;
    height: auto;
    border-radius: 24px;
    @media ${mobile} {
      border-radius: 12px;
    }
  `;

  export const ImageWrapper = styled.div`
    display: flex;
    justify-content: center;
    overflow: visible;
    pointer-events: none;
    position: absolute;
    max-height: 892px;
  `;

  export const ContinuityImage = styled(Image)`
    margin-top: -64px;
    @media ${tablet} {
      display: none;
    }
  `;

  export const ContinuityImageMobile = styled(Image)`
    display: none;
    @media ${tablet} {
      margin-left: 8px;
      display: block;
    }
  `;

  export const LineWrapper = styled.div`
    position: absolute;
    padding-right: 10px;
    top: 48px;
    left: 12px;
    right: 0;
    @media ${tablet} {
      ${DottedLine} {
        display: block;
      }
    }
    @media ${mobile} {
      ${DottedLine} {
        display: none;
      }
    }
  `;

  export const ProgressIndicator = styled.div<{
    $progress: number;
    $animate: boolean;
  }>`
    width: 217px;
    height: 4px;
    flex-shrink: 0;
    position: absolute;
    top: 2px;
    background: linear-gradient(
      90deg,
      rgba(254, 189, 58, 0) 0%,
      rgba(254, 189, 58, 1) 35%,
      rgba(252, 160, 97, 1) 100%
    );
    transition: ${(props) => (props.$animate ? "all 0.5s linear" : "none")};

    left: calc(${(props) => props.$progress}% - 217px);

    &:after {
      position: absolute;
      right: 0;
      top: -2px;
      content: "";
      display: block;
      border-radius: 50%;
      height: 8px;
      width: 8px;
      background: rgba(252, 160, 97, 1);
    }

    @media ${mobile} {
      display: none;
    }
  `;

  export const ProgressOverflow = styled.div`
    overflow-x: hidden;
    margin-top: -4px;
    height: 8px;
    width: 100%;
    position: relative;
  `;

  export const ContinuityWrapper = styled.div`
    @media ${tablet} {
      width: 100%;
      display: flex;
      justify-content: center;
    }
  `;

  export const ContentWrapper = styled.div`
    @media ${tablet} {
      display: flex;
      flex-direction: column;
      align-items: center;
    }
  `;
}

export default VideoStepper;
