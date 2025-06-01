package iteration

import "strings"

func Repeat(s string) string {
	var result strings.Builder
	for i := 0; i < 5; i++ {
		result.WriteString(s)
	}
	return result.String()
}
