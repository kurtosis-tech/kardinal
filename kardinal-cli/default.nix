{
  pkgs,
  kardinal_version,
  commit_hash ? "dirty",
}: let
  pname = "kardinal.cli";
  kardinal_version="1.0.0";

  # The CLI fails to compile as static using CGO_ENABLE (macOS and Linux). We need to manually use flags and add glibc
  # More info on: https://nixos.wiki/wiki/Go (also fails with musl!)
  static_linking_config = with pkgs; if stdenv.isLinux then {
    buildInputs = [ glibc.static ];
    nativeBuildInputs = [ stdenv ];
    CFLAGS = "-I${glibc.dev}/include";
    LDFLAGS = "-L${glibc}/lib";
  } else
    { };

  static_ldflag = with pkgs; if stdenv.isLinux then
    [ "-s -w -linkmode external -extldflags -static" ]
  else
    [ ];

  ldflags = builtins.trace kardinal_version pkgs.lib.concatStringsSep "\n" (static_ldflag ++ [
    "-X kardinal/kardinal_version.KardinalVersion=${kardinal_version}"
  ]);
in
  pkgs.buildGoApplication ({
    # pname has to match the location (folder) where the main function is or use
    # subPackges to specify the file (e.g. subPackages = ["some/folder/main.go"];)
    inherit pname ldflags;
    name = "${pname}";
    version="${kardinal_version}";
    pwd = ./.;
    src = ./.;
    modules = ./gomod2nix.toml;
    CGO_ENABLED = 0;
  } // static_linking_config)
