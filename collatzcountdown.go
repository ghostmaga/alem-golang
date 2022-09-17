package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	c := 0
	for start != 1 {
		if start%2 == 0 {
			start /= 2
			c++
		} else if start%2 == 1 {
			start = start*3 + 1
			c++
		}
	}
	return c
}

// func main() {
// 	steps := CollatzCountdown(0)
// 	fmt.Println(steps)
// }
