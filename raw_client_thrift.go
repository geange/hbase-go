package hbase

import (
	"context"
	"git.apache.org/thrift.git/lib/go/thrift"
	thrift1 "github.com/geange/hbase-go/thrift/v1"
	"net"
	"strconv"
	"time"
)

type RawClientV1 interface {
	Open() error
	Close() error
	// Brings a table on-line (enables it)
	//
	// Parameters:
	//  - TableName: name of the table
	EnableTable(tableName thrift1.Bytes) (err error)
	// Disables a table (takes it off-line) If it is being served, the master
	// will tell the servers to stop serving it.
	//
	// Parameters:
	//  - TableName: name of the table
	DisableTable(tableName thrift1.Bytes) (err error)
	// @return true if table is on-line
	//
	// Parameters:
	//  - TableName: name of the table to check
	IsTableEnabled(tableName thrift1.Bytes) (r bool, err error)
	// Parameters:
	//  - TableNameOrRegionName
	Compact(tableNameOrRegionName thrift1.Bytes) (err error)
	// Parameters:
	//  - TableNameOrRegionName
	MajorCompact(tableNameOrRegionName thrift1.Bytes) (err error)
	// List all the userspace tables.
	//
	// @return returns a list of names
	GetTableNames() (r [][]byte, err error)
	// List all the column families assoicated with a table.
	//
	// @return list of column family descriptors
	//
	// Parameters:
	//  - TableName: table name
	GetColumnDescriptors(tableName thrift1.Text) (r map[string]*thrift1.ColumnDescriptor, err error)
	// List the regions associated with a table.
	//
	// @return list of region descriptors
	//
	// Parameters:
	//  - TableName: table name
	GetTableRegions(tableName thrift1.Text) (r []*thrift1.TRegionInfo, err error)
	// Create a table with the specified column families.  The name
	// field for each ColumnDescriptor must be set and must end in a
	// colon (:). All other fields are optional and will get default
	// values if not explicitly specified.
	//
	// @throws IllegalArgument if an input parameter is invalid
	//
	// @throws AlreadyExists if the table name already exists
	//
	// Parameters:
	//  - TableName: name of table to create
	//  - ColumnFamilies: list of column family descriptors
	CreateTable(tableName thrift1.Text, columnFamilies []*thrift1.ColumnDescriptor) (err error)
	// Deletes a table
	//
	// @throws IOError if table doesn't exist on server or there was some other
	// problem
	//
	// Parameters:
	//  - TableName: name of table to delete
	DeleteTable(tableName thrift1.Text) (err error)
	// Get a single TCell for the specified table, row, and column at the
	// latest timestamp. Returns an empty list if no such value exists.
	//
	// @return value for specified row/column
	//
	// Parameters:
	//  - TableName: name of table
	//  - Row: row key
	//  - Column: column name
	//  - Attributes: Get attributes
	Get(tableName thrift1.Text, row thrift1.Text, column thrift1.Text, attributes map[string]thrift1.Text) (r []*thrift1.TCell, err error)
	// Get the specified number of versions for the specified table,
	// row, and column.
	//
	// @return list of cells for specified row/column
	//
	// Parameters:
	//  - TableName: name of table
	//  - Row: row key
	//  - Column: column name
	//  - NumVersions: number of versions to retrieve
	//  - Attributes: Get attributes
	GetVer(tableName thrift1.Text, row thrift1.Text, column thrift1.Text, numVersions int32, attributes map[string]thrift1.Text) (r []*thrift1.TCell, err error)
	// Get the specified number of versions for the specified table,
	// row, and column.  Only versions less than or equal to the specified
	// timestamp will be returned.
	//
	// @return list of cells for specified row/column
	//
	// Parameters:
	//  - TableName: name of table
	//  - Row: row key
	//  - Column: column name
	//  - Timestamp: timestamp
	//  - NumVersions: number of versions to retrieve
	//  - Attributes: Get attributes
	GetVerTs(tableName thrift1.Text, row thrift1.Text, column thrift1.Text, timestamp int64, numVersions int32, attributes map[string]thrift1.Text) (r []*thrift1.TCell, err error)
	// Get all the data for the specified table and row at the latest
	// timestamp. Returns an empty list if the row does not exist.
	//
	// @return TRowResult containing the row and map of columns to TCells
	//
	// Parameters:
	//  - TableName: name of table
	//  - Row: row key
	//  - Attributes: Get attributes
	GetRow(tableName thrift1.Text, row thrift1.Text, attributes map[string]thrift1.Text) (r []*thrift1.TRowResult_, err error)
	// Get the specified columns for the specified table and row at the latest
	// timestamp. Returns an empty list if the row does not exist.
	//
	// @return TRowResult containing the row and map of columns to TCells
	//
	// Parameters:
	//  - TableName: name of table
	//  - Row: row key
	//  - Columns: List of columns to return, null for all columns
	//  - Attributes: Get attributes
	GetRowWithColumns(tableName thrift1.Text, row thrift1.Text, columns [][]byte, attributes map[string]thrift1.Text) (r []*thrift1.TRowResult_, err error)
	// Get all the data for the specified table and row at the specified
	// timestamp. Returns an empty list if the row does not exist.
	//
	// @return TRowResult containing the row and map of columns to TCells
	//
	// Parameters:
	//  - TableName: name of the table
	//  - Row: row key
	//  - Timestamp: timestamp
	//  - Attributes: Get attributes
	GetRowTs(tableName thrift1.Text, row thrift1.Text, timestamp int64, attributes map[string]thrift1.Text) (r []*thrift1.TRowResult_, err error)
	// Get the specified columns for the specified table and row at the specified
	// timestamp. Returns an empty list if the row does not exist.
	//
	// @return TRowResult containing the row and map of columns to TCells
	//
	// Parameters:
	//  - TableName: name of table
	//  - Row: row key
	//  - Columns: List of columns to return, null for all columns
	//  - Timestamp
	//  - Attributes: Get attributes
	GetRowWithColumnsTs(tableName thrift1.Text, row thrift1.Text, columns [][]byte, timestamp int64, attributes map[string]thrift1.Text) (r []*thrift1.TRowResult_, err error)
	// Get all the data for the specified table and rows at the latest
	// timestamp. Returns an empty list if no rows exist.
	//
	// @return TRowResult containing the rows and map of columns to TCells
	//
	// Parameters:
	//  - TableName: name of table
	//  - Rows: row keys
	//  - Attributes: Get attributes
	GetRows(tableName thrift1.Text, rows [][]byte, attributes map[string]thrift1.Text) (r []*thrift1.TRowResult_, err error)
	// Get the specified columns for the specified table and rows at the latest
	// timestamp. Returns an empty list if no rows exist.
	//
	// @return TRowResult containing the rows and map of columns to TCells
	//
	// Parameters:
	//  - TableName: name of table
	//  - Rows: row keys
	//  - Columns: List of columns to return, null for all columns
	//  - Attributes: Get attributes
	GetRowsWithColumns(tableName thrift1.Text, rows [][]byte, columns [][]byte, attributes map[string]thrift1.Text) (r []*thrift1.TRowResult_, err error)
	// Get all the data for the specified table and rows at the specified
	// timestamp. Returns an empty list if no rows exist.
	//
	// @return TRowResult containing the rows and map of columns to TCells
	//
	// Parameters:
	//  - TableName: name of the table
	//  - Rows: row keys
	//  - Timestamp: timestamp
	//  - Attributes: Get attributes
	GetRowsTs(tableName thrift1.Text, rows [][]byte, timestamp int64, attributes map[string]thrift1.Text) (r []*thrift1.TRowResult_, err error)
	// Get the specified columns for the specified table and rows at the specified
	// timestamp. Returns an empty list if no rows exist.
	//
	// @return TRowResult containing the rows and map of columns to TCells
	//
	// Parameters:
	//  - TableName: name of table
	//  - Rows: row keys
	//  - Columns: List of columns to return, null for all columns
	//  - Timestamp
	//  - Attributes: Get attributes
	GetRowsWithColumnsTs(tableName thrift1.Text, rows [][]byte, columns [][]byte, timestamp int64, attributes map[string]thrift1.Text) (r []*thrift1.TRowResult_, err error)
	// Apply a series of mutations (updates/deletes) to a row in a
	// single transaction.  If an exception is thrown, then the
	// transaction is aborted.  Default current timestamp is used, and
	// all entries will have an identical timestamp.
	//
	// Parameters:
	//  - TableName: name of table
	//  - Row: row key
	//  - Mutations: list of mutation commands
	//  - Attributes: Mutation attributes
	MutateRow(tableName thrift1.Text, row thrift1.Text, mutations []*thrift1.Mutation, attributes map[string]thrift1.Text) (err error)
	// Apply a series of mutations (updates/deletes) to a row in a
	// single transaction.  If an exception is thrown, then the
	// transaction is aborted.  The specified timestamp is used, and
	// all entries will have an identical timestamp.
	//
	// Parameters:
	//  - TableName: name of table
	//  - Row: row key
	//  - Mutations: list of mutation commands
	//  - Timestamp: timestamp
	//  - Attributes: Mutation attributes
	MutateRowTs(tableName thrift1.Text, row thrift1.Text, mutations []*thrift1.Mutation, timestamp int64, attributes map[string]thrift1.Text) (err error)
	// Apply a series of batches (each a series of mutations on a single row)
	// in a single transaction.  If an exception is thrown, then the
	// transaction is aborted.  Default current timestamp is used, and
	// all entries will have an identical timestamp.
	//
	// Parameters:
	//  - TableName: name of table
	//  - RowBatches: list of row batches
	//  - Attributes: Mutation attributes
	MutateRows(tableName thrift1.Text, rowBatches []*thrift1.BatchMutation, attributes map[string]thrift1.Text) (err error)
	// Apply a series of batches (each a series of mutations on a single row)
	// in a single transaction.  If an exception is thrown, then the
	// transaction is aborted.  The specified timestamp is used, and
	// all entries will have an identical timestamp.
	//
	// Parameters:
	//  - TableName: name of table
	//  - RowBatches: list of row batches
	//  - Timestamp: timestamp
	//  - Attributes: Mutation attributes
	MutateRowsTs(tableName thrift1.Text, rowBatches []*thrift1.BatchMutation, timestamp int64, attributes map[string]thrift1.Text) (err error)
	// Atomically increment the column value specified.  Returns the next value post increment.
	//
	// Parameters:
	//  - TableName: name of table
	//  - Row: row to increment
	//  - Column: name of column
	//  - Value: amount to increment by
	AtomicIncrement(tableName thrift1.Text, row thrift1.Text, column thrift1.Text, value int64) (r int64, err error)
	// Delete all cells that match the passed row and column.
	//
	// Parameters:
	//  - TableName: name of table
	//  - Row: Row to update
	//  - Column: name of column whose value is to be deleted
	//  - Attributes: Delete attributes
	DeleteAll(tableName thrift1.Text, row thrift1.Text, column thrift1.Text, attributes map[string]thrift1.Text) (err error)
	// Delete all cells that match the passed row and column and whose
	// timestamp is equal-to or older than the passed timestamp.
	//
	// Parameters:
	//  - TableName: name of table
	//  - Row: Row to update
	//  - Column: name of column whose value is to be deleted
	//  - Timestamp: timestamp
	//  - Attributes: Delete attributes
	DeleteAllTs(tableName thrift1.Text, row thrift1.Text, column thrift1.Text, timestamp int64, attributes map[string]thrift1.Text) (err error)
	// Completely delete the row's cells.
	//
	// Parameters:
	//  - TableName: name of table
	//  - Row: key of the row to be completely deleted.
	//  - Attributes: Delete attributes
	DeleteAllRow(tableName thrift1.Text, row thrift1.Text, attributes map[string]thrift1.Text) (err error)
	// Increment a cell by the ammount.
	// Increments can be applied async if hbase.regionserver.thrift.coalesceIncrement is set to true.
	// False is the default.  Turn to true if you need the extra performance and can accept some
	// data loss if a thrift server dies with increments still in the queue.
	//
	// Parameters:
	//  - Increment: The single increment to apply
	Increment(increment *thrift1.TIncrement) (err error)
	// Parameters:
	//  - Increments: The list of increments
	IncrementRows(increments []*thrift1.TIncrement) (err error)
	// Completely delete the row's cells marked with a timestamp
	// equal-to or older than the passed timestamp.
	//
	// Parameters:
	//  - TableName: name of table
	//  - Row: key of the row to be completely deleted.
	//  - Timestamp: timestamp
	//  - Attributes: Delete attributes
	DeleteAllRowTs(tableName thrift1.Text, row thrift1.Text, timestamp int64, attributes map[string]thrift1.Text) (err error)
	// Get a scanner on the current table, using the Scan instance
	// for the scan parameters.
	//
	// Parameters:
	//  - TableName: name of table
	//  - Scan: Scan instance
	//  - Attributes: Scan attributes
	ScannerOpenWithScan(tableName thrift1.Text, scan *thrift1.TScan, attributes map[string]thrift1.Text) (r thrift1.ScannerID, err error)
	// Get a scanner on the current table starting at the specified row and
	// ending at the last row in the table.  Return the specified columns.
	//
	// @return scanner id to be used with other scanner procedures
	//
	// Parameters:
	//  - TableName: name of table
	//  - StartRow: Starting row in table to scan.
	// Send "" (empty string) to start at the first row.
	//  - Columns: columns to scan. If column name is a column family, all
	// columns of the specified column family are returned. It's also possible
	// to pass a regex in the column qualifier.
	//  - Attributes: Scan attributes
	ScannerOpen(tableName thrift1.Text, startRow thrift1.Text, columns [][]byte, attributes map[string]thrift1.Text) (r thrift1.ScannerID, err error)
	// Get a scanner on the current table starting and stopping at the
	// specified rows.  ending at the last row in the table.  Return the
	// specified columns.
	//
	// @return scanner id to be used with other scanner procedures
	//
	// Parameters:
	//  - TableName: name of table
	//  - StartRow: Starting row in table to scan.
	// Send "" (empty string) to start at the first row.
	//  - StopRow: row to stop scanning on. This row is *not* included in the
	// scanner's results
	//  - Columns: columns to scan. If column name is a column family, all
	// columns of the specified column family are returned. It's also possible
	// to pass a regex in the column qualifier.
	//  - Attributes: Scan attributes
	ScannerOpenWithStop(tableName thrift1.Text, startRow thrift1.Text, stopRow thrift1.Text, columns [][]byte, attributes map[string]thrift1.Text) (r thrift1.ScannerID, err error)
	// Open a scanner for a given prefix.  That is all rows will have the specified
	// prefix. No other rows will be returned.
	//
	// @return scanner id to use with other scanner calls
	//
	// Parameters:
	//  - TableName: name of table
	//  - StartAndPrefix: the prefix (and thus start row) of the keys you want
	//  - Columns: the columns you want returned
	//  - Attributes: Scan attributes
	ScannerOpenWithPrefix(tableName thrift1.Text, startAndPrefix thrift1.Text, columns [][]byte, attributes map[string]thrift1.Text) (r thrift1.ScannerID, err error)
	// Get a scanner on the current table starting at the specified row and
	// ending at the last row in the table.  Return the specified columns.
	// Only values with the specified timestamp are returned.
	//
	// @return scanner id to be used with other scanner procedures
	//
	// Parameters:
	//  - TableName: name of table
	//  - StartRow: Starting row in table to scan.
	// Send "" (empty string) to start at the first row.
	//  - Columns: columns to scan. If column name is a column family, all
	// columns of the specified column family are returned. It's also possible
	// to pass a regex in the column qualifier.
	//  - Timestamp: timestamp
	//  - Attributes: Scan attributes
	ScannerOpenTs(tableName thrift1.Text, startRow thrift1.Text, columns [][]byte, timestamp int64, attributes map[string]thrift1.Text) (r thrift1.ScannerID, err error)
	// Get a scanner on the current table starting and stopping at the
	// specified rows.  ending at the last row in the table.  Return the
	// specified columns.  Only values with the specified timestamp are
	// returned.
	//
	// @return scanner id to be used with other scanner procedures
	//
	// Parameters:
	//  - TableName: name of table
	//  - StartRow: Starting row in table to scan.
	// Send "" (empty string) to start at the first row.
	//  - StopRow: row to stop scanning on. This row is *not* included in the
	// scanner's results
	//  - Columns: columns to scan. If column name is a column family, all
	// columns of the specified column family are returned. It's also possible
	// to pass a regex in the column qualifier.
	//  - Timestamp: timestamp
	//  - Attributes: Scan attributes
	ScannerOpenWithStopTs(tableName thrift1.Text, startRow thrift1.Text, stopRow thrift1.Text, columns [][]byte, timestamp int64, attributes map[string]thrift1.Text) (r thrift1.ScannerID, err error)
	// Returns the scanner's current row value and advances to the next
	// row in the table.  When there are no more rows in the table, or a key
	// greater-than-or-equal-to the scanner's specified stopRow is reached,
	// an empty list is returned.
	//
	// @return a TRowResult containing the current row and a map of the columns to TCells.
	//
	// @throws IllegalArgument if ScannerID is invalid
	//
	// @throws NotFound when the scanner reaches the end
	//
	// Parameters:
	//  - ID: id of a scanner returned by scannerOpen
	ScannerGet(id thrift1.ScannerID) (r []*thrift1.TRowResult_, err error)
	// Returns, starting at the scanner's current row value nbRows worth of
	// rows and advances to the next row in the table.  When there are no more
	// rows in the table, or a key greater-than-or-equal-to the scanner's
	// specified stopRow is reached,  an empty list is returned.
	//
	// @return a TRowResult containing the current row and a map of the columns to TCells.
	//
	// @throws IllegalArgument if ScannerID is invalid
	//
	// @throws NotFound when the scanner reaches the end
	//
	// Parameters:
	//  - ID: id of a scanner returned by scannerOpen
	//  - NbRows: number of results to return
	ScannerGetList(id thrift1.ScannerID, nbRows int32) (r []*thrift1.TRowResult_, err error)
	// Closes the server-state associated with an open scanner.
	//
	// @throws IllegalArgument if ScannerID is invalid
	//
	// Parameters:
	//  - ID: id of a scanner returned by scannerOpen
	ScannerClose(id thrift1.ScannerID) (err error)
	// Get the regininfo for the specified row. It scans
	// the metatable to find region's start and end keys.
	//
	// @return value for specified row/column
	//
	// Parameters:
	//  - Row: row key
	GetRegionInfo(row thrift1.Text) (r *thrift1.TRegionInfo, err error)
	// Appends values to one or more columns within a single row.
	//
	// @return values of columns after the append operation.
	//
	// Parameters:
	//  - Append: The single append operation to apply
	Append(append *thrift1.TAppend) (r []*thrift1.TCell, err error)
	// Atomically checks if a row/family/qualifier value matches the expected
	// value. If it does, it adds the corresponding mutation operation for put.
	//
	// @return true if the new put was executed, false otherwise
	//
	// Parameters:
	//  - TableName: name of table
	//  - Row: row key
	//  - Column: column name
	//  - Value: the expected value for the column parameter, if not
	// provided the check is for the non-existence of the
	// column in question
	//  - Mput: mutation for the put
	//  - Attributes: Mutation attributes
	CheckAndPut(tableName thrift1.Text, row thrift1.Text, column thrift1.Text, value thrift1.Text, mput *thrift1.Mutation, attributes map[string]thrift1.Text) (r bool, err error)
}

type rawClientV1 struct {
	*thrift1.HbaseClient
	transport thrift.TTransport
	host      string
	port      int
}

func (c *rawClientV1) Open() error {
	if c.transport.IsOpen() {
		return nil
	}
	return c.transport.Open()
}

func (c *rawClientV1) Close() error {
	if err := c.transport.Flush(); err != nil {
		return err
	}
	return c.transport.Close()
}

func NewRawClientV1(ctx context.Context, option RawClientOption) (rc RawClientV1, err error) {
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
		transport = thrift.NewTTransportFactory().GetTransport(socket)
	} else {
		transport = thrift.NewTBufferedTransportFactory(option.BufferSize).GetTransport(socket)
	}
	if err != nil {
		return nil, err
	}

	client := thrift1.NewHbaseClientFactory(transport, thrift.NewTBinaryProtocolFactoryDefault())

	return &rawClientV1{
		HbaseClient: client,
		host:        option.Host,
		port:        option.Port,
		transport:   transport,
	}, nil
}
