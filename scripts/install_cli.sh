#!/bin/bash

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

LATEST_RELEASE=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

DOWNLOAD_URL="https://github.com/$REPO/releases/download/$LATEST_RELEASE/${BINARY_NAME}-${OS}-${ARCH}"
curl -L $DOWNLOAD_URL -o /tmp/${BINARY_NAME}

BIN_FOLDER="$HOME/.local/bin"
mkdir -p $BIN_FOLDER

mv /tmp/$BINARY_NAME $BIN_FOLDER/$BINARY_NAME

if [ -f $BIN_FOLDER/$BINARY_NAME ]; then
	if [ -d /usr/share/bash-completion/completions ]; then
		source $($BIN_FOLDER/$BINARY_NAME completion bash)
	fi

	if [ -d ~/.zsh/completions ]; then
		source $($BIN_FOLDER/$BINARY_NAME completion zsh)
	fi

	if [ -d ~/.config/fish/completions ]; then
		source $($BIN_FOLDER/$BINARY_NAME completion fish)
	fi
fi

echo "$BINARY_NAME has been installed successfully!"

if ! command -v $BINARY_NAME &>/dev/null; then
	echo "export PATH=\$PATH:$BIN_FOLDER" >>~/.bashrc
	source ~/.bashrc
	echo "export PATH=\$PATH:$BIN_FOLDER" >>~/.zshrc
	source ~/.zshrc
	echo "set -gx PATH \$PATH $BIN_FOLDER" >>~/.config/fish/config.fish
	source ~/.config/fish/config.fish
fi
