package piscine

import "github.com/01-edu/z01"

func intables(base string, s string) int {
	j := 0
	for i := 0; i < len(base); i++ {
		if base[i] == s[j] {
			j++
		} else if base[i] != s[j] {
			j = 0
		}
		if j == len(s) {
			return i + 1 - j
		}
	}
	return -1
}

// func SplitWhiteSpaces(str string) []string {
// 	strarr := []rune(str)
// 	newarr := ""
// 	index := 0
// 	check := false
// 	size := 0
// 	length := 0
// 	for range str {
// 		size++
// 	}
// 	for i, s := range strarr {
// 		if s != '\n' && s != '\t' && s != ' ' && !check {
// 			index = i
// 			check = true
// 		}
// 		if (s == '\n' || s == '\t' || s == ' ') && check {
// 			newarr += string(strarr[index:i])
// 			newarr += " "
// 			check = false
// 			length++
// 		}
// 		if i == size-1 {
// 			newarr += string(strarr[index : i+1])
// 			newarr += " "
// 			length++
// 		}
// 	}
// 	arr := make([]string, length)
// 	for i := 0; i < length; i++ {
// 		arr[i] = newarr[0:intables(newarr, " ")]
// 		newarr = newarr[intables(newarr, " ")+1:]
// 	}
// 	return arr
// }

func PrintWordsTables(a []string) {
	for _, i := range a {
		for j := range i {
			z01.PrintRune(rune(i[j]))
		}
		z01.PrintRune('\n')
	}
}

// func main() {
// 	a := SplitWhiteSpaces("Hello how are you?")
// 	PrintWordsTables(a)
// }
