package helper

import "strings"

// reverse a string
func Reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

// uppcase a stirng
func Uppcase(s string) string {
	return strings.ToUpper(s)
}

func Modify(s string, opt bool) string {
	if opt == true {
		return s + "_MODIFIED"
	} else {
		return s + "_modified"
	}

}
