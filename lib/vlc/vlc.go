package vlc

import (
	"strings"
	"unicode"
)

func Encode(str string) string {

	return ""
}

func prepareText(str string) string {
	var buf strings.Builder

	var res string

	for _, ch := range str {
		if unicode.IsSpace(ch) {
			res += "!" + string(ch)
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(ch))
		} else {
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}
