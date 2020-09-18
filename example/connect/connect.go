package main

import (
	"context"
	"fmt"
	"time"

	hb "github.com/geange/hbase-go"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := hb.NewRawClient(ctx, hb.RawClientOption{
		Host:       "127.0.0.1",
		Port:       9090,
		BufferSize: 8192,
	})
	if err != nil {
		panic(err)
	}

	if err := client.Open(); err != nil {
		panic(err)
	}

	tables, err := client.ListNamespaceDescriptors(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(tables)
}
