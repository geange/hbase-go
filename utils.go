package hbase

import "fmt"

func stringSliceToString(strs []string) string {
	result := make([]byte, 0, 20)
	for i, v := range strs {
		result = append(result, []byte(v)...)
		if i != len(strs)-1 {
			result = append(result, ',')
		}
	}
	return string(result)
}

func int64SliceTostring(nums []int64) string {
	result := make([]byte, 0, 20)
	for i, v := range nums {
		result = append(result, []byte(fmt.Sprint(v))...)
		if i != len(nums)-1 {
			result = append(result, ',')
		}
	}
	return string(result)
}
