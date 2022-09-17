package piscine

// func Printint(n int) {
// 	if n == 0 {
// 		z01.PrintRune('0')
// 		return
// 	}
// 	if n < 0 {
// 		z01.PrintRune('-')
// 		n *= -1
// 	}
// 	if n > 0 {
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

func ForEach(f func(int), a []int) {
	for i := range a {
		f(a[i])
	}
	// z01.PrintRune('\n')
}

// func main() {
// 	a := []int{1, 2, 3, 4, 5, 6}
// 	ForEach(Printint, a)
// }
