package main

import "fmt"

func main() {
  fmt.Println(".intel_syntax noprefix")
  fmt.Println(".globl main")
  fmt.Println("")
  fmt.Println("main:")
  fmt.Println("  mov rax, 42")
  fmt.Println("  ret")
}
