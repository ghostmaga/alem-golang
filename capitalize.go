package piscine

func Capitalize(s string) string {
	var done bool = true
	s = tolow(s)
	r := []byte(s)
	for i := range r {
		if r[i] >= '0' && r[i] <= '9' {
			done = false
		}
		if (r[i] > 96 && r[i] < 123) && done == true {
			r[i] = r[i] - 32
			done = false
		} else if s[i] < '0' || s[i] > '9' && s[i] < 'A' || s[i] > 'Z' && s[i] < 'a' || s[i] > 'z' {
			done = true
		}
	}
	return string(r)
}

func tolow(s string) string {
	r := []byte(s)
	for i := 0; i < len(s); i++ {
		if r[i] >= 'A' && r[i] <= 'Z' {
			r[i] = r[i] + 32
		}
	}
	return string(r)
}

// func main() {
// 	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
// }
