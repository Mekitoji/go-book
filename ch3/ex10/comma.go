package comma

import "bytes"

// Comma insert commas in string spliting every 3 symbols
func Comma(s string) string {
	// n := len(s)
	// if n <= 3 {
	// return s
	// }
	// return Comma(s[:n-3]) + "," + s[n-3:]

	var b bytes.Buffer
	n := len(s)
	if n <= 3 {
		return s
	}

	position := n % 3
	if position == 0 {
		position = 3
	}

	for i := 0; i < n; i++ {
		if i == position {
			// write comma and move position
			b.WriteByte(',')
			position += 3
		}
		b.WriteByte(s[i])
	}

	return b.String()
}
