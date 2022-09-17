package piscine

func BasicJoin(elems []string) string {
	var r string
	for i := 0; i < len(elems); i++ {
		r += elems[i]
	}
	return r
}

// func main() {
// 	elems := []string{"Hello!", " How", " are", " you?"}
// 	fmt.Println(BasicJoin(elems))
// }
