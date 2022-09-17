package piscine

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	a := 0
	j := 0
	k := 1
	if s[0] == '+' {
		j++
	}
	if s[0] == '-' {
		k = -1
		j++
	}
	for i := j; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			a = 0
			break
		}
		a = a*10 + (int(s[i]) - 48)
	}
	return a * k
}
