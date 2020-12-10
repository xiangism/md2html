package md2html

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Node
}

func NewListNode(name, line string) *ListNode {
	t := &ListNode{}
	t.Name = name
	t.chNodes = make([]Parse, 0)

	t.parse(line)
	return t
}

func (n *ListNode) parse(line string) bool {
	if n.Name == "ul" {
		if strings.HasPrefix(line, "* ") {

			li := NewNodeWithText("li", line[2:])
			n.append(li)

			return true
		}
	} else if n.Name == "ol" {
		if isOlStart(line) {
			li := NewNodeWithText("li", parseOl(line))
			n.append(li)

			return true
		}
	}
	return false
}

func isOlStart(line string) bool {
	index := strings.Index(line, ".")

	if index == -1 {
		return false
	}
	t := line[0:index]

	_, err := strconv.Atoi(t)
	if err == nil {
		return true
	}
	return false
}

func parseOl(line string) string {
	index := strings.Index(line, ".")
	return strings.TrimSpace(line[index+1:])
}
