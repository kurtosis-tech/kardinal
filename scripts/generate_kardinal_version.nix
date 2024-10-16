{ pkgs, ... }:
pkgs.writeShellApplication {
  name = "generate-kardinal-version";
  runtimeInputs = with pkgs; [ git ];

  text = ''
    set -euo pipefail
    script_dirpath="$(cd "$(dirname "$0")" && pwd)"
    root_dirpath="$(dirname "$script_dirpath")"

    KARDINAL_VERSION_PACKAGE_DIR="kardinal_version"
    KARDINAL_VERSION_GO_FILE="kardinal_version.go"

    if ! git status > /dev/null; then
      echo "Error: This command was run from outside a git repo" >&2
      exit 1
    fi

    commit_sha="$(git rev-parse --short=6 HEAD)"
    suffix="$(git diff --quiet || echo '-dirty')"
    echo "$commit_sha$suffix"

    new_version=""

    if [ -z "$new_version" ]; then
        if ! cd "$root_dirpath"; then
            echo "Error: Couldn't cd to the root of this repo, '$root_dirpath', which is required to get the Git tag" >&2
            show_helptext_and_exit
        fi
        if ! new_version="$(./scripts/get-docker-tag.sh)"; then
            echo "Error: No new version provided and couldn't generate one" >&2
            show_helptext_and_exit
        fi
    fi

    kardinal_version_go_file_abs_path="$root_dirpath/$KARDINAL_VERSION_PACKAGE_DIR/$KARDINAL_VERSION_GO_FILE"

    cat << EOF > "$kardinal_version_go_file_abs_path"
package $KARDINAL_VERSION_PACKAGE_DIR

const (
    // !!!!!!!!!!!!!!!!!! DO NOT MODIFY THIS! IT WILL BE UPDATED AUTOMATICALLY DURING THE BUILD PROCESS !!!!!!!!!!!!!!!
    KardinalVersion = "$new_version"
    // !!!!!!!!!!!!!!!!!! DO NOT MODIFY THIS! IT WILL BE UPDATED AUTOMATICALLY DURING THE BUILD PROCESS !!!!!!!!!!!!!!!
)
EOF
  '';
}
