#!/bin/bash

set -e

DOWNLOAD_URL=https://github.com/Angryman18/filesystem-size/releases/download/1.0.0/gtds
INSTALL_PATH="/usr/local/bin"


echo "Downloading gtds from github release"
if ! curl -sSLf "$DOWNLOAD_URL" -o "$INSTALL_PATH/gtds"; then
    echo "Failed to Download & Install gtds please try with sudo"
    exit 1
fi

if ! chmod +x "$INSTALL_PATH/gtds"; then
    echo "Failed"
    exit 1
fi

echo "gtds downloaded successfully!"