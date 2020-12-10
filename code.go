package md2html

type CodeAreaNode struct {
	code     string
	nextStop bool
}

const codeKey = "```"

func NewCodeArea() *CodeAreaNode {
	t := &CodeAreaNode{}
	t.nextStop = false

	return t
}

func (n *CodeAreaNode) parse(line string) bool {

	if n.nextStop {
		return false
	}

	if line == codeKey {
		n.nextStop = true
		// 这里应该是下一行才返回 false

		return true
	}
	n.code += line + "\n"

	return true
}

func (n *CodeAreaNode) toString() string {
	return "<pre><code>" + n.code + "</code></pre>"
}
