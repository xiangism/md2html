package md2html

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Content struct {
	body    *Node
	nowNode Parse

	Css string

	anchorId int

	alllevel []*HeadLevel
}

func NewContent() *Content {
	c := &Content{}
	c.body = NewNode("body")
	c.Css = "body { font-family:  \"Consolas\" ; } \n table\n{border-collapse:collapse;\n table-layout:fixed;\n}\n table, td, th {border:1px solid black;padding: 5px;padding-left: 10px;    padding-right: 10px;}\n" +
		"pre { padding: 4pt; max-width: 100%%white-space; line-height: 1.5; border: 1pt solid #ddd; background-color: #f7f7f7;  }\n" +
		"code { font-family: DejaVu Sans Mono, \\\\5FAE\\\\8F6F\\\\96C5\\\\9ED1; line-height: 1.5; background-color: #f7f7f7; }\n"

	top := NewHeadLevel("", "top", 0)
	c.alllevel = []*HeadLevel{top}

	return c
}

func (c *Content) ParseLines(lines []string) {
	for _, line := range lines {
		c.parse(line)
	}
}

func (c *Content) Html(title string) string {
	return "<html>" +
		"<head><title>" + title + "</title> <meta charset=\"UTF-8\"> \n<style type=\"text/css\">" + c.Css + "</style>\n</head>" +
		"<meta name=\"viewport\" content=\"width=device-width,initial-scale=1\"> <meta name=\"apple-touch-fullscreen\" content=\"YES\"> <meta name=\"apple-mobile-web-app-capable\" content=\"yes\">" +
		"<body>" +
		c.body.toString() +
		"</body></html>"
}

/*
返回 head 结构的层次，json 结构
*/
func (c *Content) HeadLevel() string {
	t := c.parseHeadLevel()

	if t == nil {
		return "[]"
	}
	if len(t.Children) == 0 {
		return "[]"
	}

	bs, _ := json.Marshal(t.Children)

	return string(bs)
}

func (c *Content) parse(line string) {

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
			id := c.getAnchorId()
			a := NewAnchor(id)
			c.body.append(a)

			tree := NewHeadLevel(id, title, level)
			c.alllevel = append(c.alllevel, tree)

			t := NewNodeWithText(fmt.Sprintf("h%v", level), title)
			c.body.append(t)
			return
		}
	}
	{ // code
		if strings.HasPrefix(line, codeKey) {
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

func (c *Content) getAnchorId() string {
	s := fmt.Sprintf("md_%v", c.anchorId)
	c.anchorId += 1
	return s
}
