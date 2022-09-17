package piscine

func TrimAtoi(s string) int {
	n := 0
	sign := 1
	index := 0
	for j := 0; j < len(s); j++ {
		if s[j] >= '0' && s[j] <= '9' {
			index = j
			break
		}
	}
	for i := range s {
		if s[i] == '-' {
			if i < index {
				sign = -1
			}
		} else if s[i] >= '0' && s[i] <= '9' {
			n = n*10 + (int(s[i]) - '0')
		}
	}
	// fmt.Printf("index of num %v ", index)
	return n * sign
}

// func main() {
// 	fmt.Println(TrimAtoi("12345"))
// 	fmt.Println(TrimAtoi("str123ing45"))
// 	fmt.Println(TrimAtoi("012 345"))
// 	fmt.Println(TrimAtoi("Hello World!"))
// 	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
// 	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
// 	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
// 	fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
// }
