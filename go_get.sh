#!/bin/sh

export GOARCH=wasm GOOS=js

go get -u -v github.com/sternix/wasm
