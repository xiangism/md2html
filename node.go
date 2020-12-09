package md2html

import "fmt"

type AnchorNode struct {
}

type Parse interface {
	toString() string

	parse(string) bool

	//CreateAnchor(int) *AnchorNode
}

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

	return fmt.Sprintf("<%v>%v</%v>", n.Name, v, n.Name)
}

func (n *Node) parse(line string) bool {
	return false
}

// <h1> </h1>
type HeadNode struct {
	Node
}

// <ul>
type UlNode struct {
	Node
}

// <ol>
type OlNode struct {
	Node
}

// <li>
type LiNode struct {
	Node
}

type CodeNode struct {
	Node
}

type TextNode struct {
	Node
}
