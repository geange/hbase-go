package hbase

import thrift2 "github.com/geange/hbase-go/thrift/v2"

func toThrift2ColumnValueList(vars []*ColumnValue) []*thrift2.TColumnValue {
	if len(vars) == 0 {
		return nil
	}
	list := make([]*thrift2.TColumnValue, 0, len(vars))
	for _, v := range vars {
		list = append(list, v.toThrift2())
	}
	return list
}
