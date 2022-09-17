package piscine

// import "fmt"

func FindNextPrime(nb int) int {
	if nb < 2 {
		return 2
	}
	if primecheck(nb) {
		return nb
	} else {
		for !primecheck(nb) {
			nb++
		}
	}
	return nb
}

func primecheck(nb int) bool {
	c := 0
	if nb < 133 {
		c = nb
	} else {
		c = nb / 3
	}
	for i := 2; i < c; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(FindNextPrime(-5))
// 	fmt.Println(FindNextPrime(134))
// }
