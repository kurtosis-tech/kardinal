{ pkgs, commit_hash ? "dirty" }: let
  pname = "kardinal.cli";

  # Use the generate-kardinal-version Nix expression
  generateKardinalVersion = pkgs.callPackage ./scripts/generate_kardinal_version.nix {
    inherit pkgs;
  };

  ldflags = pkgs.lib.concatStringsSep "\n" [
    "-X github.com/kurtosis-tech/kurtosis/kardinal.AppName=${pname}"
    "-X github.com/kurtosis-tech/kurtosis/kardinal.Commit=${commit_hash}"
  ];

in
  pkgs.stdenv.mkDerivation {
    name = pname;
    src = ./.;
    buildInputs = [ generateKardinalVersion ] ++ (with pkgs; [
      go
      # Add any other dependencies your build needs
    ]);

    phases = [ "buildPhase" "installPhase" ]; # Ensure buildPhase is included

    buildPhase = ''
      # Run the version generation before building
      ${generateKardinalVersion}/bin/generate-kardinal-version

      # Now build the Go application
      ${pkgs.go}/bin/go build -ldflags "${ldflags}" -o ${pname} ./...
    '';

    installPhase = ''
      # Install the built binary
      mkdir -p $out/bin
      mv ${pname} $out/bin/
    '';

    # This will ensure that the generated file is part of the final output
    installCheckPhase = ''
      # Check if the version file exists and is generated correctly
      if [ ! -f "kardinal_version/kardinal_version.go" ]; then
        echo "Version file not generated!"
        exit 1
      fi
    '';
  }
