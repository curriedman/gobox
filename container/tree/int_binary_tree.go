package tree

import (
	"github.com/recursivecurry/gobox/instance"
)

type IntNode struct {
	instance.Integer
	left Node
	right Node
}

func (i IntNode) Left() Node {
	return i.left
}

func (i *IntNode) SetLeft(n Node) {
	i.left = n
}

func (i IntNode) Right() Node {
	return i.right
}

func (i *IntNode) SetRight(n Node) {
	i.right = n
}


