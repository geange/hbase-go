package hbase

import (
	"context"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/geange/hbase-go/thrift/hbase"
	"net"
	"strconv"
	"time"
)

type rawClient struct {
	*hbase.THBaseServiceClient
	transport thrift.TTransport
	host      string
	port      int
}

func (c *rawClient) Open() error {
	return c.transport.Open()
}

func (c *rawClient) Close(ctx context.Context) error {
	if err := c.transport.Flush(ctx); err != nil {
		return err
	}
	return c.transport.Close()
}

type RawClientOption struct {
	Host       string
	Port       int
	BufferSize int
}

func NewRawClient(ctx context.Context, option RawClientOption) (rc RawClient, err error) {
	hostPort := net.JoinHostPort(option.Host, strconv.Itoa(option.Port))

	var socket *thrift.TSocket
	deadline, ok := ctx.Deadline()
	if ok {
		socket, err = thrift.NewTSocketTimeout(hostPort, deadline.Sub(time.Now()))
	} else {
		socket, err = thrift.NewTSocketTimeout(hostPort, 0)
	}
	if err != nil {
		return nil, err
	}

	var transport thrift.TTransport
	if option.BufferSize == 0 {
		transport, err = thrift.NewTTransportFactory().GetTransport(socket)
	} else {
		transport, err = thrift.NewTBufferedTransportFactory(option.BufferSize).GetTransport(socket)
	}
	if err != nil {
		return nil, err
	}

	client := hbase.NewTHBaseServiceClientFactory(transport, thrift.NewTBinaryProtocolFactoryDefault())

	return &rawClient{
		THBaseServiceClient: client,
		host:                option.Host,
		port:                option.Port,
		transport:           transport,
	}, nil
}
