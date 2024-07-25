import styled from "styled-components";

interface Props {
  $height?: string | number;
  $width?: string | number;
  $top?: boolean;
  $right?: boolean;
  $bottom?: boolean;
  $left?: boolean;
  $display?: string;
  $offsetTop?: number;
  $offsetLeft?: number;
}

const border = "2px dashed #e3e3e3";

const DottedLine = styled.div<Props>`
  height: ${({ $height }) => $height}px;
  width: ${({ $width }) => $width}px;
  display: ${({ $display }) => $display};
  margin-top: ${({ $offsetTop }) => $offsetTop}px;
  margin-left: ${({ $offsetLeft }) => $offsetLeft}px;
  border-top: ${({ $top }) => ($top ? border : "none")};
  border-right: ${({ $right }) => ($right ? border : "none")};
  border-bottom: ${({ $bottom }) => ($bottom ? border : "none")};
  border-left: ${({ $left }) => ($left ? border : "none")};
  border-top-left-radius: ${({ $left, $top }) =>
    $left && $top ? "24px" : "0"};
  border-top-right-radius: ${({ $right, $top }) =>
    $right && $top ? "24px" : "0"};
  border-bottom-left-radius: ${({ $left, $bottom }) =>
    $left && $bottom ? "24px" : "0"};
  border-bottom-right-radius: ${({ $right, $bottom }) =>
    $right && $bottom ? "24px" : "0"};
`;

export default DottedLine;
