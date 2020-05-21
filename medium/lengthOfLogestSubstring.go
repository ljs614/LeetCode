package main

import "fmt"

func main() {
	fmt.Println("d d"[1])
	lengthOfLongestSubstring("ddss")
}

//베낌
func lengthOfLongestSubstring(s string) int {
	index := [128]int{}
	start, ans := 0, 0
	len := len(s)
	for end := 0; end < len; end++ {
		start = max(index[s[end]], start)
		ans = max(ans, end-start+1)
		index[s[end]] = end + 1
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//그냥 푼거
// func lengthOfLongestSubstring(s string) int {
// 	sArr := strings.Split(s, "")
// 	max := 0
// 	for i := 0; i < len(sArr); i++ {
// 		length := getSubstringLength(i, sArr)
// 		if max < length {
// 			max = length
// 		}
// 		if len(sArr) - i < max {
// 			break
// 		}
// 	}
// 	return max
// }

// func getSubstringLength(i int, arr []string) int {
// 	chars := []string{}
// 	count := 0
// Go:
// 	for j := i; j < len(arr); j++ {
// 		for _, char := range chars {
//             if arr[j] == char {
// 				break Go
// 			}
// 		}
// 		chars = append(chars, arr[j])
// 		count++
// 	}
// 	return count
// }
