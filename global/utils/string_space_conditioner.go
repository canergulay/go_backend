package utils

import "strings"

func StringSpaceConditioner(target string, replaceWith string) string {
	// this guy splits our string into a slice depending on the spaces between words
	words := strings.Fields(target)
	if len(words) == 1 {
		// which means that our string does not have any space, consisting of just one word
		return target
	}
	justString := strings.Join(words, replaceWith)
	return justString
}
