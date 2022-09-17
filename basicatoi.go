package piscine

func BasicAtoi(s string) int {
	n := 0
	length := 0
	for c := range s {
		length++
		if s[c] == 0 {
			length++
		}
	}
	for i := 0; i < length; i++ {
		n = n*10 + (int(s[i]) - 48)
	}
	return n
}
