package main

import (
	"fmt"
	"regexp"
	"strconv"
)

const (
	TK_NUM = iota
	TK_ADD
)

type Token struct {
	Next *Token
	Kind int

	// Number
	Num int
}

func tokenize(input string) *Token {
	var rest string = input

	var head Token = Token{}
	cur := &head

	for {
		if rest == "" {
			break
		}

		c := string(rest[0])

		if c == " " {
			rest = rest[1:len(rest)]
			continue
		}

		if isNum(c) {
			num, _ := getNum(rest)
			cur.Next = &Token{Kind: TK_NUM, Num: num}
			cur = cur.Next
			rest = rest[len(strconv.Itoa(num)):len(rest)]
			continue
		}

		if c == "+" {
			cur.Next = &Token{Kind: TK_ADD}
			cur = cur.Next
			rest = rest[1:len(rest)]
			continue
		}

		raiseUnexpectedChar(input, len(input)-len(rest))
	}

	return head.Next
}

func isNum(str string) bool {
	switch str {
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		return true
	default:
		return false
	}
}

func getNum(str string) (int, error) {
	reg := regexp.MustCompile(`^\d+`)
	return strconv.Atoi(string(reg.Find([]byte(str))))
}

func raiseUnexpectedChar(input string, i int) {
	fmt.Println(input)
	for j := 0; j < i; j++ {
		fmt.Print(" ")
	}
	fmt.Println("^cannot recognize this character")
	panic("unexpected character")
}
