#!/bin/sh
set -e

# Function to handle errors
handle_error() {
    echo "Ops! Failed to setup integration with your shell. Please add the following lines to
your shell configuration manually (changes may not be persistent)
export PATH=\$PATH:$BIN_FOLDER
source <($BIN_FOLDER/$BINARY_NAME completion $PARENT_SHELL)"
    exit 1
}

# Rest of your script goes here
REPO="kurtosis-tech/kardinal"
BINARY_NAME="kardinal"
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)
case "$ARCH" in
    x86_64) ARCH="amd64" ;;
    arm64|aarch64) ARCH="arm64" ;;
esac

BIN_FOLDER="$HOME/.local/bin"
mkdir -p $BIN_FOLDER
LATEST_RELEASE=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
echo "Downloading $BINARY_NAME $LATEST_RELEASE for $OS $ARCH..."
DOWNLOAD_URL="https://github.com/$REPO/releases/download/$LATEST_RELEASE/${BINARY_NAME}-${OS}-${ARCH}"
curl -L $DOWNLOAD_URL -o "$BIN_FOLDER/$BINARY_NAME"
chmod +x "$BIN_FOLDER/$BINARY_NAME"

PARENT_SHELL=$(ps -o comm= -p $PPID)
echo "Detected parent shell: $PARENT_SHELL"

if [ -f "$BIN_FOLDER/$BINARY_NAME" ]; then
    case "$PARENT_SHELL" in
        bash|bash)
            if ! echo "export PATH=\$PATH:$BIN_FOLDER" >> "$HOME/.bashrc" || \
               ! echo "source <($BIN_FOLDER/$BINARY_NAME completion bash)" >> "$HOME/.bashrc"; then
                handle_error
	    else
     		echo "Run the following command to load Kardinal in the current shell (new shell will already load it):"
     		echo "source $HOME/.bashrc"
            fi
            ;;
        zsh|zsh)
            if ! echo "export PATH=\$PATH:$BIN_FOLDER" >> "$HOME/.zshrc" || \
               ! echo "autoload -U +X compinit && compinit" >> "$HOME/.zshrc" || \
               ! echo "source <($BIN_FOLDER/$BINARY_NAME completion zsh)" >> "$HOME/.zshrc"; then
                handle_error
	    else
     		echo "Run the following command to load Kardinal in the current shell (new shell will already load it):"
     		echo "source $HOME/.zshrc"
            fi
            ;;
        fish)
            if ! echo "set -gx PATH \$PATH $BIN_FOLDER" >> "$HOME/.config/fish/config.fish" || \
               ! echo "source <($BIN_FOLDER/$BINARY_NAME completion fish)" >> "$HOME/.config/fish/config.fish"; then
                handle_error
	    else
     		echo "Run the following command to load Kardinal in the current shell (new shell will already load it):"
     		echo "source $HOME/.config/fish/config.fish"
            fi
            ;;
        *)
            echo "Unrecognized shell: $PARENT_SHELL"
            handle_error
            ;;
    esac
else
    handle_error
fi

echo "$BINARY_NAME has been installed successfully!"
