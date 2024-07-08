{pkgs, ...}: let
  pyEnv = pkgs.python3.buildEnv.override {
    extraLibs = [pkgs.python3Packages.click pkgs.python3Packages.requests];
    ignoreCollisions = true;
  };

  pname = "demo-load-generator";
  demo-load-genarator = pkgs.stdenv.mkDerivation {
    inherit pname;
    version = "1.0.0";

    src = ./.;

    installPhase = ''
      mkdir -p $out/bin
      echo "#!${pyEnv}/bin/python3" > $out/bin/${pname}
      cat load-generator.py >> $out/bin/${pname}
      chmod +x $out/bin/${pname}
    '';
  };
in
  pkgs.mkShell {
    buildInputs = [
      demo-load-genarator
      pyEnv
    ];
  }
