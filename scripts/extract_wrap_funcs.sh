#!/bin/sh

fgrep 'func wrap' *.go | cut -d: -f 2 | awk '/ wrap/{print substr($2,0,length($2) -2)}' | sort | uniq
