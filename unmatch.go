package piscine

func Unmatch(a []int) int {
	c := 0
	for i := 0; i < len(a); i++ {
		for _, j := range a {
			if a[i] == j {
				c++
			}
		}
		if c%2 == 1 {
			return a[i]
		}
	}
	return -1
}

// func main() {
// 	a := []int{1, 3, 1, 3, 4, 20, 4}
// 	unmatch := Unmatch(a)
// 	fmt.Println(unmatch)
// }
