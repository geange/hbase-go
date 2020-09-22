package hbase

import (
	"fmt"
	"testing"
)

func TestColumnRangeFilter_FilterExpress(t *testing.T) {
	min := "001"
	max := "002"
	filter := columnRangeFilter{
		minColumn: min,
		minFlag:   true,
		maxColumn: max,
		maxFlag:   true,
	}

	express := filter.FilterExpress()
	fmt.Println(express)
}
