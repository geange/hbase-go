package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/geange/hbase-go/thrift/hbase"
)

func main() {
	sock, err := thrift.NewTSocket("127.0.0.1:9090")
	if err != nil {
		panic(err)
	}

	factory := thrift.NewTTransportFactory()
	transport, err := factory.GetTransport(sock)
	if err != nil {
		panic(err)
	}

	client := hbase.NewTHBaseServiceClientFactory(transport, thrift.NewTBinaryProtocolFactoryDefault())

	if err := transport.Open(); err != nil {
		panic(err)
	}
	defer transport.Close()

	tables, err := client.ListNamespaceDescriptors(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(tables)
}
