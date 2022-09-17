package piscine

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
	return arr[2]
}

// func main() {
// 	middle := Abort(8588436943631415843, 2089606424447550122, -5259421004519017324, 5612258996908273977, -9068070812324729669)
// 	fmt.Println(middle)
// }
