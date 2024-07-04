{pkgs}:
with pkgs;
  stdenv.mkDerivation {
    name = "cli-kontrol-api";
    src = ./.;
    installPhase = ''
      mkdir -p $out
      cp -R ./api $out/
      cp package.json $out/
      cp go.mod $out/
      cp go.sum $out/
      cp README.md $out/
    '';
  }
