package main

import (
	"fmt"
)

/* Debugger */
func debugToken(tok *Token) {
	cnt := 0
	for {
		fmt.Println("token", cnt, ":", tok)
		tok = tok.next
		cnt++

		if tok == nil {
			break
		}
	}
}
