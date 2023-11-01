#!/bin/bash

VERSION="v0.1.0"

WINDOWS="./build/windows-x86-64bit-$VERSION"
LINUX="./build/linux-x86-64bit-$VERSION"
MAC_OS_ARM="./build/macOs-arm-64bit-$VERSION"
MAC_OS="./build/macOs-x86-64bit-$VERSION"
MAIN_FILE="./main.go"
APP_NAME="webPer"

export GOOS=windows
go build -o "$WINDOWS/$APP_NAME.exe" $MAIN_FILE
zip -r $WINDOWS.zip $WINDOWS/

export GOOS=linux
go build -o "$LINUX/$APP_NAME" $MAIN_FILE
zip -r $LINUX.zip $LINUX/

export GOOS=darwin
go build -o "$MAC_OS/$APP_NAME" $MAIN_FILE
zip -r $MAC_OS.zip $MAC_OS/

export GOARCH=arm64
go build -o "$MAC_OS_ARM/$APP_NAME" $MAIN_FILE
zip -r $MAC_OS_ARM.zip $MAC_OS_ARM/