package hbase

import (
	"context"
	"sync"
)

type TableName struct {
	Namespace []byte
	TableName []byte
}

type ThriftWithPool interface {
	EnableTable(ctx context.Context, tableName *TableName) (err error)
	DisableTable(ctx context.Context, tableName *TableName) (err error)
	IsTableEnabled(ctx context.Context, tableName *TableName) (r bool, err error)
	IsTableDisabled(ctx context.Context, tableName *TableName) (r bool, err error)
}

type client struct {
	sync.Mutex
}
