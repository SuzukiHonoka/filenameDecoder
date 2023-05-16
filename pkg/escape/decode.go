package escape

import "net/url"

func Decode(s string) (string, error) {
	return url.QueryUnescape(s)
}
