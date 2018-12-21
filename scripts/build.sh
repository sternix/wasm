#!/bin/sh

export GOARCH=wasm GOOS=js

go build .
