package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			t := j + 1
			for s := i; s <= '9'; s++ {
				for t <= '9' {
					if j == '8' && i == '9' {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(s)
						z01.PrintRune(t)
					} else {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(s)
						z01.PrintRune(t)
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
					t++
				}
				t = '0'
			}
		}
	}
	z01.PrintRune('\n')
}
