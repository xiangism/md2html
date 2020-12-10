package md2html

func Convert(lines []string) string {
	c := NewContent()

	for _, line := range lines {
		//fmt.Printf("line:(%v)\n", line)
		c.Parse(line)
	}
	return c.Html()
}
