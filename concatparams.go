package piscine

func ConcatParams(args []string) string {
	l := len(args)
	var a string
	for i := 0; i < l; i++ {
		a += args[i]
		if i != l-1 {
			a += "\n"
		}
	}
	return a
}

// func main() {
// 	test := []string{"Hello", "how", "are", "you?"}
// 	fmt.Println(ConcatParams(test))
// }
