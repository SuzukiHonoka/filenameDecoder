package rename

import "strings"

type FileNameFilter []rune

var replacer *strings.Replacer

func init() {
	tokens := make([]string, 0, len(Filter)*2)
	for _, c := range Filter {
		tokens = append(tokens, string(c))
		tokens = append(tokens, "")
	}
	replacer = strings.NewReplacer(tokens...)
}

func DoFilter(s string) string {
	return replacer.Replace(s)
}
