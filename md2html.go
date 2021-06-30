package md2html

func Convert(lines []string) string {
	c := NewContent()
	c.ParseLines(lines)
	return c.Html()
}
