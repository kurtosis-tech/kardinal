{
  pkgs,
  commit_hash ? "dirty",
}: let
  pname = "redis-proxy-overlay";
  ldflags = pkgs.lib.concatStringsSep "\n" [
    "-X github.com/kurtosis-tech/kurtosis/kardinal.AppName=${pname}"
    "-X github.com/kurtosis-tech/kurtosis/kardinal.Commit=${commit_hash}"
  ];
in
  pkgs.buildGoApplication {
    # pname has to match the location (folder) where the main function is or use
    # subPackges to specify the file (e.g. subPackages = ["some/folder/main.go"];)
    inherit pname ldflags;
    name = "${pname}";
    pwd = ./.;
    src = ./.;
    modules = ./gomod2nix.toml;
    CGO_ENABLED = 0;
  }
