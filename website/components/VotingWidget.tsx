import { usePathname } from "next/navigation";
import {
  BiDownvote,
  BiSolidDownvote,
  BiSolidUpvote,
  BiUpvote,
} from "react-icons/bi";
import styled from "styled-components";

import Tooltip from "@/components/Tooltip";
import { useVoting, Vote } from "@/context/VotingContext";

// replace anything that is not a letter with an underscore, prefix with "page_"
function normalizePath(path: string) {
  return `page_${(path.startsWith("/") ? path.slice(1) : path).replace(/[^a-z]/g, "_")}`;
}

const VotingWidget = () => {
  const {
    upvote,
    downvote,
    getVoteStatus,
    isUpvoted,
    isDownvoted,
    getVoteCount,
  } = useVoting();
  const pathname = usePathname();
  // const feature = "skylarskardib"; // test feature
  const feature = normalizePath(pathname);
  const voteStatus = getVoteStatus(feature);
  const voteCount = getVoteCount(feature);

  return (
    <S.Wrapper>
      <Tooltip text="Was this page helpful? Use this to provide anonymous feedback">
        <S.VotingWidget $variant={voteStatus}>
          <S.IconButton
            aria-label="upvote feature"
            onClick={() => !isUpvoted(feature) && upvote(feature)}
          >
            {isUpvoted(feature) ? (
              <BiSolidUpvote size={20} />
            ) : (
              <BiUpvote size={20} />
            )}
          </S.IconButton>
          {voteCount >= 5 ? voteCount : "—"}
          <S.IconButton
            aria-label="downvote feature"
            onClick={() => !isDownvoted(feature) && downvote(feature)}
          >
            {isDownvoted(feature) ? (
              <BiSolidDownvote size={20} />
            ) : (
              <BiDownvote size={20} />
            )}
          </S.IconButton>
        </S.VotingWidget>
      </Tooltip>
    </S.Wrapper>
  );
};

namespace S {
  export const Wrapper = styled.div`
    display: flex;
  `;
  export const IconButton = styled.button`
    border: none;
    background: transparent;
    cursor: pointer;
    color: red;
  `;
  export const VotingWidget = styled.div<{ $variant?: Vote }>`
    display: flex;
    align-items: center;
    justify-content: space-evenly;
    border: ${(props) =>
      props.$variant === "none" ? "1px solid #eee" : "1px solid transparent"};
    border-radius: 999px;
    height: 36px;
    width: 98px;
    font-weight: 500;
    font-size: 14px;
    color: var(--white);
    background-color: ${({ $variant }) =>
      $variant === "upvote"
        ? "var(--brand-primary)"
        : $variant === "downvote"
          ? "var(--blue-light)"
          : "white"};
    color: ${({ $variant }) =>
      $variant === "none" ? "var(--gray)" : "var(--white)"};

    ${IconButton} {
      color: ${({ $variant }) =>
        $variant === "none" ? "var(--gray)" : "var(--white)"};
    }
  `;
}

export default VotingWidget;
