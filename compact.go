package piscine

const N = 6

// func remove(s []string, i int) []string {
// 	return append(s[:i], s[(i+1):]...)
// }

func Compact(ptr *[]string) int {
	var str []string
	for _, val := range *ptr {
		if val != "" {
			str = append(str, val)
		}
	}
	*ptr = str
	return len(str)
	// a := make([]string, N)
	// for i, val := range *ptr {
	// 	a[i] = val
	// }
	// for i, val := range a {
	// 	if val == nil {
	// 		a = remove(a, i)
	// 	}
	// }
	// *ptr = a
	// return len(a)
}

// func main() {
// 	a := make([]string, N)
// 	a[0] = "a"
// 	a[2] = "b"
// 	a[4] = "c"

// 	for _, v := range a {
// 		fmt.Println(v)
// 	}

// 	fmt.Println("Size after compacting:", Compact(&a))

// 	for _, v := range a {
// 		fmt.Println(v)
// 	}
// }
