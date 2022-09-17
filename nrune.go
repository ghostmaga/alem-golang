package piscine

func NRune(s string, n int) rune {
	r := []rune(s)
	l := len(r)
	// ret := rune("")
	if n > l || n < 1 {
		return 0
	}
	return r[n-1]
}
