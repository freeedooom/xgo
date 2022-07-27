package xutil

import "strings"

// Blank returns true if trimmed string is empty
func Blank(src string) bool {
	return len(strings.TrimSpace(src)) == 0
}

func GetKeyValue(src string) (key, value string) {
	kv := strings.Split(src, "=")
	if len(kv) == 2 {
		key = kv[0]
		value = kv[1]
	} else {
		key = src
	}

	return
}
