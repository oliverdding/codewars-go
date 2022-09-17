package kata

func FindMissingLetter(chars []rune) rune {
	last := int(chars[0])
	for i := 1; i < len(chars); i++ {
		if last+i != int(chars[i]) {
			return rune(last + i)
		}
	}
	return 'X'
}
