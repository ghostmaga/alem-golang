package piscine

func IterativeFactorial(nb int) int {
	var result int = 1
	if nb < 0 || nb > 25 {
		return 0
	} else if nb > 1 {
		for i := 1; i <= nb; i++ {
			result = result * i
		}
	} else {
		return 1
	}
	return result
}
