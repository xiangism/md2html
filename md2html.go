package md2html

func Convert(lines []string, title string) string {
	c := NewContent()
	c.ParseLines(lines)
	return c.Html(title)
}
