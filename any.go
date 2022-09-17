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

func Any(f func(string) bool, a []string) bool {
	for i := range a {
		if f(a[i]) == true {
			return true
		}
	}
	return false
}

// func main() {
// 	a1 := []string{"Hello", "how", "are", "you"}
// 	a2 := []string{"This", "is", "4", "you"}

// 	result1 := Any(IsNumeric, a1)
// 	result2 := Any(IsNumeric, a2)

// 	fmt.Println(result1)
// 	fmt.Println(result2)
// }
