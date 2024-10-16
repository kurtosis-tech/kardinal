{ pkgs, ... }:
pkgs.writeShellApplication {
  name = "generate-kardinal-version";
  runtimeInputs = with pkgs; [ git, get-docker-tag ];

  text = ''
    set -euo pipefail
    script_dirpath="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
    root_dirpath="$(dirname "${script_dirpath}")"

    KARDINAL_VERSION_PACKAGE_DIR="kardinal_version"
    KARDINAL_VERSION_GO_FILE="kardinal_version.go"
    KARDINAL_VERSION_PACKAGE_NAME="github.com/kurtosis-tech/kardinal/kardinal_version"
    KARDINAL_VERSION_PACKAGE_GOSUM_PATH="go.sum"

    show_helptext_and_exit() {
        echo "Usage: $(basename "$0") new_version"
        echo ""
        echo "  new_version     The version to generate the version constants with, otherwise uses 'get-docker-tag.sh'"
        echo ""
        exit 1
    }

    new_version="${1:-}"

    if [ -z "$new_version" ]; then
        if ! cd "$root_dirpath"; then
            echo "Error: Couldn't cd to the root of this repo, '${root_dirpath}', which is required to get the Git tag" >&2
            show_helptext_and_exit
        fi
        if ! new_version="$(./scripts/get-docker-tag.sh)"; then
            echo "Error: No new version provided and couldn't generate one" >&2
            show_helptext_and_exit
        fi
    fi

    kardinal_version_go_file_abs_path="$root_dirpath/$KARDINAL_VERSION_PACKAGE_DIR/$KARDINAL_VERSION_GO_FILE"

    cat << EOF > "$kardinal_version_go_file_abs_path"
package ${KARDINAL_VERSION_PACKAGE_DIR}

const (
    // !!!!!!!!!!!!!!!!!! DO NOT MODIFY THIS! IT WILL BE UPDATED AUTOMATICALLY DURING THE BUILD PROCESS !!!!!!!!!!!!!!!
    KardinalVersion = "$new_version"
    // !!!!!!!!!!!!!!!!!! DO NOT MODIFY THIS! IT WILL BE UPDATED AUTOMATICALLY DURING THE BUILD PROCESS !!!!!!!!!!!!!!!
)
EOF
  '';
}
