package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
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

		isDone := false

		for !isDone {
			isDone = true
			i := 0
			for i < len(array)-1 {
				if array[i] > array[i+1] {
					array[i], array[i+1] = array[i+1], array[i]
					isDone = false
				}
				i++
			}
		}

		for i := 0; i < len(array); i++ {
			z01.PrintRune(rune(array[i] + 48))
		}

	}
}
