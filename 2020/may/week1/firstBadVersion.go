package may2020

import "fmt"

//FirstBadVersion ...
func FirstBadVersion(n int) int {
	start, end, mid := 0, n, n/2
	for {
		fmt.Println(start, mid, end)
		if isBadVersion(mid) {
			if !isBadVersion(mid - 1) {
				break
			}
			end = mid
			mid = (start + mid) / 2
		} else {
			if isBadVersion(mid + 1) {
				mid++
				break
			}
			start = mid
			mid = (end + mid) / 2

		}
	}
	return mid
}

func isBadVersion(n int) bool {
	if n >= 5 {
		return true
	}
	return false
}
