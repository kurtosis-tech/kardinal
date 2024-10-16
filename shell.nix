{pkgs}: let
  mergeShells = shells:
    pkgs.mkShell {
      shellHook = builtins.concatStringsSep "\n" (map (s: s.shellHook or "") shells);
      buildInputs = builtins.concatLists (map (s: s.buildInputs or []) shells);
      nativeBuildInputs = builtins.concatLists (map (s: s.nativeBuildInputs or []) shells);
      paths = builtins.concatLists (map (s: s.paths or []) shells);
    };

  kardinal = pkgs.writeShellScriptBin "kardinal" ''
    nix run .#kardinal-cli -- "$@"
  '';

  go-tidy-all = import ./scripts/go-tidy-all.nix {inherit pkgs;};
  generate-kardinal-version = import ./kardinal-cli/scripts/generate_kardinal_version.nix {inherit pkgs;};
  manager_shell = pkgs.callPackage ./kardinal-manager/shell.nix {inherit pkgs;};
  cli_shell = pkgs.callPackage ./kardinal-cli/shell.nix {inherit pkgs;};
  cli_kontrol_api_shell = pkgs.callPackage ./libs/cli-kontrol-api/shell.nix {inherit pkgs;};
  demo_shell = pkgs.callPackage ./examples/voting-app/shell.nix {inherit pkgs;};
  website_shell = pkgs.callPackage ./website/shell.nix {inherit pkgs;};

  kardinal_shell = with pkgs;
    pkgs.mkShell {
      nativeBuildInputs = [bashInteractive bash-completion];
      buildInputs = [
        kardinal
        go-tidy-all
        generate-kardinal-version
        kubectl
        kustomize
        kubernetes-helm
        minikube
        istioctl
        tilt
        reflex
      ];
      shellHook = ''
        export SHELLNAME=$(basename $shell)
        source <(kubectl completion $SHELLNAME)
        source <(minikube completion $SHELLNAME)
        source <(kardinal completion $SHELLNAME)
        printf '\u001b[31m

                                          :::::
                                           :::::::
                                           ::   :::
                                          :::     ::
                                          ::   ::- :::
                                        :::         :::
                                       ::: :::    :::
                                     :::    ::    ::
                                   :::      ::   :::
                                 :::       :::   ::
                               :::        ::     ::
                            ::::       ::::     ::
                          ::::      ::::      :::
                       ::::::::::::::       ::::
                                       ::::::
                   :::::::::::::::::::::
               ::::::
            :::::
          :::



        \u001b[0m
        Starting Kardinal dev shell.
        \e[32m
        \e[0m
        '
      '';
    };
in
  mergeShells [
    manager_shell
    cli_shell
    kardinal_shell
    cli_kontrol_api_shell
    demo_shell
    website_shell
  ]
