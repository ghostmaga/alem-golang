package piscine

func Join(strs []string, sep string) string {
	l := len(strs)
	var r string
	c := l
	for i := range strs {
		r += strs[i]
		if c == 1 {
			break
		}
		if i != l {
			r += sep
			c--
		}

	}
	return r
}

// func main() {
// 	toConcat := []string{"Hello!", " How", " are", " you?"}
// 	fmt.Println(Join(toConcat, ":"))
// }
