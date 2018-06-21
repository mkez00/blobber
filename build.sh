#!/bin/sh

# Package for MacOS
echo "Packaging for MacOS"
mkdir blobber-mac
env GOOS=darwin GOARCH=amd64 go build
rm -rf blobber-mac.zip
zip -r -X -j blobber-mac.zip blobber
rm -rf blobber-mac

# Package for Linux
echo "Packaging for Linux"
mkdir blobber-linux
env GOOS=linux GOARCH=amd64 go build
rm -rf blobber-linux.zip
zip -r -X -j blobber-linux.zip blobber
rm -rf blobber-linux

# Package for Windows
echo "Packaging for Windows"
mkdir blobber-win
env GOOS=windows GOARCH=amd64 go build
rm -rf blobber-win.zip
zip -r -X -j blobber-win.zip blobber.exe
rm -rf blobber-win