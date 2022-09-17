package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	a := make([]int, max-min)
	c := 0
	for i := min; i < max; i++ {
		a[c] = i
		c++
	}
	return a
}

// func main() {
// 	fmt.Println(MakeRange(5, 10))
// 	fmt.Println(MakeRange(10, 5))
// }
