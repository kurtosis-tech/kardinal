#!/bin/sh

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
mkdir -p $BIN_FOLDER

LATEST_RELEASE=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

echo "Downloading $BINARY_NAME $LATEST_RELEASE for $OS $ARCH..."
DOWNLOAD_URL="https://github.com/$REPO/releases/download/$LATEST_RELEASE/${BINARY_NAME}-${OS}-${ARCH}"
curl -L $DOWNLOAD_URL -o /$BIN_FOLDER/$BINARY_NAME
chmod +x $BIN_FOLDER/$BINARY_NAME

USER_SHELL=$(basename $SHELL)
echo "Detected shell: $USER_SHELL"

handle_error() {
	local exit_code=$?
	echo "Ops! Failed to setup integration with your $USER_SHELL shell. Please add the following lines to 
  your shell configuration manually (changes may not be persistent)
export PATH=\$PATH:$BIN_FOLDER
source <($BIN_FOLDER/$BINARY_NAME completion $USER_SHELL)"
	exit $exit_code
}

trap 'handle_error' ERR

if [ -f $BIN_FOLDER/$BINARY_NAME ]; then
	if [ $USER_SHELL == 'bash' ]; then
		echo "export PATH=\$PATH:$BIN_FOLDER" >>~/.bashrc
		echo "source <($BIN_FOLDER/$BINARY_NAME completion bash)" >>~/.bashrc
		source ~/.bashrc
	fi

	if [ $USER_SHELL == 'zsh' ]; then
		echo "export PATH=\$PATH:$BIN_FOLDER" >>~/.zshrc
		echo "source <($BIN_FOLDER/$BINARY_NAME completion zsh)" >>~/.bashrc
		source ~/.zshrc
	fi

	if [ $USER_SHELL == 'fish' ]; then
		echo "set -gx PATH \$PATH $BIN_FOLDER" >>~/.config/fish/config.fish
		echo "source <($BIN_FOLDER/$BINARY_NAME completion fish)" >>~/.bashrc
		source ~/.config/fish/config.fish
	fi
fi

echo "$BINARY_NAME has been installed successfully!"
