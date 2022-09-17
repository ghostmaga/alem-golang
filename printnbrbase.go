package piscine

import (
	"github.com/01-edu/z01"
)

func notvalid(base string) bool {
	if len(base) < 2 {
		return true
	}
	for i := range base {
		if base[i] == '+' || base[i] == '-' {
			return true
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return true
			}
		}
	}
	return false
}

func PrintNbrBase(nbr int, base string) {
	if notvalid(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
	} else {
		n := len(base)
		var str string
		i := 0
		if nbr < 0 {
			z01.PrintRune('-')
			if nbr == -9223372036854775808 {
				i = nbr % n
				i *= -1
				nbr /= n
				str = str + string(base[i])
			}
			nbr *= -1
		}
		for nbr >= n {
			i = nbr % n
			nbr = nbr / n
			str = str + string(base[i])
		}
		str = str + string(base[nbr])

		strrev := []byte(str)
		l := len(strrev) - 1
		j := 0
		for l >= 0 {
			strrev[j] = str[l]
			j++
			l--
		}
		str = string(strrev)

		for i := range str {
			z01.PrintRune(rune(str[i]))
		}
	}
}

// func main() {
// 	PrintNbrBase(125, "0123456789")
// 	z01.PrintRune('\n')
// 	PrintNbrBase(-125, "01")
// 	z01.PrintRune('\n')
// 	PrintNbrBase(125, "0123456789ABCDEF")
// 	z01.PrintRune('\n')
// 	PrintNbrBase(-125, "choumi")
// 	z01.PrintRune('\n')
// 	PrintNbrBase(125, "aa")
// 	z01.PrintRune('\n')
// }
