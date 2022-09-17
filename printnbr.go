package piscine

import "github.com/01-edu/z01"

// func PrintNbr(n int) {
// 	if n == 0 {
// 		z01.PrintRune('0')
// 		return
// 	}
// 	if n < 0 {
// 		z01.PrintRune('-')
// 		n *= -1
// 	} else if n > 0 {
// 		var array []int
// 		eachValue := 0
// 		for n != 0 {
// 			eachValue = n % 10
// 			n /= 10
// 			array = append(array, eachValue)
// 		}

// 		for i := len(array) - 1; i >= 0; i-- {
// 			z01.PrintRune(rune(array[i] + 48))
// 		}

// 	}
// }

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		if n == -9223372036854775808 {
			z01.PrintRune('9')
			n = n % 1000000000000000000
		}
		n *= -1
	}

	if n > 0 {
		var array []int
		eachValue := 0
		for n != 0 {
			eachValue = n % 10
			n /= 10
			array = append(array, eachValue)
		}

		for i := len(array) - 1; i >= 0; i-- {
			z01.PrintRune(rune(array[i] + 48))
		}

	}
}

// func main() {
// 	PrintNbr(-9223372036854775808)
// 	// PrintNbr(0)
// 	// PrintNbr(123)
// 	z01.PrintRune('\n')
// }
