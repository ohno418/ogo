#!/bin/bash

rm tmp.s
go run main.go >> tmp.s
gcc --static -o main tmp.s
./main
echo $?
