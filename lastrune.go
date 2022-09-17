package piscine

func LastRune(s string) rune {
	r := []rune(s)
	l := len(r)
	return r[l-1]
}
