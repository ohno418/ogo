#!/bin/bash

assert()
{
  input="$1"
  expected="$2"

  ./ogo "$input" > tmp.s
  gcc --static -o tmp tmp.s
  ./tmp
  actual="$?"

  if [ "$actual" = "$expected" ]
  then
    echo "\"$input\" => $actual"
  else
    echo "\"$input\" => $expected expected, but got $actual"
    exit 1
  fi
}

go build -o ogo main.go

assert 0 0
assert 42 42

echo OK
exit 0
