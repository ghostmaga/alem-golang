package piscine

func f(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	}
	return -1
}

func IsSorted(f func(a, b int) int, a []int) bool {
	// count := 0
	if len(a) < 2 {
		return true
	}
	if a[0] <= a[1] {
		for i := 0; i < len(a)-1; i++ {
			if f(a[i], a[i+1]) > 0 {
				return false
			}
		}
	}
	if a[0] >= a[1] {
		for i := 0; i < len(a)-1; i++ {
			if f(a[i], a[i+1]) < 0 {
				return false
			}
		}
	}
	return true
}

// func main() {
// 	// a1 := []int{0, 1, 2, 3, 4, 5}
// 	a2 := []int{9, 8}

// 	result1 := IsSorted(f, []int{807254, 750650, 618809, 524808, 347288, 220340, -351866, -995517})
// 	result2 := IsSorted(f, a2)

// 	fmt.Println(result1)
// 	fmt.Println(result2)
// }
