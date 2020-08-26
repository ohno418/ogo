package main

import (
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("Args required")
	}

	input := args[1]
	tok := tokenize(input)
	node := parse(tok)
	codegen(node)
}
