"use client";
import { PropsWithChildren } from "react";
import styled from "styled-components";

const Main = ({ children }: PropsWithChildren) => {
  return <S.Main>{children}</S.Main>;
};

namespace S {
  export const Main = styled.main`
    align-items: center;
    display: flex;
    flex-direction: column;
    margin: 0 auto;
  `;
}
export default Main;
