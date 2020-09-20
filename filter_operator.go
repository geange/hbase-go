package hbase

import "fmt"

func AND(f1, f2 Filter) Filter {
	return stringFilter(fmt.Sprintf("(%s) AND (%s)", f1.FilterExpress(), f2.FilterExpress()))
}

func OR(f1, f2 Filter) Filter {
	return stringFilter(fmt.Sprintf("(%s) OR (%s)", f1.FilterExpress(), f2.FilterExpress()))
}

func SKIP(f Filter) Filter {
	return stringFilter(fmt.Sprintf("SKIP %s", f.FilterExpress()))
}

func WHILE(f Filter) Filter {
	return stringFilter(fmt.Sprintf("WHILE %s", f.FilterExpress()))
}
