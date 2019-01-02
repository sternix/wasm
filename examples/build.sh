#!/bin/sh

# this compiles to native
go build -o server/server ./server/server.go 

DIRS='
ball_drop_game
dom
life
webgl/rotating-cube
webgl/splashy
webgl/triangle
w3schools/accordion
w3schools/clickable_dropdown
w3schools/modal
w3schools/slideshow
w3schools/draggable
w3schools/closable_list_items
w3schools/filter_list
w3schools/filter_table
w3schools/autocomplete
'

OUT_DIR="./server/assets"

export GOOS=js GOARCH=wasm

for d in $DIRS; do
	CUR_DIR=`basename $d`
	echo "Compiling: ${CUR_DIR}"
	OUT_FILE="${OUT_DIR}/${CUR_DIR}.wasm"
	SRC_FILES="$d/*.go"
	go build -o $OUT_FILE $SRC_FILES
	if [ $? -ne 0 ]; then
		exit
	fi
done
