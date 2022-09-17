package piscine

func IterativePower(nb int, power int) int {
	var result int = 1
	if power < 0 {
		return 0
	} else if power == 1 {
		return nb
	} else if power == 0 {
		return 1
	} else {
		for i := 1; i <= power; i++ {
			result = result * nb
		}
	}
	return result
}
