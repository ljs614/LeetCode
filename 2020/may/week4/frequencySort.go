package may2020

import (
	"sort"
	"strings"
)

//FrequencySort ...
func FrequencySort(s string) string {
	max := 0
	ans := []byte{}
	arr := make([]myString, 0)
	for len(s) > 0 {
		c := s[0]
		l := len(s)
		s = strings.ReplaceAll(s, string(c), "")
		arr = append(arr, myString{c, l - len(s)})
		i := len(arr) - 1
		if arr[i].count > max {
			max = arr[i].count
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].count > arr[j].count
	})
	for _, val := range arr {
		for i := 0; i < val.count; i++ {
			ans = append(ans, val.value)
		}
	}

	return string(ans)
}

type myString struct {
	value byte
	count int
}
