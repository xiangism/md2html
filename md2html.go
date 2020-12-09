package md2html

func MarkdownToHtml(lines []string) string {
	c := NewContent()

	for _, line := range lines {
		//fmt.Printf("line:(%v)\n", line)
		c.Parse(line)
	}
	return c.Html()
}
