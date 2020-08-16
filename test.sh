#!/bin/bash

go build -o occ main.go

input=0
./occ $input > tmp.s
gcc --static -o tmp tmp.s
./tmp
output="$?"
echo "${input} => ${output}"

input=42
./occ $input > tmp.s
gcc --static -o tmp tmp.s
./tmp
output="$?"
echo "${input} => ${output}"
