#!/bin/bash
# PokeSplash Installer for Linux/macOS

INSTALL_DIR="$HOME/.local/bin"
BINARY_NAME="pokesplash"

echo "Installing PokeSplash..."

# Check if binary exists
if [ ! -f "$BINARY_NAME" ]; then
    echo "Error: $BINARY_NAME not found! Please build it first."
    exit 1
fi

# Create install directory
mkdir -p "$INSTALL_DIR"

# Copy binary
cp "$BINARY_NAME" "$INSTALL_DIR/$BINARY_NAME"
chmod +x "$INSTALL_DIR/$BINARY_NAME"

echo "Binary installed to $INSTALL_DIR"

# Detect shell config file
SHELL_CONFIG=""
case "$SHELL" in
  */bash) SHELL_CONFIG="$HOME/.bashrc" ;;
  */zsh) SHELL_CONFIG="$HOME/.zshrc" ;;
  *) SHELL_CONFIG="$HOME/.profile" ;;
esac

# Check if PATH needs update
if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    echo "Adding to PATH in $SHELL_CONFIG..."
    echo "" >> "$SHELL_CONFIG"
    echo "# PokeSplash" >> "$SHELL_CONFIG"
    echo "export PATH=\"\$PATH:$INSTALL_DIR\"" >> "$SHELL_CONFIG"
    echo "Path updated! Please run: source $SHELL_CONFIG"
else
    echo "Path already configured."
fi

echo "Installation Complete! ⚡"
