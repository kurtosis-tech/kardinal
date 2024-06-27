{pkgs ? import <nixpkgs> {} }: let
  goEnv = pkgs.mkGoEnv {pwd = ./.;};
  node-devtools = import ./nix/. {
    inherit pkgs;
    nodejs = pkgs.nodejs_20;
  };
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

      oapi-codegen
      nodejs
      node2nix
      node-devtools.nodeDependencies
    ];
  }
