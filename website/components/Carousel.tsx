"use client";
import Image from "next/image";
import { useState } from "react";
import styled from "styled-components";

const Carousel = ({ imageUrls }: { imageUrls: string[] }) => {
  const [currentImageIndex, setCurrentImageIndex] = useState(0);

  const handleDotClick = (index: number) => {
    setCurrentImageIndex(index);
  };

  return (
    <S.Carousel>
      <S.Container>
        <Image
          src={imageUrls[currentImageIndex]}
          alt="Logo"
          unoptimized
          height={446}
          width={550}
          style={{
            width: "100%",
            height: "auto",
          }}
        />
        <S.Dots>
          {imageUrls.map((_, index) => (
            <S.Dot
              key={index}
              $active={index === currentImageIndex}
              onClick={() => handleDotClick(index)}
            />
          ))}
        </S.Dots>
      </S.Container>
    </S.Carousel>
  );
};

namespace S {
  export const Carousel = styled.div`
    display: flex;
    flex-direction: column;
    gap: 16px;
    width: 100%;
  `;

  export const Container = styled.div`
    width: 100%;
    display: flex;
    align-items: center;
    flex-direction: column;
    gap: 16px;
    justify-content: space-between;
    padding: 16px 0;
    overflow: hidden;

    img {
      flex-shrink: 0;
    }
  `;

  export const Dots = styled.div`
    display: flex;
    justify-content: center;
    margin-top: 10px;
  `;

  export const Dot = styled.div<{ $active: boolean }>`
    height: 16px;
    width: 16px;
    border-radius: 50%;
    background: ${(props) =>
      props.$active ? "var(--brand-secondary)" : "gray"};
    margin: 0 5px;
    cursor: pointer;
  `;
}

export default Carousel;
