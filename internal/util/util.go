package util

import "strings"

func RemoveName(input, name string) string {
	condensedStr := strings.ToLower(input)
	if !strings.Contains(condensedStr, name) {
		return ""
	}

	// Find the index of the name
	return strings.Replace(input, name, "", -1)
}
