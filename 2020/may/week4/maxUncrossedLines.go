package may2020

//MaxUncrossedLines ... 어려움 모든경우의수 다해야하나...
func MaxUncrossedLines(A, B []int) int {
	max := 0
	for i := 0; i < len(A); i++ {
		temp := getUncrossedLine(i, A, B)
		if max < temp {
			max = temp
		}
	}
	return max
}

func getUncrossedLine(start int, A, B []int) int {
	link := [][]int{}
	i := start
	for {
		for j := 0; j < len(B); j++ {
			if A[i] == B[j] {
				flag := true
				if link != nil {
					for _, line := range link {
						if (line[0] < i && line[1] > j) || (line[0] > i && line[1] < j) || line[1] == j {
							flag = false
							break
						}
					}
				}
				if flag {
					link = append(link, []int{i, j})
					break
				}

			}
		}
		i++
		if i >= len(A) {
			i = 0
		}
		if i == start {
			break
		}
	}
	return len(link)
}
