package may2020

import "fmt"

// Count Square Submatrices with All Ones
//   Given a m * n matrix of ones and zeros,
//   return how many square submatrices have all ones.

//CountSquare ...  시간초과..
func CountSquare(matrix [][]int) int {
	xl := len(matrix)
	yl := len(matrix[0])
	var n int
	if xl > yl {
		n = yl
	} else {
		n = xl
	}
	count := 0
	for n > 0 {
		for i := 0; i < xl; i++ {
			if xl-i < n {
				break
			}
			for j := 0; j < yl; j++ {
				if yl-j < n {
					break
				}
				fmt.Println(i, j, n)
				if isSquare(matrix, i, j, n) {
					count++
				}
			}
		}
		n--
	}
	return count
}

func isSquare(m [][]int, i, j, n int) bool {
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			if m[i+x][j+y] == 0 {
				return false
			}
		}
	}
	return true
}
