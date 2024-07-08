#!/bin/bash

REPO="owner/repo"         # Replace with the actual repository, e.g., "cli/cli"
BINARY_NAME="binary_name" # Replace with the actual binary name, e.g., "gh"

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

if [[ "$ARCH" == "x86_64" ]]; then
	ARCH="amd64"
elif [[ "$ARCH" == "arm64" ]]; then
	ARCH="arm64"
elif [[ "$ARCH" == "aarch64" ]]; then
	ARCH="arm64"
elif [[ "$ARCH" == "i386" ]]; then
	ARCH="386"
fi

LATEST_RELEASE=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

DOWNLOAD_URL="https://github.com/$REPO/releases/download/$LATEST_RELEASE/${BINARY_NAME}-${OS}-${ARCH}.tar.gz"
curl -L $DOWNLOAD_URL -o /tmp/${BINARY_NAME}.tar.gz

tar -xzf /tmp/${BINARY_NAME}.tar.gz -C /tmp
sudo mv /tmp/$BINARY_NAME /usr/local/bin/$BINARY_NAME

rm /tmp/${BINARY_NAME}.tar.gz

if [ -f /usr/local/bin/$BINARY_NAME ]; then
	if [ -d /usr/share/bash-completion/completions ]; then
		sudo curl -L "https://raw.githubusercontent.com/$REPO/$LATEST_RELEASE/completions/$BINARY_NAME.bash" -o /usr/share/bash-completion/completions/$BINARY_NAME
		source /usr/share/bash-completion/completions/$BINARY_NAME
	fi

	if [ -d ~/.zsh/completions ]; then
		sudo curl -L "https://raw.githubusercontent.com/$REPO/$LATEST_RELEASE/completions/_$BINARY_NAME" -o ~/.zsh/completions/_$BINARY_NAME
		fpath=(~/.zsh/completions $fpath)
		autoload -Uz compinit && compinit
	elif [ -d /usr/local/share/zsh/site-functions ]; then
		sudo curl -L "https://raw.githubusercontent.com/$REPO/$LATEST_RELEASE/completions/_$BINARY_NAME" -o /usr/local/share/zsh/site-functions/_$BINARY_NAME
	fi

	if [ -d ~/.config/fish/completions ]; then
		sudo curl -L "https://raw.githubusercontent.com/$REPO/$LATEST_RELEASE/completions/$BINARY_NAME.fish" -o ~/.config/fish/completions/$BINARY_NAME.fish
	elif [ -d /usr/share/fish/vendor_completions.d ]; then
		sudo curl -L "https://raw.githubusercontent.com/$REPO/$LATEST_RELEASE/completions/$BINARY_NAME.fish" -o /usr/share/fish/vendor_completions.d/$BINARY_NAME.fish
	fi
fi

echo "$BINARY_NAME has been installed successfully!"

if ! command -v $BINARY_NAME &>/dev/null; then
	echo "export PATH=\$PATH:/usr/local/bin" >>~/.bashrc
	source ~/.bashrc
	echo "export PATH=\$PATH:/usr/local/bin" >>~/.zshrc
	source ~/.zshrc
	echo "set -gx PATH \$PATH /usr/local/bin" >>~/.config/fish/config.fish
	source ~/.config/fish/config.fish
fi
