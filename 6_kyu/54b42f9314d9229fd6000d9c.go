package kata

import "strings"

func DuplicateEncode(word string) string {
	word = strings.ToLower(word)
	var cnts [128]byte // ASCII
	for _, c := range word {
		cnts[c]++
	}
	var sb strings.Builder
	for _, c := range word {
		if cnts[c] > 1 {
			sb.WriteString(")")
		} else {
			sb.WriteString("(")
		}
	}
	return sb.String()
}
