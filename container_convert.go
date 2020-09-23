package hbase

import thrift2 "github.com/geange/hbase-go/thrift/v2"

func toThrift2TGetList(vars []*Get) []*thrift2.TGet {
	if len(vars) == 0 {
		return nil
	}
	list := make([]*thrift2.TGet, 0, len(vars))
	for _, v := range vars {
		list = append(list, toThrift2TGet(v))
	}
	return list
}

func toThrift2ColumnValueList(vars []*ColumnValue) []*thrift2.TColumnValue {
	if len(vars) == 0 {
		return nil
	}
	list := make([]*thrift2.TColumnValue, 0, len(vars))
	for _, v := range vars {
		list = append(list, toThrift2TColumnValue(v))
	}
	return list
}

func toThrift2ColumnList(vars []*Column) []*thrift2.TColumn {
	if len(vars) == 0 {
		return nil
	}
	list := make([]*thrift2.TColumn, 0, len(vars))
	for _, v := range vars {
		list = append(list, toThrift2TColumn(v))
	}
	return list
}

func toThrift2MutationList(vars []*Mutation) []*thrift2.TMutation {
	if len(vars) == 0 {
		return nil
	}
	list := make([]*thrift2.TMutation, 0, len(vars))
	for _, v := range vars {
		list = append(list, toThrift2TMutation(v))
	}
	return list
}

func toThriftTTimeRangeMap(ptr map[string]*TimeRange) map[string]*thrift2.TTimeRange {
	if ptr == nil {
		return nil
	}

	tmap := make(map[string]*thrift2.TTimeRange)
	for k, v := range ptr {
		tmap[k] = toThrift2TTimeRange(v)
	}
	return tmap
}

func toThrift2ColumnFamilyDescriptorList(vars []*ColumnFamilyDescriptor) []*thrift2.TColumnFamilyDescriptor {
	if len(vars) == 0 {
		return nil
	}
	list := make([]*thrift2.TColumnFamilyDescriptor, 0, len(vars))
	for _, v := range vars {
		list = append(list, toThrift2TColumnFamilyDescriptor(v))
	}
	return list
}

func toThrift2TTableNameList(vars []*TableName) []*thrift2.TTableName {
	if len(vars) == 0 {
		return nil
	}
	list := make([]*thrift2.TTableName, 0, len(vars))
	for _, v := range vars {
		list = append(list, toThrift2TTableName(v))
	}
	return list
}

func toThrift2TPutList(vars []*Put) []*thrift2.TPut {
	if len(vars) == 0 {
		return nil
	}
	list := make([]*thrift2.TPut, 0, len(vars))
	for _, v := range vars {
		list = append(list, toThrift2TPut(v))
	}
	return list
}

func toThrift2TDeleteList(vars []*Delete) []*thrift2.TDelete {
	if len(vars) == 0 {
		return nil
	}
	list := make([]*thrift2.TDelete, 0, len(vars))
	for _, v := range vars {
		list = append(list, toThrift2TDelete(v))
	}
	return list
}

func toThrift2TColumnIncrementList(vars []*ColumnIncrement) []*thrift2.TColumnIncrement {
	if len(vars) == 0 {
		return nil
	}
	list := make([]*thrift2.TColumnIncrement, 0, len(vars))
	for _, v := range vars {
		list = append(list, toThrift2TColumnIncrement(v))
	}
	return list
}

func fromThrift2TableDescriptorList(vars []*thrift2.TTableDescriptor) []*TableDescriptor {
	if len(vars) == 0 {
		return nil
	}
	list := make([]*TableDescriptor, 0, len(vars))
	for _, v := range vars {
		list = append(list, fromThrift2TableDescriptor(v))
	}
	return list
}

func fromThrift2NamespaceDescriptorList(vars []*thrift2.TNamespaceDescriptor) []*NamespaceDescriptor {
	if len(vars) == 0 {
		return nil
	}
	list := make([]*NamespaceDescriptor, 0, len(vars))
	for _, v := range vars {
		list = append(list, fromThrift2TNamespaceDescriptor(v))
	}
	return list
}

func fromThrift2ColumnFamilyDescriptorList(vars []*thrift2.TColumnFamilyDescriptor) []*ColumnFamilyDescriptor {
	if len(vars) == 0 {
		return nil
	}
	list := make([]*ColumnFamilyDescriptor, 0, len(vars))
	for _, v := range vars {
		list = append(list, fromThrift2ColumnFamilyDescriptor(v))
	}
	return list
}

func fromThrift2TTableNameList(vars []*thrift2.TTableName) []*TableName {
	if len(vars) == 0 {
		return nil
	}
	list := make([]*TableName, 0, len(vars))
	for _, v := range vars {
		list = append(list, fromThrift2TableName(v))
	}
	return list
}

func fromThrift2HRegionLocationList(vars []*thrift2.THRegionLocation) []*HRegionLocation {
	if len(vars) == 0 {
		return nil
	}
	list := make([]*HRegionLocation, 0, len(vars))
	for _, v := range vars {
		list = append(list, fromThrift2RegionLocation(v))
	}
	return list
}

func fromThrift2ColumnList(vars []*thrift2.TColumn) []*Column {
	if len(vars) == 0 {
		return nil
	}
	list := make([]*Column, 0, len(vars))
	for _, v := range vars {
		list = append(list, fromThrift2Column(v))
	}
	return list
}

func fromThrift2ColumnValueList(vars []*thrift2.TColumnValue) []*ColumnValue {
	if len(vars) == 0 {
		return nil
	}
	list := make([]*ColumnValue, 0, len(vars))
	for _, v := range vars {
		list = append(list, fromThrift2ColumnValue(v))
	}
	return list
}

func fromThrift2ResultList(vars []*thrift2.TResult) []*Result {
	if len(vars) == 0 {
		return nil
	}
	list := make([]*Result, 0, len(vars))
	for _, v := range vars {
		list = append(list, fromThrift2Result(v))
	}
	return list
}

func fromThrift2DeleteList(vars []*thrift2.TDelete) []*Delete {
	if len(vars) == 0 {
		return nil
	}
	list := make([]*Delete, 0, len(vars))
	for _, v := range vars {
		list = append(list, fromThrift2Delete(v))
	}
	return list
}
