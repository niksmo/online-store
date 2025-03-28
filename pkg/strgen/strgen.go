package strgen

import "math/rand/v2"

const (
	min = 65
	max = 90
)

func Len(n int) string {
	if n < 1 {
		n = 1
	}

	var buf []byte
	for range n {
		charPos := min + rand.IntN(max-min)
		buf = append(buf, byte(charPos))
	}

	return string(buf)
}
