package md2html

import "strings"

var headKey = []string{"######", "#####", "####", "###", "##", "#"}

// 返回head的层级, 1-6
// 如果不是head，则返回0
// @return (leven, title)
func parseHead(line string) (int, string) {
	if !strings.HasPrefix(line, "#") {
		return 0, ""
	}
	for _, key := range headKey {
		if strings.HasPrefix(line, key) {
			l := len(key)
			return l, line[l:]
		}
	}
	return 0, ""
}
