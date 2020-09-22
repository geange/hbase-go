package hbase

import "sync"

type ThriftWithPool interface {
}

type client struct {
	sync.Mutex
}
