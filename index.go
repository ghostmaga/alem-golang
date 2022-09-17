package piscine

func Index(s string, toFind string) int {
	r1 := []rune(s)
	r2 := []rune(toFind)
	c := 0
	if len(r2) == 0 {
		return 0
	}
	v := 0
	if len(r1) == len(r2) {
		for j := range r1 {
			if r1[j] == r2[j] {
				v++
			}
		}
		if v == len(r1) {
			return 0
		}
	}
	for i := range r1 {
		if c == len(r2) {
			return i - len(r2)
		} else if r1[i] == r2[c] {
			c++
		} else {
			c = 0
		}
	}
	return -1
	// c := 0
	// if len(toFind) < 1 {
	// 	return 0
	// }
	// for i := 0; i < len(s); i++ {
	// 	if s[i] == toFind[0] {
	// 		if len(toFind) == 1 {
	// 			return i
	// 		} else {
	// 			for j := i; j < i+len(toFind); j++ {
	// 				if s[j] == toFind[c] {
	// 					c++
	// 				}
	// 			}
	// 			if c == len(toFind) {
	// 				return i
	// 			}
	// 		}
	// 	}
	// }
	// return -1
}

// func main() {
// 	fmt.Println(Index("nHvXD'd6@jsnV", "'"))
// 	fmt.Println(Index("", ""))
// 	fmt.Println(Index("Ola!", "Ola!"))
// }
