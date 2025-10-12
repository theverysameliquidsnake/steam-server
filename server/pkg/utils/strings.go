package utils

import "strings"

func Strip(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			result.WriteByte(s[i])
		}
	}

	return result.String()
}
