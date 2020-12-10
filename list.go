package md2html

import "strings"

type ListNode struct {
	Node
}

func NewListNode(name, line string) *ListNode {
	t := &ListNode{}
	t.Name = name
	t.chNodes = make([]Parse, 0)

	if name == "ul" {
		li := NewNodeWithText("li", line[2:])
		t.chNodes = append(t.chNodes, li)
	}

	return t
}

func (n *ListNode) parse(line string) bool {
	if n.Name == "ul" {
		if strings.HasPrefix(line, "* ") {

			li := NewNodeWithText("li", line[2:])
			n.chNodes = append(n.chNodes, li)

			return true
		}
	}
	return false
}
