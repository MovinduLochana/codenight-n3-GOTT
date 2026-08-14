#!/usr/bin/env bash
set -e

REPO="Mozilla-Campus-Club-of-SLIIT/codenight-n3-GO"

echo "🦊 Downloading Gostlings Go Exercises..."

# 1. Detect OS
OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
case "$OS" in
  darwin) OS="darwin" ;;
  linux)  OS="linux" ;;
  *)
    echo "Unsupported OS: $OS"
    exit 1
    ;;
esac

# 2. Detect Arch
ARCH="$(uname -m)"
case "$ARCH" in
  x86_64)  ARCH="amd64" ;;
  arm64|aarch64) ARCH="arm64" ;;
  *)
    echo "Unsupported architecture: $ARCH"
    exit 1
    ;;
esac

# 3. Fetch latest release tag
LATEST_TAG=$(curl -fsSL "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
if [ -z "$LATEST_TAG" ]; then
  LATEST_TAG="v1.0.0"
fi

FILE_NAME="gostlings-${OS}-${ARCH}.zip"
DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${LATEST_TAG}/${FILE_NAME}"

echo "Fetching ${FILE_NAME} (${LATEST_TAG})..."

TMP_DIR=$(mktemp -d)
curl -fsSL "$DOWNLOAD_URL" -o "${TMP_DIR}/${FILE_NAME}"

unzip -q "${TMP_DIR}/${FILE_NAME}" -d "${TMP_DIR}/extracted"

# Move extracted files to current directory
EXTRACTED_SUBDIR=$(ls "${TMP_DIR}/extracted")
cp -r "${TMP_DIR}/extracted/${EXTRACTED_SUBDIR}/"* .

chmod +x gostlings 2>/dev/null || true

rm -rf "$TMP_DIR"

echo "✅ Setup complete! Run ./gostlings to start learning Go!"
