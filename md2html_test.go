package md2html

import (
	"fmt"
	"testing"
)

/*
func TestMk2ToHtml(t *testing.T) {
	ls := util.ReadFileLines("README.md")
	//ls := util.ReadFileLines("lite.md")

	c := NewContent()
	c.ParseLines(ls)
	html := c.Html("test html")

	level := c.HeadLevel()

	util.WriteFile("out.html", html)
	util.WriteFile("out_table.json", level)

}
*/

func _TestString(t *testing.T) {
	line := "|head1 | head2 | head3 |head4"
	//rs := strings.Split(line, "|")
	rs := splitTableLine(line)

	for _, item := range rs {
		fmt.Printf("item:(%v)\n", item)
	}
}
