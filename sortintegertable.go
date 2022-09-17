package piscine

func SortIntegerTable(table []int) {
	// r := 0
	for i := 0; i < len(table)-1; i++ {
		for j := 0; j < len(table)-1; j++ {
			if table[j] > table[j+1] {
				// r = table[j+1]
				// table[j+1] = table[j]
				// table[j] = r
				table[j+1], table[j] = table[j], table[j+1]
			}
		}
	}
}

// func main() {
// 	s := []int{5, 4, 3, 2, 1, 0}
// 	SortIntegerTable(s)
// 	fmt.Println(s)
// }
