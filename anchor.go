package md2html

import "fmt"

type AnchorNode struct {
	Id string
}

func NewAnchor(id string) *AnchorNode {
	a := &AnchorNode{Id: id}
	return a
}

func (a *AnchorNode) toString() string {
	return fmt.Sprintf("<a id=\"%v\"></a>", a.Id)
}

func (a *AnchorNode) parse(s string) bool {
	return false
}

/*
找其anchorn上面小于它的节点做为其父节点
1
2
3
2
1
1
3
2
1
*/
type HeadLevel struct {
	Id       string       `json:"id"`
	Name     string       `json:"name"`
	Children []*HeadLevel `json:"children"`

	level int
}

func NewHeadLevel(id, name string, l int) *HeadLevel {
	t := &HeadLevel{Id: id, Name: name, level: l}
	return t
}

func (l *HeadLevel) addChild(c *HeadLevel) {
	l.Children = append(l.Children, c)
}

func (c *Content) parseHeadLevel() *HeadLevel {
	for i := 1; i < len(c.alllevel); i += 1 {
		item := c.alllevel[i]

		p := c.findParent(i-1, item.level)
		p.addChild(item)
	}

	return c.alllevel[0]
}

func (c *Content) findParent(lastindex int, myLevel int) *HeadLevel {
	for i := lastindex; i >= 0; i -= 1 {
		if c.alllevel[i].level < myLevel {
			return c.alllevel[i]
		}
	}
	return c.alllevel[0]
}
