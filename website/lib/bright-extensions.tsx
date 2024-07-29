// source: https://github.com/code-hike/bright/tree/main/web/app/recipes/title-bar

import type { BrightProps, Extension } from "bright";

export const titleBar: Extension = {
  name: "titleBar",
  TitleBarContent: Title,
};

function Title(props: BrightProps) {
  const { title, colors } = props;
  const { foreground, background } = colors;

  const circle = {
    borderRadius: "100%",
    height: "0.8em",
    width: "0.8em",
    background: foreground,
    opacity: 0.2,
  };

  return (
    <div
      style={{
        background,
        color: foreground,
        padding: "3px 0",
        fontSize: "0.9em",
        display: "flex",
        alignItems: "center",
        justifyContent: "space-between",
        width: "100%",
      }}
    >
      <div style={{ gap: 4, display: "flex", marginLeft: 8 }}>
        <div style={circle} />
        <div style={circle} />
        <div style={circle} />
      </div>
      <span style={{ opacity: 0.8 }}>{title}</span>
      <div style={{ width: 45 }} />
    </div>
  );
}
