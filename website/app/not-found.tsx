"use client";

import { NextPage } from "next";
import Image from "next/image";
import styled from "styled-components";

import notFoundImg from "@/public/not-found.jpg";

const Custom404: NextPage = () => {
  return (
    <S.Container>
      <S.SubHeading>Oops! Page not found.</S.SubHeading>
      <S.ImageContainer>
        <Image
          src={notFoundImg}
          alt="Page not found"
          width={350}
          height={280}
        />
      </S.ImageContainer>
      <S.SubHeading>
        We are sorry but the page you requested was not found.
      </S.SubHeading>
    </S.Container>
  );
};

namespace S {
  export const Container = styled.div`
    height: 69vh;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
    margin-top: 55px;
  `;

  export const SubHeading = styled.div`
    width: 80%;
    display: flex;
    align-items: center;
    justify-content: center;
    text-transform: uppercase;
    color: var(--gray-600);
    font-size: 16px;
    font-weight: 700;
    text-align: center;
  `;

  export const ImageContainer = styled.div``;
}

export default Custom404;
