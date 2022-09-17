package piscine

func ToUpper(s string) string {
	r := []byte(s)
	for i := 0; i < len(s); i++ {
		if r[i] > 96 && r[i] < 123 {
			r[i] = r[i] - 32
		}
	}
	return string(r)
}

// func main() {
// 	fmt.Println(ToUpper("Hello! How are you?"))
// }
