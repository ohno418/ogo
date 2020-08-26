package main

import (
	"fmt"
	"strconv"
)

func codegen(node *Node) {
	fmt.Println(".intel_syntax noprefix")
	fmt.Println(".globl main")
	fmt.Println("")
	fmt.Println("main:")
	fmt.Println("  mov rax, " + strconv.Itoa(node.val))
	fmt.Println("  ret")
}
