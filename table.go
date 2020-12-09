package md2html

import "strings"

type TableNode struct {
	head  Node
	body  Node
	index int
}

func NewTableNode(line string) *TableNode {
	t := &TableNode{index: 0}
	// thead
	t.head.Name = "thead"

	rs := splitTableLine(line)

	for _, item := range rs {
		n := NewNodeWithText("th", item)
		t.head.chNodes = append(t.head.chNodes, n)
	}

	t.body.Name = "tbody"

	return t
}

func (n *TableNode) parse(line string) bool {

	if !strings.Contains(line, "|") {
		return false
	}

	if n.index == 0 {
		if strings.Contains(line, "|") {
			n.index = 1
			return true
		}
		return false
	}
	// tbody

	tr := NewNode("tr")

	rs := splitTableLine(line)
	for _, item := range rs {
		td := NewNodeWithText("td", item)
		tr.chNodes = append(tr.chNodes, td)
	}

	n.body.chNodes = append(n.body.chNodes, tr)

	return true
}

func (n *TableNode) toString() string {
	s := "<table>"
	s += n.head.toString()
	s += n.body.toString()
	s += "</table>"
	return s
}

// 分隔 md table一行，拆分成数组
func splitTableLine(line string) []string {
	line = strings.TrimSpace(line)

	// 移除前后 | 后，再用标准 Split 分隔
	if strings.HasPrefix(line, "|") {
		line = line[1:]
	}
	if strings.HasSuffix(line, "|") {
		line = line[0 : len(line)-1]
	}

	rs := strings.Split(line, "|")

	for i, item := range rs {
		rs[i] = strings.TrimSpace(item)
	}

	return rs
}
