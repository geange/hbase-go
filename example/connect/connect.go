package main

import (
	"context"
	"fmt"
	"github.com/geange/hbase-go/thrift/v2"
	"time"

	hb "github.com/geange/hbase-go"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := hb.NewRawClient(ctx, hb.RawClientOption{
		Host:       "172.23.58.228",
		Port:       9090,
		BufferSize: 8192,
	})
	if err != nil {
		panic(err)
	}

	if err := client.Open(); err != nil {
		panic(err)
	}

	table := []byte("demo")

	err = client.Put(context.Background(), table, &hbase.TPut{
		Row: []byte("abc"),
		ColumnValues: []*hbase.TColumnValue{&hbase.TColumnValue{
			Family:    []byte("d"),
			Qualifier: []byte("001"),
			Value:     []byte("1234567890"),
		}},
	})
	if err != nil {
		panic(err)
	}

	result, err := client.Get(context.Background(), table, &hbase.TGet{
		Row: []byte("abc"),
		Columns: []*hbase.TColumn{&hbase.TColumn{
			Family:    []byte("d"),
			Qualifier: []byte("001"),
		}},
		FilterString: []byte("PrefixFilter ('ab')"),
	})
	if err != nil {
		panic(err)
	}

	for _, v := range result.ColumnValues {
		fmt.Println(string(v.Family), string(v.Qualifier), string(v.Value), *v.Timestamp)
	}

	tables, err := client.ListNamespaceDescriptors(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(tables)
}
