package piscine

// func IsNumeric(s string) bool {
// 	if len(s) == 0 {
// 		return false
// 	}
// 	for i := 0; i < len(s); i++ {
// 		if s[i] < '0' || s[i] > '9' {
// 			return false
// 		}
// 	}
// 	return true
// }

func CountIf(f func(string) bool, tab []string) int {
	c := 0
	for i := range tab {
		if f(tab[i]) == true {
			c++
		}
	}
	return c
}

// func main() {
// 	tab1 := []string{"Hello", "how", "are", "you"}
// 	tab2 := []string{"This", "1", "is", "4", "you"}
// 	answer1 := CountIf(IsNumeric, tab1)
// 	answer2 := CountIf(IsNumeric, tab2)
// 	fmt.Println(answer1)
// 	fmt.Println(answer2)
// }
