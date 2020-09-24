package hbase

import (
	"context"
	thrift2 "github.com/geange/hbase-go/thrift/v2"
)

var (
	Thrift  = 1
	Thrift2 = 2
)

type client struct {
	version       int
	thrift1Client Thrift1Client
	thrift2Client Thrift2Client
}

func (c *client) Exists(ctx context.Context, table []byte, get *Get) (r bool, err error) {
	switch c.version {
	case Thrift:

	case Thrift2:
		return c.thrift2Client.Exists(ctx, table, toThrift2TGet(get))
	}
	return false, ErrorClientVersion
}

func (c *client) ExistsAll(ctx context.Context, table []byte, gets []*Get) (r []bool, err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		return c.thrift2Client.ExistsAll(ctx, table, toThrift2TGetList(gets))
	}
	return nil, ErrorClientVersion
}

func (c *client) Get(ctx context.Context, table []byte, get *Get) (r *Result, err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		result, err := c.thrift2Client.Get(ctx, table, toThrift2TGet(get))
		if err != nil {
			return nil, err
		}
		return fromThrift2Result(result), nil
	}
	return nil, ErrorClientVersion
}

func (c *client) GetMultiple(ctx context.Context, table []byte, Gets []*Get) (r []*Result, err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		result, err := c.thrift2Client.GetMultiple(ctx, table, toThrift2TGetList(Gets))
		if err != nil {
			return nil, err
		}
		return fromThrift2ResultList(result), nil
	}
	return nil, ErrorClientVersion
}

func (c *client) Put(ctx context.Context, table []byte, put *Put) (err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		return c.thrift2Client.Put(ctx, table, toThrift2TPut(put))
	}
	return ErrorClientVersion
}

func (c *client) CheckAndPut(ctx context.Context, table []byte, row []byte, family []byte, qualifier []byte, value []byte, put *Put) (r bool, err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		return c.thrift2Client.CheckAndPut(ctx, table, row, family, qualifier, value, toThrift2TPut(put))
	}
	return false, ErrorClientVersion
}

func (c *client) PutMultiple(ctx context.Context, table []byte, puts []*Put) (err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		return c.thrift2Client.PutMultiple(ctx, table, toThrift2TPutList(puts))
	}
	return ErrorClientVersion
}

func (c *client) DeleteSingle(ctx context.Context, table []byte, delete *Delete) (err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		return c.thrift2Client.DeleteSingle(ctx, table, toThrift2TDelete(delete))
	}
	return ErrorClientVersion
}

func (c *client) DeleteMultiple(ctx context.Context, table []byte, deletes []*Delete) (r []*Delete, err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		list, err := c.thrift2Client.DeleteMultiple(ctx, table, toThrift2TDeleteList(deletes))
		if err != nil {
			return nil, err
		}
		return fromThrift2DeleteList(list), nil
	}
	return nil, ErrorClientVersion
}

func (c *client) CheckAndDelete(ctx context.Context, table []byte, row []byte, family []byte, qualifier []byte, value []byte, delete *Delete) (r bool, err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		return c.thrift2Client.CheckAndDelete(ctx, table, row, family, qualifier, value, toThrift2TDelete(delete))
	}
	return false, ErrorClientVersion
}

func (c *client) Increment(ctx context.Context, table []byte, increment *Increment) (r *Result, err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		v, err := c.thrift2Client.Increment(ctx, table, toThrift2TIncrement(increment))
		if err != nil {
			return nil, err
		}
		return fromThrift2Result(v), nil
	}
	return nil, ErrorClientVersion
}

func (c *client) Append(ctx context.Context, table []byte, append *Append) (r *Result, err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		v, err := c.thrift2Client.Append(ctx, table, toThrift2TAppend(append))
		if err != nil {
			return nil, err
		}
		return fromThrift2Result(v), nil
	}
	return nil, ErrorClientVersion
}

func (c *client) OpenScanner(ctx context.Context, table []byte, scan *Scan) (r int32, err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		return c.thrift2Client.OpenScanner(ctx, table, toThrift2TScan(scan))
	}
	return 0, ErrorClientVersion
}

func (c *client) GetScannerRows(ctx context.Context, scannerId int32, numRows int32) (r []*Result, err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		list, err := c.thrift2Client.GetScannerRows(ctx, scannerId, numRows)
		if err != nil {
			return nil, err
		}
		return fromThrift2ResultList(list), nil
	}
	return nil, ErrorClientVersion
}

func (c *client) CloseScanner(ctx context.Context, scannerId int32) (err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		return c.thrift2Client.CloseScanner(ctx, scannerId)
	}
	return ErrorClientVersion
}

