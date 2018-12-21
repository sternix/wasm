#!/bin/sh

for f in `find . -name "*.go" -print` ; do
	head -1 $f | grep -v "build js,wasm" >/dev/null 2>&1
	if [ $? -eq 0 ]; then
		echo $f
	fi
done
