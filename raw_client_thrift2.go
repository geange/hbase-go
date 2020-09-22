package hbase

import (
	"context"
	"github.com/apache/thrift/lib/go/thrift"
	thrift2 "github.com/geange/hbase-go/thrift/v2"
	"net"
	"strconv"
	"time"
)

type rawClientV2 struct {
	*thrift2.THBaseServiceClient
	transport thrift.TTransport
	host      string
	port      int
}

func (c *rawClientV2) Open() error {
	if c.transport.IsOpen() {
		return nil
	}
	return c.transport.Open()
}

func (c *rawClientV2) Close(ctx context.Context) error {
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

func NewRawClientV2(ctx context.Context, option RawClientOption) (rc RawClientV2, err error) {
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

	client := thrift2.NewTHBaseServiceClientFactory(transport, thrift.NewTBinaryProtocolFactoryDefault())

	return &rawClientV2{
		THBaseServiceClient: client,
		host:                option.Host,
		port:                option.Port,
		transport:           transport,
	}, nil
}
