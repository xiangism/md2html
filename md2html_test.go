package md2html

import (
	"fmt"
	"testing"

	"git.lakatv.com/xbc/go-lib/util"
)

func TestMk2ToHtml(t *testing.T) {
	//s := "| | head1 | head2 | head3 |\n|--|--|\n|row1|1.1|1.2|1.3|\n|row2|2.1|2.2|2.3|\n"
	//ls := strings.Split(s, "\n")
	//ls := util.ReadFileLines("t.md")
	ls := util.ReadFileLines("15m.txt")
	html := MarkdownToHtml(ls)
	util.WriteFile("out.html", html)

}

func _TestString(t *testing.T) {
	line := "|head1 | head2 | head3 |head4"
	//rs := strings.Split(line, "|")
	rs := splitTableLine(line)

	for _, item := range rs {
		fmt.Printf("item:(%v)\n", item)
	}
}
