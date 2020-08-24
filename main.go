package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("Args required")
	}

	var input string = args[1]
	tok := tokenize(input)
	debugToken(tok)
	// TODO

	fmt.Println(".intel_syntax noprefix")
	fmt.Println(".globl main")
	fmt.Println("")
	fmt.Println("main:")
	fmt.Println("  mov rax, " + input)
	fmt.Println("  ret")
}
