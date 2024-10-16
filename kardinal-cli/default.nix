{
    pkgs,
    commit_hash ? "dirty"
}: let
  pname = "kardinal.cli";

  generateKardinalVersion = pkgs.callPackage ./scripts/generate_kardinal_version.nix {
    inherit pkgs;
  };
  _ = pkgs.runCommand "generate-version" {
    buildInputs = [ generateKardinalVersion ];
  };

  ldflags = pkgs.lib.concatStringsSep "\n" [
    "-X github.com/kurtosis-tech/kurtosis/kardinal.AppName=${pname}"
    "-X github.com/kurtosis-tech/kurtosis/kardinal.Commit=${commit_hash}"
  ];
in
  pkgs.buildGoApplication {
    # pname has to match the location (folder) where the main function is or use
    # subPackages to specify the file (e.g. subPackages = ["some/folder/main.go"];)
    inherit pname ldflags;
    name = "${pname}";
    pwd = ./.;
    src = ./.;
    modules = ./gomod2nix.toml;
    CGO_ENABLED = 0;
  }