#!/bin/sh

find .. -name "*.go" -exec go fmt {} \;
