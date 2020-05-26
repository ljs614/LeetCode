package main

import (
	"fmt"

	may2020 "github.com/ljs614/leetcode/2020/may/week4"
)

func main() {
	test()
}

func test() {
	a := []int{1, 3, 7, 1, 7, 5}
	b := []int{1, 9, 2, 5, 1}
	result := may2020.MaxUncrossedLines(a, b)
	if result != 2 {
		fmt.Printf("incorrect. you: %d, expected: 2\n", result)
	}
}
