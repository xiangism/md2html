package md2html

type AnchorNode struct {
}

type Node interface {
	ToString() string
	CreateAnchor(int) *AnchorNode
}

type Nodes struct {
	chNodes []Node
}

type HeadNode struct {
	Nodes
}

type UlNode struct {
	Nodes
}

type OlNode struct {
	Nodes
}

type LiNode struct {
	Nodes
}

type CodeNode struct {
	Nodes
}

type TextNode struct {
	Nodes
}