func (c *client) MutateRow(ctx context.Context, table []byte, trowMutations *RowMutations) (err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		return c.thrift2Client.MutateRow(ctx, table, toThrift2TRowMutations(trowMutations))
	}
	return ErrorClientVersion
}

func (c *client) GetScannerResults(ctx context.Context, table []byte, tscan *Scan, numRows int32) (r []*Result, err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		list, err := c.thrift2Client.GetScannerResults(ctx, table, toThrift2TScan(tscan), numRows)
		if err != nil {
			return nil, err
		}
		return fromThrift2ResultList(list), nil
	}
	return nil, ErrorClientVersion
}

func (c *client) GetRegionLocation(ctx context.Context, table []byte, row []byte, reload bool) (r *HRegionLocation, err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		location, err := c.thrift2Client.GetRegionLocation(ctx, table, row, reload)
		if err != nil {
			return nil, err
		}
		return fromThrift2RegionLocation(location), nil
	}
	return nil, ErrorClientVersion
}

func (c *client) GetAllRegionLocations(ctx context.Context, table []byte) (r []*HRegionLocation, err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		list, err := c.thrift2Client.GetAllRegionLocations(ctx, table)
		if err != nil {
			return nil, err
		}
		return fromThrift2HRegionLocationList(list), nil
	}
	return nil, ErrorClientVersion
}

func (c *client) CheckAndMutate(ctx context.Context, table []byte, row []byte, family []byte, qualifier []byte, compareOp TCompareOp, value []byte, rowMutations *RowMutations) (r bool, err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		return c.thrift2Client.CheckAndMutate(ctx, table, row, family, qualifier, thrift2.TCompareOp(compareOp), value, toThrift2TRowMutations(rowMutations))
	}
	return false, ErrorClientVersion
}

func (c *client) GetTableDescriptor(ctx context.Context, table *TableName) (r *TableDescriptor, err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		desc, err := c.thrift2Client.GetTableDescriptor(ctx, toThrift2TTableName(table))
		if err != nil {
			return nil, err
		}
		return fromThrift2TableDescriptor(desc), nil
	}
	return nil, ErrorClientVersion
}

func (c *client) GetTableDescriptors(ctx context.Context, tables []*TableName) (r []*TableDescriptor, err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		descList, err := c.thrift2Client.GetTableDescriptors(ctx, toThrift2TTableNameList(tables))
		if err != nil {
			return nil, err
		}
		return fromThrift2TableDescriptorList(descList), nil
	}
	return nil, ErrorClientVersion
}

func (c *client) TableExists(ctx context.Context, tableName *TableName) (r bool, err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		return c.thrift2Client.TableExists(ctx, toThrift2TTableName(tableName))
	}
	return false, ErrorClientVersion
}

func (c *client) GetTableDescriptorsByPattern(ctx context.Context, regex string, includeSysTables bool) (r []*TableDescriptor, err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		list, err := c.thrift2Client.GetTableDescriptorsByPattern(ctx, regex, includeSysTables)
		if err != nil {
			return nil, err
		}
		return fromThrift2TableDescriptorList(list), nil
	}
	return nil, ErrorClientVersion
}

func (c *client) GetTableDescriptorsByNamespace(ctx context.Context, name string) (r []*TableDescriptor, err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		list, err := c.thrift2Client.GetTableDescriptorsByNamespace(ctx, name)
		if err != nil {
			return nil, err
		}
		return fromThrift2TableDescriptorList(list), nil
	}
	return nil, ErrorClientVersion
}

func (c *client) GetTableNamesByPattern(ctx context.Context, regex string, includeSysTables bool) (r []*TableName, err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		list, err := c.thrift2Client.GetTableNamesByPattern(ctx, regex, includeSysTables)
		if err != nil {
			return nil, err
		}
		return fromThrift2TTableNameList(list), nil
	}
	return nil, ErrorClientVersion
}

func (c *client) GetTableNamesByNamespace(ctx context.Context, name string) (r []*TableName, err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		list, err := c.thrift2Client.GetTableNamesByNamespace(ctx, name)
		if err != nil {
			return nil, err
		}
		return fromThrift2TTableNameList(list), nil
	}
	return nil, ErrorClientVersion
}

func (c *client) CreateTable(ctx context.Context, desc *TableDescriptor, splitKeys [][]byte) (err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		return c.thrift2Client.CreateTable(ctx, toThrift2TTableDescriptor(desc), splitKeys)
	}
	return ErrorClientVersion
}

