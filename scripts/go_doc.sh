#!/bin/sh

export GOARCH=wasm GOOS=js

godoc -http=:6060
