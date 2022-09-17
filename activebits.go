package piscine

func find(num int) int {
	if num <= 0 {
		return 0
	} else {
		return find(num/2) + num%2
	}
}

func ActiveBits(n int) int {
	c := find(n)
	return c
}

// func main() {
// 	fmt.Println(ActiveBits(4863))
// }
