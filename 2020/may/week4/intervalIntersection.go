package may2020

import "fmt"

//IntervalIntersection ...
func IntervalIntersection(A, B [][]int) [][]int {
	var length int
	aLength := len(A)
	bLength := len(B)
	if aLength == 0 || bLength == 0 {
		return [][]int{}
	}
	if A[aLength-1][1] > B[bLength-1][1] {
		length = A[aLength-1][1]
	} else {
		length = B[bLength-1][1]
	}

	temp := make([]float64, length+1)
	setSection(&temp, A)
	setSection(&temp, B)
	ans := [][]int{}
	fmt.Println(temp)
	for i := 0; i < len(temp); i++ {
		if temp[i] <= float64(i+1) {
			continue
		}
		start, end, v := i, i-1, temp[i]
		for i < len(temp) && temp[i] == v {
			end++
			i++
		}

		if end >= start {
			ans = append(ans, []int{start, end})
			i--
		}
	}
	return ans
}

func setSection(arr *[]float64, a [][]int) {
	for _, val := range a {
		t := val[0] + 1
		for i := val[0]; i <= val[1]; i++ {
			(*arr)[i] += float64(t)
		}
	}
}
