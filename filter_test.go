package hbase

import (
	"fmt"
	"testing"
)

func TestColumnRangeFilter_FilterExpress(t *testing.T) {
	min := "001"
	max := "002"
	filter := columnRangeFilter{
		minColumn: &min,
		maxColumn: &max,
	}

	express := filter.FilterExpress()
	fmt.Println(express)
}
