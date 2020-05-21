package medium

import "fmt"

//LongestPalindrome ...
func LongestPalindrome(s string) string {
	ans := getPalindrome(s)
	ans2 := getPalindrome(reverse(s))
	if len(ans) > len(ans2) {
		return ans
	}
	return ans2

}

func getPalindrome(s string) string {
	l := len(s)
	ans := ""
	for i := 0; i < l; i++ {
		index, end := i, l
		for j := l - 1; j >= 0; j-- {
			if index < j {
				if s[index] == s[j] {
					if end == l {
						end = j
					}
					index++
				} else {
					if end != l {
						end = l
						index = i
					}
				}
			} else {
				if end == l {
					end = j
				}
				if len(ans) < end-i+1 {
					ans = s[i : end+1]
					fmt.Println(ans)
				}
				break
			}
		}
	}
	return ans
}

func reverse(str string) string {
	result := ""
	for i := len(str) - 1; i >= 0; i-- {
		result += string(str[i])
	}
	return result
}
