package piscine

func IsLower(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] < 'a' || s[i] > 'z' {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(IsLower("hello"))
// 	fmt.Println(IsLower("hello!"))
// }
