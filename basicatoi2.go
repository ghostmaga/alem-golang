package piscine

func BasicAtoi2(s string) int {
	a := 0
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			a = 0
			break
		}
		a = a*10 + (int(s[i]) - 48)
	}
	return a
}
