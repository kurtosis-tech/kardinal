{pkgs, ...}:
pkgs.writeShellApplication {
  name = "go-tidy-all";
  runtimeInputs = with pkgs; [go git gomod2nix];
  text = ''
    root_dirpath=$(git rev-parse --show-toplevel)
    find "$root_dirpath" -type f -name 'go.mod' -exec sh -c 'dir=$(dirname "$1") && cd "$dir" && echo "$dir" && go mod tidy && gomod2nix' shell {} \;
  '';
}
