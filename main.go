package main

import (
	"fmt"

	may2020 "github.com/ljs614/leetcode/2020/may/week4"
)

func main() {
	a := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
	b := [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
	fmt.Println(may2020.IntervalIntersection(a, b))
}
