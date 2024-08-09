#!/bin/sh
set -e

# Install Kardinal CLI - supports bash, zsh, fish and assumes you have curl procps installed

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
arm64 | aarch64) ARCH="arm64" ;;
esac

BIN_FOLDER="$HOME/.local/bin"
mkdir -p "$BIN_FOLDER"

WAS_INTALLED_BEFORE=0
if [ -f "$BIN_FOLDER/$BINARY_NAME" ]; then
	WAS_INTALLED_BEFORE=1
fi

LATEST_RELEASE=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
# TODO: Inject the latest release version here during CI
LATEST_RELEASE=${LATEST_RELEASE:-0.1.10}

echo "Downloading $BINARY_NAME $LATEST_RELEASE for $OS $ARCH..."
DOWNLOAD_URL="https://github.com/$REPO/releases/download/$LATEST_RELEASE/${BINARY_NAME}-${OS}-${ARCH}"
curl -L "$DOWNLOAD_URL" -o "$BIN_FOLDER/$BINARY_NAME"
chmod +x "$BIN_FOLDER/$BINARY_NAME"

PARENT_SHELL=$(ps -o comm= -p $PPID)
echo "Detected parent shell: $PARENT_SHELL"

# There are simpler interactive checks such as checking $PS1 or $-,
# but since we are piping the script to sh, these methods don't work
is_interactive_shell() {
  # Check if stdin is a terminal
  if [ -t 0 ]; then
    return 0
  fi
  # Check if stdout is a terminal
  if [ -t 1 ]; then
    return 0
  fi
  # Additional check: see if we can read from /dev/tty
  if [ -t 2 ] && [ -r /dev/tty ]; then
    return 0
  fi
  return 1
}


# Function to prompt user with Y/n question
prompt_user() {
  prompt="$1 [Y/n] "
  if is_interactive_shell; then
    # Use /dev/tty to read input if available
    if [ -r /dev/tty ]; then
      printf "%s" "$prompt" > /dev/tty
      read -r answer < /dev/tty
    else
      printf "%s" "$prompt"
      read -r answer
    fi
  else
    # No terminal available, use a sensible default
    answer="N"
  fi

  case "$answer" in
    [Nn]*) return 1 ;;
    *) return 0 ;;
  esac
}

post_install_questions() {
  # Optionally skip post-install questions and analytics, used in the playground
  if [ -n "$SKIP_KARDINAL_POST_INSTALL" ]; then
    return 0
  fi
  if is_interactive_shell; then
    if prompt_user "Would you like to help us improve Kardinal by anonymously reporting your install?"; then
      "$BIN_FOLDER/$BINARY_NAME" report-install
    else
      echo "No problem! Consider giving us a ⭐ on github to help us grow 😊"
    fi
  else
    echo "Non-interactive shell detected. Skipping post-install questions."
  fi
}

if [ -f "$BIN_FOLDER/$BINARY_NAME" ]; then
	if [ $WAS_INTALLED_BEFORE -eq 0 ]; then
		case "$PARENT_SHELL" in
		-bash | bash)
			CONFIG_FILE="$HOME/.bashrc"
			if ! echo "# Kardinal CLI config" >>"$CONFIG_FILE"; then
				handle_error
			fi
			echo "export PATH=\$PATH:$BIN_FOLDER" >>"$CONFIG_FILE"
			echo "source <($BIN_FOLDER/$BINARY_NAME completion bash)" >>"$CONFIG_FILE"
			;;
		-zsh | zsh)
			CONFIG_FILE="$HOME/.zshrc"
			if ! echo "# Kardinal CLI config" >>"$CONFIG_FILE"; then
				handle_error
			fi
			echo "export PATH=\$PATH:$BIN_FOLDER" >>"$CONFIG_FILE"
			echo "autoload -U +X compinit && compinit" >>"$CONFIG_FILE"
			echo "source <($BIN_FOLDER/$BINARY_NAME completion zsh)" >>"$CONFIG_FILE"
			;;
		-fish | fish)
			CONFIG_FILE="$HOME/.config/fish/config.fish"
			if ! echo "# Kardinal CLI config" >>"$CONFIG_FILE"; then
				handle_error
			fi
			echo "set -gx PATH \$PATH $BIN_FOLDER" >>"$CONFIG_FILE"
			echo "source ($BIN_FOLDER/$BINARY_NAME completion fish | psub)" >>"$CONFIG_FILE"
			;;
		*)
			echo "Unrecognized shell: $PARENT_SHELL"
			handle_error
			;;
		esac
		echo "$BINARY_NAME has been installed successfully!"
		echo "Run the following command to load Kardinal in the current shell (new shells will already load it):"
		echo ""
		echo "> source $CONFIG_FILE"
		echo ""
    post_install_questions
	else
		echo "Kardinal was installed before, just updated it."
		echo ""
    post_install_questions
	fi
else
	echo "Failed to install $BINARY_NAME. Please try again."
	echo ""
	exit 1
fi
