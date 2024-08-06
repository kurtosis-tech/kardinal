#!/usr/bin/env bash

set -e

REPO="kurtosis-tech/kardinal"
BINARY_NAME="kardinal"

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

if [[ "$ARCH" == "x86_64" ]]; then
	ARCH="amd64"
elif [[ "$ARCH" == "arm64" ]]; then
	ARCH="arm64"
elif [[ "$ARCH" == "aarch64" ]]; then
	ARCH="arm64"
fi

BIN_FOLDER="$HOME/.local/bin"
mkdir -p "$BIN_FOLDER"

LATEST_RELEASE=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

echo "Downloading $BINARY_NAME $LATEST_RELEASE for $OS $ARCH..."
DOWNLOAD_URL="https://github.com/$REPO/releases/download/$LATEST_RELEASE/${BINARY_NAME}-${OS}-${ARCH}"
curl -L "$DOWNLOAD_URL" -o "$BIN_FOLDER/$BINARY_NAME"
chmod +x "$BIN_FOLDER/$BINARY_NAME"

USER_SHELL=$(basename "$SHELL")

handle_error() {
	local exit_code=$?
	echo "Ops! Failed to setup integration with your $USER_SHELL shell. Please add the following lines to your shell configuration manually (changes may not be persistent)
export PATH=\$PATH:$BIN_FOLDER
source <($BIN_FOLDER/$BINARY_NAME completion $USER_SHELL)"
	exit $exit_code
}

trap 'handle_error' ERR

# Patch the shell configuration file to include the kardinal binary and completion
patch_shell_config() {
  local user_shell_config_file
  user_shell_config_file=''
  echo "Detected shell: $USER_SHELL"

  # bash and zsh
  common_shell_config() {
    local config_content="
# Kardinal CLI config
export PATH=\$PATH:$BIN_FOLDER
source <($BIN_FOLDER/$BINARY_NAME completion $USER_SHELL)
"
    echo "$config_content">>"$1"
  }

  if [ -f "$BIN_FOLDER/$BINARY_NAME" ]; then
    if [ "$USER_SHELL" == 'bash' ]; then
      # shellcheck disable=SC2088
      user_shell_config_file=~/.bashrc
      common_shell_config $user_shell_config_file
    fi

    if [ "$USER_SHELL" == "zsh" ]; then
      # shellcheck disable=SC2088
      user_shell_config_file=~/.zshrc
      common_shell_config $user_shell_config_file
    fi

    if [ "$USER_SHELL" == "fish" ]; then
      # shellcheck disable=SC2088
      user_shell_config_file=~/.config/fish/config.fish
      local fish_config_content="
# Kardinal CLI config
set -gx PATH \$PATH $BIN_FOLDER
kardinal completion fish | source
"
    echo "$fish_config_content">>"$user_shell_config_file"
    fi
  fi

  echo "Kardinal CLI has been installed successfully!"
  echo "Please reload your shell or 'source $user_shell_config_file' to start using kardinal"
}

patch_shell_config
