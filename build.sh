#!/usr/bin/env bash
set -e

OUT=dist
BIN=pggoose

mkdir -p $OUT

build() {
  GOOS=$1
  GOARCH=$2
  EXT=$3

  echo "Building $GOOS/$GOARCH"
  CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH \
    go build -ldflags="-s -w" \
    -o $OUT/$BIN-$GOOS-$GOARCH$EXT \
    ./main.go
}

build linux amd64 ""
build linux arm64 ""
build darwin amd64 ""
build darwin arm64 ""
