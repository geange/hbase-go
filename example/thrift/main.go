package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/geange/hbase-go/thrift/v1"
	"net"
	"time"
)

func main() {
	hostPort := net.JoinHostPort("0.0.0.0", "9090")
	socket, err := thrift.NewTSocketTimeout(hostPort, 10*time.Second)
	transport := thrift.NewTTransportFactory().GetTransport(socket)
	client := hbase.NewHbaseClientFactory(transport, thrift.NewTBinaryProtocolFactoryDefault())

	transport.Open()

	tables, err := client.GetTableNames()
	if err != nil {
		panic(err)
	}

	for _, v := range tables {
		fmt.Println(string(v))
	}

	if transport.IsOpen() {
		transport.Close()
	}

	transport.Open()

	tables, err = client.GetTableNames()
	if err != nil {
		panic(err)
	}

	for _, v := range tables {
		fmt.Println(string(v))
	}
}
