package md2html

import (
	"fmt"
	"strings"
)

type Content struct {
	top     *Node
	nowNode Parse

	Css string
}

func NewContent() *Content {
	c := &Content{}
	c.top = NewNode("body")
	c.Css = "table\n{border-collapse:collapse;\n    table-layout:fixed;\n}\n\ntable, td, th\n{\n    border:1px solid black;\n    padding: 5px;\n    padding-left: 10px;\n    padding-right: 10px;\n}\n"
	return c
}

func (c *Content) Parse(line string) {

	if c.nowNode != nil {
		if !c.nowNode.parse(line) {
			c.findNode(line)
		}
	} else {
		c.findNode(line)
	}

}

func (c *Content) findNode(line string) {
	line = strings.TrimSpace(line)

	// table
	if strings.HasPrefix(line, "|") && strings.HasSuffix(line, "|") {
		t := NewTableNode(line)

		c.nowNode = t
		c.top.append(t)
		return
	}

	if strings.HasPrefix(line, "* ") {
		// ul
		t := NewListNode("ul", line)

		c.nowNode = t
		c.top.append(t)

		return
	}

	if isOlStart(line) {
		// ol
		t := NewListNode("ol", line)

		c.nowNode = t
		c.top.append(t)

		return
	}
	{ // head
		level, title := parseHead(line)
		if level > 0 {
			t := NewNodeWithText(fmt.Sprintf("h%v", level), title)
			c.top.append(t)
			return
		}
	}
	{ // code
		if line == codeKey {
			t := NewCodeArea()
			c.top.append(t)

			c.nowNode = t
			return
		}

	}

	{ // text
		t := NewNodeWithText("p", line)
		c.top.append(t)

		c.nowNode = nil
	}
}

func (c *Content) Html() string {
	return "<html>" +
		"<head>\n<style type=\"code/css\">" + c.Css + "</style>\n</head>" +
		c.top.toString() +
		"</html>"
}