func (c *client) DeleteTable(ctx context.Context, tableName *TableName) (err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		return c.thrift2Client.DeleteTable(ctx, toThrift2TTableName(tableName))
	}
	return ErrorClientVersion
}

func (c *client) TruncateTable(ctx context.Context, tableName *TableName, preserveSplits bool) (err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		return c.thrift2Client.TruncateTable(ctx, toThrift2TTableName(tableName), preserveSplits)
	}
	return ErrorClientVersion
}

func (c *client) EnableTable(ctx context.Context, tableName *TableName) (err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		return c.thrift2Client.EnableTable(ctx, toThrift2TTableName(tableName))
	}
	return ErrorClientVersion
}

func (c *client) DisableTable(ctx context.Context, tableName *TableName) (err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		return c.thrift2Client.DisableTable(ctx, toThrift2TTableName(tableName))
	}
	return ErrorClientVersion
}

func (c *client) IsTableEnabled(ctx context.Context, tableName *TableName) (r bool, err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		return c.thrift2Client.IsTableEnabled(ctx, toThrift2TTableName(tableName))
	}
	return false, ErrorClientVersion
}

func (c *client) IsTableDisabled(ctx context.Context, tableName *TableName) (r bool, err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		return c.thrift2Client.IsTableDisabled(ctx, toThrift2TTableName(tableName))
	}
	return false, ErrorClientVersion
}

func (c *client) IsTableAvailable(ctx context.Context, tableName *TableName) (r bool, err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		return c.thrift2Client.IsTableAvailable(ctx, toThrift2TTableName(tableName))
	}
	return false, ErrorClientVersion
}

func (c *client) IsTableAvailableWithSplit(ctx context.Context, tableName *TableName, splitKeys [][]byte) (r bool, err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		return c.thrift2Client.IsTableAvailableWithSplit(ctx, toThrift2TTableName(tableName), splitKeys)
	}
	return false, ErrorClientVersion
}

func (c *client) AddColumnFamily(ctx context.Context, tableName *TableName, column *ColumnFamilyDescriptor) (err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		return c.thrift2Client.AddColumnFamily(ctx, toThrift2TTableName(tableName), toThrift2TColumnFamilyDescriptor(column))
	}
	return ErrorClientVersion
}

func (c *client) DeleteColumnFamily(ctx context.Context, tableName *TableName, column []byte) (err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		return c.thrift2Client.DeleteColumnFamily(ctx, toThrift2TTableName(tableName), column)
	}
	return ErrorClientVersion
}

func (c *client) ModifyColumnFamily(ctx context.Context, tableName *TableName, column *ColumnFamilyDescriptor) (err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		return c.thrift2Client.ModifyColumnFamily(ctx, toThrift2TTableName(tableName), toThrift2TColumnFamilyDescriptor(column))
	}
	return ErrorClientVersion
}

func (c *client) ModifyTable(ctx context.Context, desc *TableDescriptor) (err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		desc := toThrift2TTableDescriptor(desc)
		return c.thrift2Client.ModifyTable(ctx, desc)
	}
	return ErrorClientVersion
}

func (c *client) CreateNamespace(ctx context.Context, namespaceDesc *NamespaceDescriptor) (err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		desc := toThrift2TNamespaceDescriptor(namespaceDesc)
		return c.thrift2Client.CreateNamespace(ctx, desc)
	}
	return ErrorClientVersion
}

func (c *client) ModifyNamespace(ctx context.Context, namespaceDesc *NamespaceDescriptor) (err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		return c.thrift2Client.ModifyNamespace(ctx, toThrift2TNamespaceDescriptor(namespaceDesc))
	}
	return ErrorClientVersion
}

func (c *client) DeleteNamespace(ctx context.Context, name string) (err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		return c.thrift2Client.DeleteNamespace(ctx, name)
	}
	return ErrorClientVersion
}

func (c *client) GetNamespaceDescriptor(ctx context.Context, name string) (r *NamespaceDescriptor, err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		desc, err := c.thrift2Client.GetNamespaceDescriptor(ctx, name)
		if err != nil {
			return nil, err
		}
		return fromThrift2TNamespaceDescriptor(desc), nil
	}
	return nil, ErrorClientVersion
}

func (c *client) ListNamespaceDescriptors(ctx context.Context) (r []*NamespaceDescriptor, err error) {
	switch c.version {
	case Thrift:
	case Thrift2:
		list, err := c.thrift2Client.ListNamespaceDescriptors(ctx)
		if err != nil {
			return nil, err
		}
		return fromThrift2NamespaceDescriptorList(list), nil
	}
	return nil, ErrorClientVersion
}
