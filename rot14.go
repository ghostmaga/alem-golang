package piscine

func Rot14(s string) string {
	var str string
	for _, i := range s {
		if i >= 'A' && i <= 'Z' {
			c := i + 14
			if c > 90 {
				c = c - 26
			}
			str += string(c)
		} else if i >= 'a' && i <= 'z' {
			c := i + 14
			if c > 122 {
				c = c - 26
			}
			str += string(c)
		} else {
			str = str + string(i)
		}
	}
	return str
}

// func main() {
// 	result := Rot14("Hello! How are You?")

// 	for _, r := range result {
// 		z01.PrintRune(r)
// 	}
// 	z01.PrintRune('\n')
// }
