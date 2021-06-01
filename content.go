package md2html

import (
	"fmt"
	"strings"
)

type Content struct {
	body    *Node
	nowNode Parse

	Css string
}

func NewContent() *Content {
	c := &Content{}
	c.body = NewNode("body")
	c.Css = "body { font-family:  \"Consolas\" ; } \n table\n{border-collapse:collapse;\n table-layout:fixed;\n}\n table, td, th {border:1px solid black;padding: 5px;padding-left: 10px;    padding-right: 10px;}\n" +
		"pre { padding: 4pt; max-width: 100%%white-space; line-height: 1.5; border: 1pt solid #ddd; background-color: #f7f7f7;  }\n" +
		"code { font-family: DejaVu Sans Mono, \\\\5FAE\\\\8F6F\\\\96C5\\\\9ED1; line-height: 1.5; background-color: #f7f7f7; }\n"
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
		c.body.append(t)
		return
	}

	if strings.HasPrefix(line, "* ") {
		// ul
		t := NewListNode("ul", line)

		c.nowNode = t
		c.body.append(t)

		return
	}

	if isOlStart(line) {
		// ol
		t := NewListNode("ol", line)

		c.nowNode = t
		c.body.append(t)

		return
	}
	{ // head
		level, title := parseHead(line)
		if level > 0 {
			t := NewNodeWithText(fmt.Sprintf("h%v", level), title)
			c.body.append(t)
			return
		}
	}
	{ // code
		if line == codeKey {
			t := NewCodeArea()

			c.nowNode = t
			c.body.append(t)

			return
		}

	}

	{ // text
		t := NewNodeWithText("p", line)
		c.body.append(t)

		c.nowNode = nil
	}
}

func (c *Content) Html() string {
	return "<html>" +
		"<head>\n<style type=\"text/css\">" + c.Css + "</style>\n</head>" +
		c.body.toString() +
		"</html>"
}
