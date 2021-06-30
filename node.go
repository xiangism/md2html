package md2html

import "fmt"

type Parse interface {
	toString() string

	parse(string) bool
}

// <Name> Text </Name>
type Node struct {
	Name string

	Text    string
	chNodes []Parse
}

func NewNode(name string) *Node {
	n := &Node{Name: name}
	n.chNodes = make([]Parse, 0)
	return n
}

func NewNodeWithText(name, text string) *Node {
	n := &Node{Name: name, Text: text}
	n.chNodes = make([]Parse, 0)
	return n
}

func (n *Node) toString() string {
	v := ""

	if len(n.chNodes) > 0 {
		for _, item := range n.chNodes {
			v += item.toString()
		}

	} else {
		v = n.Text
	}

	return fmt.Sprintf("<%v>%v</%v>\n", n.Name, v, n.Name)
}

func (n *Node) parse(line string) bool {
	return false
}

func (n *Node) append(c Parse) {
	n.chNodes = append(n.chNodes, c)
}
