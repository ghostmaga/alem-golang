package piscine

func Max(a []int) int {
	max := 0
	for i := range a {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}

// func main() {
// 	a := []int{23, 123, 1, 11, 55, 93}
// 	max := Max(a)
// 	fmt.Println(max)
// }
