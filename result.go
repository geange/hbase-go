package hbase

type Result struct {
	Row []byte
}

type ColumnValue struct {
	Family    []byte
	Qualifier []byte
	Value     []byte
	Timestamp *int64
}
