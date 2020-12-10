package md2html

func Convert(lines []string) string {
	c := NewContent()

	for _, line := range lines {
		c.Parse(line)
	}
	return c.Html()
}
