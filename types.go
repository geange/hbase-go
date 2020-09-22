package hbase

type RowValue struct {
	Row          []byte
	ColumnValues []*ColumnValue
	Stale        bool
	Partial      bool
}

type ColumnValue struct {
	Family    []byte
	Qualifier []byte
	Value     []byte
	Timestamp *int64
	Tags      []byte
	Type      *int8
}
