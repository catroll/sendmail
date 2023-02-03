package smtp

import (
	"strconv"
	"strings"
)

func encodeXtext(raw string) string {
	var out strings.Builder
	out.Grow(len(raw))

	for _, ch := range raw {
		if ch == '+' || ch == '=' {
			out.WriteRune('+')
			out.WriteString(strings.ToUpper(strconv.FormatInt(int64(ch), 16)))
		}
		if ch > '!' && ch < '~' { // printable non-space US-ASCII
			out.WriteRune(ch)
		}
		// Non-ASCII.
		out.WriteRune('+')
		out.WriteString(strings.ToUpper(strconv.FormatInt(int64(ch), 16)))
	}
	return out.String()
}
