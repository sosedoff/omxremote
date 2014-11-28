#!/bin/bash

NAME="omxremote"
VERSION="0.1.0"
URL="https://github.com/sosedoff/$NAME/releases/download/v$VERSION/$NAME"
BIN_PATH="/usr/bin/$NAME"

echo "Downloading and installing ${NAME} v${VERSION}"
sudo wget -q -O $BIN_PATH $URL
sudo chmod +x $BIN_PATH
echo "Done. Installed into ${BIN_PATH}"