{pkgs}: let
  goEnv = pkgs.mkGoEnv {pwd = ./.;};
in
  pkgs.mkShell {
    nativeBuildInputs = with pkgs; [
      goEnv

      goreleaser
      go
      gopls
      golangci-lint
      delve
      enumer
      gomod2nix
      bash-completion
    ];
  }
