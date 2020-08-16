#!/bin/bash

go build -o occ main.go
./occ > tmp.s
gcc --static -o tmp tmp.s
./tmp
output="$?"
echo $output
