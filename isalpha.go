package piscine

func IsAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' && s[i] < 'A' || s[i] > 'Z' && s[i] < 'a' || s[i] > 'z' {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(IsAlpha(" "))
// 	fmt.Println(IsAlpha("HelloHowareyou"))
// 	fmt.Println(IsAlpha("What's this 4?"))
// 	fmt.Println(IsAlpha(""))
// }
