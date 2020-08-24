package main

type NodeKind int
const (
	ND_NUM NodeKind = iota
	ND_ADD
)

type Node struct {
	kind NodeKind

	// ND_NUM
	val int

	// ND_ADD
	lhs *Node
	rhs *Node
}

// num ("+" num)?
func parse(tok *Token) *Node {
	curTok := tok

	node := numNode(curTok)
	curTok = curTok.next

	if curTok.kind != TK_ADD {
		return node
	}

	node = &Node{kind: ND_ADD, lhs: node, rhs: nil}
	curTok = curTok.next

	node.rhs = numNode(curTok)
	curTok = curTok.next

	if curTok != nil {
		panic("extra token")
	}

	return node
}

func numNode(tok *Token) *Node {
	if tok.kind != TK_NUM {
		panic("number expected")
	}

	return &Node{kind: ND_NUM, val: tok.val}
}
