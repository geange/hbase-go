# hbase-go

## Installation

go get github.com/geange/hbase-go

## Example

#### create client

```go
import hb "github.com/geange/hbase-go"

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
```

#### Insert a cell

```go
table := []byte("demo")
err = client.Put(context.Background(), table, &hbase.TPut{
    Row: []byte("abc"),
    ColumnValues: []*hbase.TColumnValue{&hbase.TColumnValue{
        Family:    []byte("d"),
        Qualifier: []byte("001"),
        Value:     []byte("1234567890"),
    }},
})
```

#### Get an entire row

```go
table := []byte("demo")

result, err := client.Get(context.Background(), table, &hbase.TGet{
    Row: []byte("abc"),
})
if err != nil {
    panic(err)
}

for _, v := range result.ColumnValues {
    fmt.Println(string(v.Family), string(v.Qualifier), string(v.Value))
}
```

#### Get a special cell

```go
family := []byte("d")
qualifier := []byte("001")
result, err := client.Get(context.Background(), table, &hbase.TGet{
    Row: []byte("abc"),
    Columns: []*hbase.TColumn{&hbase.TColumn{
        Family:    family,
        Qualifier: qualifier,
    }},
})
```

#### Get a specific cell with a filter

```go
result, err := client.Get(context.Background(), table, &hbase.TGet{
    Row: []byte("abc"),
    Columns: []*hbase.TColumn{&hbase.TColumn{
        Family:    []byte("d"),
        Qualifier: []byte("001"),
    }},
    FilterString: []byte("PrefixFilter ('ab')"),
})
```