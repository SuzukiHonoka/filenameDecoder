package escape

import "regexp"

var RgxEscaped = regexp.MustCompile("%[A-F\\d]{2}")

func IsEscape(s string) bool {
	return RgxEscaped.MatchString(s)
}
