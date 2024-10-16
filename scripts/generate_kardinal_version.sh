#!/bin/zsh

set -euo pipefail
root_dirpath=$(git rev-parse --show-toplevel)

KARDINAL_VERSION_PACKAGE_DIR="kardinal_version"
KARDINAL_VERSION_GO_FILE="kardinal_version.go"

# Check if the current commit has a tag
if git describe --tags --exact-match > /dev/null 2>&1; then
    # If there's an exact match to a tag, use the tag as the version
    new_version="$(git describe --tags --exact-match)"
else
    # Otherwise, use the short commit SHA
    commit_sha="$(git rev-parse --short=6 HEAD)"
    # Check if the working directory is dirty
    suffix="$(git diff --quiet || echo '-dirty')"
    new_version="$commit_sha$suffix"
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