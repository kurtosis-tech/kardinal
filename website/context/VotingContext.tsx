"use client";

import React, {
  createContext,
  PropsWithChildren,
  useContext,
  useEffect,
} from "react";

interface Feature {
  name: string;
  upvotes: number;
  downvotes: number;
}

interface LocalStorageVote {
  feature: string;
  vote: "upvote" | "downvote";
}

export type Vote = "upvote" | "downvote" | "none";

interface VotingContextProps {
  localStorageVotes: LocalStorageVote[]; // features the current user has voted on
  features: Feature[];
  // eslint-disable-next-line no-unused-vars
  upvote: (feature: string) => Promise<void>;
  // eslint-disable-next-line no-unused-vars
  downvote: (feature: string) => Promise<void>;
  // eslint-disable-next-line no-unused-vars
  getVoteStatus: (feature: string) => Vote;
  // eslint-disable-next-line no-unused-vars
  isUpvoted: (feature: string) => boolean;
  // eslint-disable-next-line no-unused-vars
  isDownvoted: (feature: string) => boolean;
  // eslint-disable-next-line no-unused-vars
  getVoteCount: (feature: string) => number;
}

const VotingContext = createContext<VotingContextProps | undefined>(undefined);

const doApiRequest = async (
  method: "GET" | "POST",
  route: string,
  feature: string,
) => {
  const baseUrl = `${process.env.NEXT_PUBLIC_DOCS_VOTING_API}/${route}`;
  // avoid trailing slash which causes a 301 redirect
  const url = feature ? `${baseUrl}/${feature}` : baseUrl;

  try {
    const response = await fetch(url, {
      method,
      headers: {
        "Content-Type": "application/json",
        // "http://localhost:3000" or "https://kardinal.dev" etc,
        // This Origin header must be set for the server to respond with an
        // access-control-allow-origin header
        Origin: `${window.location.protocol}//${window.location.host}`,
      },
      mode: "cors",
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error("There was an error with the fetch operation: ", error);
  }
};

export const VotingProvider = ({ children }: PropsWithChildren) => {
  const [localStorageVotes, setLocalStorageVotes] = React.useState<
    LocalStorageVote[]
  >(
    typeof window === "undefined"
      ? []
      : JSON.parse(localStorage.getItem("votes") || "[]"),
  );
  const [features, setFeatures] = React.useState<Feature[]>([]);

  const fetchFeatures = async () => {
    const response = await doApiRequest("GET", "features", "");
    setFeatures(response);
  };

  useEffect(() => {
    fetchFeatures();
  }, []);

  useEffect(() => {
    localStorage.setItem("votes", JSON.stringify(localStorageVotes));
  }, [localStorageVotes]);

  const downvote = async (feature: string) => {
    setLocalStorageVotes([
      ...localStorageVotes.filter((vote) => vote.feature !== feature),
      { feature, vote: "downvote" },
    ]);
    // if user has already voted on this feature, prevent them from duplicating votes
    if (getVoteStatus(feature) !== "none") {
      return;
    }
    await doApiRequest("POST", "downvote", feature);
    fetchFeatures(); // update vote count
  };

  const upvote = async (feature: string) => {
    setLocalStorageVotes([
      ...localStorageVotes.filter((vote) => vote.feature !== feature),
      { feature, vote: "upvote" },
    ]);
    // if user has already voted on this feature, prevent them from duplicating votes
    if (getVoteStatus(feature) !== "none") {
      return;
    }
    await doApiRequest("POST", "upvote", feature);
    fetchFeatures(); // update vote count
  };

  const getVoteCount = (feature: string) => {
    return features.find((f) => f.name === feature)?.upvotes || 0;
  };

  const getVoteStatus = (feature: string) => {
    if (
      localStorageVotes.some(
        (vote) => vote.feature === feature && vote.vote === "upvote",
      )
    ) {
      return "upvote";
    }
    if (
      localStorageVotes.some(
        (vote) => vote.feature === feature && vote.vote === "downvote",
      )
    ) {
      return "downvote";
    }
    return "none";
  };
  const isUpvoted = (feature: string) => getVoteStatus(feature) === "upvote";
  const isDownvoted = (feature: string) =>
    getVoteStatus(feature) === "downvote";

  return (
    <VotingContext.Provider
      value={{
        features,
        upvote,
        downvote,
        getVoteStatus,
        getVoteCount,
        isUpvoted,
        isDownvoted,
        localStorageVotes,
      }}
    >
      {children}
    </VotingContext.Provider>
  );
};

export const useVoting = (): VotingContextProps => {
  const context = useContext(VotingContext);
  if (!context) {
    throw new Error("useVoting must be used within a VotingProvider");
  }
  return context;
};
