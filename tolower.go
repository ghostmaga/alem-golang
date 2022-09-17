package piscine

func ToLower(s string) string {
	r := []byte(s)
	for i := 0; i < len(s); i++ {
		if r[i] >= 'A' && r[i] <= 'Z' {
			r[i] = r[i] + 32
		}
	}
	return string(r)
}

// func main() {
// 	fmt.Println(ToLower("HeAlo! How are yZou?"))
// }
