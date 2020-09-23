package hbase

import (
	"context"
	"sync"
)

type ThriftWithPool interface {
	// Test for the existence of columns in the table, as specified in the TGet.
	//
	// @return true if the specified TGet matches one or more keys, false if not
	//
	// Parameters:
	//  - Table: the table to check on
	//  - Get: the TGet to check for
	Exists(ctx context.Context, table []byte, get *Get) (r bool, err error)
	// Test for the existence of columns in the table, as specified by the TGets.
	//
	// This will return an array of booleans. Each value will be true if the related Get matches
	// one or more keys, false if not.
	//
	// Parameters:
	//  - Table: the table to check on
	//  - Gets: a list of TGets to check for
	ExistsAll(ctx context.Context, table []byte, gets []*Get) (r []bool, err error)
	// Method for getting data from a row.
	//
	// If the row cannot be found an empty Result is returned.
	// This can be checked by the empty field of the TResult
	//
	// @return the result
	//
	// Parameters:
	//  - Table: the table to get from
	//  - Get: the TGet to fetch
	Get(ctx context.Context, table []byte, get *Get) (r *Result, err error)
	// Method for getting multiple rows.
	//
	// If a row cannot be found there will be a null
	// value in the result list for that TGet at the
	// same position.
	//
	// So the Results are in the same order as the TGets.
	//
	// Parameters:
	//  - Table: the table to get from
	//  - Gets: a list of TGets to fetch, the Result list
	// will have the Results at corresponding positions
	// or null if there was an error
	GetMultiple(ctx context.Context, table []byte, Gets []*Get) (r []*Result, err error)
	// Commit a TPut to a table.
	//
	// Parameters:
	//  - Table: the table to put data in
	//  - Tput: the TPut to put
	Put(ctx context.Context, table []byte, put *Put) (err error)
	// Atomically checks if a row/family/qualifier value matches the expected
	// value. If it does, it adds the TPut.
	//
	// @return true if the new put was executed, false otherwise
	//
	// Parameters:
	//  - Table: to check in and put to
	//  - Row: row to check
	//  - Family: column family to check
	//  - Qualifier: column qualifier to check
	//  - Value: the expected value, if not provided the
	// check is for the non-existence of the
	// column in question
	//  - Tput: the TPut to put if the check succeeds
	CheckAndPut(ctx context.Context, table []byte, row []byte, family []byte, qualifier []byte, value []byte, put *Put) (r bool, err error)
	// Commit a List of Puts to the table.
	//
	// Parameters:
	//  - Table: the table to put data in
	//  - Tputs: a list of TPuts to commit
	PutMultiple(ctx context.Context, table []byte, puts []*Put) (err error)
	// Deletes as specified by the TDelete.
	//
	// Note: "delete" is a reserved keyword and cannot be used in Thrift
	// thus the inconsistent naming scheme from the other functions.
	//
	// Parameters:
	//  - Table: the table to delete from
	//  - Delete: the TDelete to delete
	DeleteSingle(ctx context.Context, table []byte, delete *Delete) (err error)
	// Bulk commit a List of TDeletes to the table.
	//
	// Throws a TIOError if any of the deletes fail.
	//
	// Always returns an empty list for backwards compatibility.
	//
	// Parameters:
	//  - Table: the table to delete from
	//  - Deletes: list of TDeletes to delete
	DeleteMultiple(ctx context.Context, table []byte, deletes []*Delete) (r []*Delete, err error)
	// Atomically checks if a row/family/qualifier value matches the expected
	// value. If it does, it adds the delete.
	//
	// @return true if the new delete was executed, false otherwise
	//
	// Parameters:
	//  - Table: to check in and delete from
	//  - Row: row to check
	//  - Family: column family to check
	//  - Qualifier: column qualifier to check
	//  - Value: the expected value, if not provided the
	// check is for the non-existence of the
	// column in question
	//  - Tdelete: the TDelete to execute if the check succeeds
	CheckAndDelete(ctx context.Context, table []byte, row []byte, family []byte, qualifier []byte, value []byte, delete *Delete) (r bool, err error)
	// Parameters:
	//  - Table: the table to increment the value on
	//  - Increment: the TIncrement to increment
	Increment(ctx context.Context, table []byte, increment *Increment) (r *Result, err error)
	// Parameters:
	//  - Table: the table to append the value on
	//  - Tappend: the TAppend to append
	Append(ctx context.Context, table []byte, append *TAppend) (r *Result, err error)
	// Get a Scanner for the provided TScan object.
	//
	// @return Scanner Id to be used with other scanner procedures
	//
	// Parameters:
	//  - Table: the table to get the Scanner for
	//  - Tscan: the scan object to get a Scanner for
	OpenScanner(ctx context.Context, table []byte, scan *Scan) (r int32, err error)
	// Grabs multiple rows from a Scanner.
	//
	// @return Between zero and numRows TResults
	//
	// Parameters:
	//  - ScannerId: the Id of the Scanner to return rows from. This is an Id returned from the openScanner function.
	//  - NumRows: number of rows to return
	GetScannerRows(ctx context.Context, scannerId int32, numRows int32) (r []*Result, err error)
	// Closes the scanner. Should be called to free server side resources timely.
	// Typically close once the scanner is not needed anymore, i.e. after looping
	// over it to get all the required rows.
	//
	// Parameters:
	//  - ScannerId: the Id of the Scanner to close *
	CloseScanner(ctx context.Context, scannerId int32) (err error)
	// mutateRow performs multiple mutations atomically on a single row.
	//
	// Parameters:
	//  - Table: table to apply the mutations
	//  - TrowMutations: mutations to apply
	MutateRow(ctx context.Context, table []byte, trowMutations *RowMutations) (err error)
	// Get results for the provided TScan object.
	// This helper function opens a scanner, get the results and close the scanner.
	//
	// @return between zero and numRows TResults
	//
	// Parameters:
	//  - Table: the table to get the Scanner for
	//  - Tscan: the scan object to get a Scanner for
	//  - NumRows: number of rows to return
	GetScannerResults(ctx context.Context, table []byte, tscan *Scan, numRows int32) (r []*Result, err error)
	// Given a table and a row get the location of the region that
	// would contain the given row key.
	//
	// reload = true means the cache will be cleared and the location
	// will be fetched from meta.
	//
	// Parameters:
	//  - Table
	//  - Row
	//  - Reload
	GetRegionLocation(ctx context.Context, table []byte, row []byte, reload bool) (r *THRegionLocation, err error)
	// Get all of the region locations for a given table.
	//
	//
	// Parameters:
	//  - Table
	GetAllRegionLocations(ctx context.Context, table []byte) (r []*THRegionLocation, err error)
	// Atomically checks if a row/family/qualifier value matches the expected
	// value. If it does, it mutates the row.
	//
	// @return true if the row was mutated, false otherwise
	//
	// Parameters:
	//  - Table: to check in and delete from
	//  - Row: row to check
	//  - Family: column family to check
	//  - Qualifier: column qualifier to check
	//  - CompareOp: comparison to make on the value
	//  - Value: the expected value to be compared against, if not provided the
	// check is for the non-existence of the column in question
	//  - RowMutations: row mutations to execute if the value matches
	CheckAndMutate(ctx context.Context, table []byte, row []byte, family []byte, qualifier []byte, compareOp TCompareOp, value []byte, rowMutations *RowMutations) (r bool, err error)
	// Get a table descriptor.
	// @return the TableDescriptor of the giving tablename
	//
	//
	// Parameters:
	//  - Table: the tablename of the table to get tableDescriptor
	GetTableDescriptor(ctx context.Context, table *TableName) (r *TableDescriptor, err error)
	// Get table descriptors of tables.
	// @return the TableDescriptor of the giving tablename
	//
	//
	// Parameters:
	//  - Tables: the tablename list of the tables to get tableDescriptor
	GetTableDescriptors(ctx context.Context, tables []*TableName) (r []*TableDescriptor, err error)
	//
	// @return true if table exists already, false if not
	//
	//
	// Parameters:
	//  - TableName: the tablename of the tables to check
	TableExists(ctx context.Context, tableName *TableName) (r bool, err error)
	// Get table descriptors of tables that match the given pattern
	// @return the tableDescriptors of the matching table
	//
	//
	// Parameters:
	//  - Regex: The regular expression to match against
	//  - IncludeSysTables: set to false if match only against userspace tables
	GetTableDescriptorsByPattern(ctx context.Context, regex string, includeSysTables bool) (r []*TableDescriptor, err error)
	// Get table descriptors of tables in the given namespace
	// @return the tableDescriptors in the namespce
	//
	//
	// Parameters:
	//  - Name: The namesapce's name
	GetTableDescriptorsByNamespace(ctx context.Context, name string) (r []*TableDescriptor, err error)
	// Get table names of tables that match the given pattern
	// @return the table names of the matching table
	//
	//
	// Parameters:
	//  - Regex: The regular expression to match against
	//  - IncludeSysTables: set to false if match only against userspace tables
	GetTableNamesByPattern(ctx context.Context, regex string, includeSysTables bool) (r []*TableName, err error)
	// Get table names of tables in the given namespace
	// @return the table names of the matching table
	//
	//
	// Parameters:
	//  - Name: The namesapce's name
	GetTableNamesByNamespace(ctx context.Context, name string) (r []*TableName, err error)
	// Creates a new table with an initial set of empty regions defined by the specified split keys.
	// The total number of regions created will be the number of split keys plus one. Synchronous
	// operation.
	//
	//
	// Parameters:
	//  - Desc: table descriptor for table
	//  - SplitKeys: rray of split keys for the initial regions of the table
	CreateTable(ctx context.Context, desc *TableDescriptor, splitKeys [][]byte) (err error)
	// Deletes a table. Synchronous operation.
	//
	//
	// Parameters:
	//  - TableName: the tablename to delete
	DeleteTable(ctx context.Context, tableName *TableName) (err error)
	// Truncate a table. Synchronous operation.
	//
	//
	// Parameters:
	//  - TableName: the tablename to truncate
	//  - PreserveSplits: whether to  preserve previous splits
	TruncateTable(ctx context.Context, tableName *TableName, preserveSplits bool) (err error)
	// Enalbe a table
	//
	//
	// Parameters:
	//  - TableName: the tablename to enable
	EnableTable(ctx context.Context, tableName *TableName) (err error)
	// Disable a table
	//
	//
	// Parameters:
	//  - TableName: the tablename to disable
	DisableTable(ctx context.Context, tableName *TableName) (err error)
	//
	// @return true if table is enabled, false if not
	//
	//
	// Parameters:
	//  - TableName: the tablename to check
	IsTableEnabled(ctx context.Context, tableName *TableName) (r bool, err error)
	//
	// @return true if table is disabled, false if not
	//
	//
	// Parameters:
	//  - TableName: the tablename to check
	IsTableDisabled(ctx context.Context, tableName *TableName) (r bool, err error)
	//
	// @return true if table is available, false if not
	//
	//
	// Parameters:
	//  - TableName: the tablename to check
	IsTableAvailable(ctx context.Context, tableName *TableName) (r bool, err error)
	//  * Use this api to check if the table has been created with the specified number of splitkeys
	//  * which was used while creating the given table. Note : If this api is used after a table's
	//  * region gets splitted, the api may return false.
	//  *
	//  * @return true if table is available, false if not
	//  *
	//  * @deprecated Since 2.2.0. Because the same method in Table interface has been deprecated
	//  * since 2.0.0, we will remove it in 3.0.0 release.
	//  * Use {@link #isTableAvailable(TTableName tableName)} instead
	// *
	//
	// Parameters:
	//  - TableName: the tablename to check
	//  - SplitKeys: keys to check if the table has been created with all split keys
	IsTableAvailableWithSplit(ctx context.Context, tableName *TableName, splitKeys [][]byte) (r bool, err error)
	// Add a column family to an existing table. Synchronous operation.
	//
	//
	// Parameters:
	//  - TableName: the tablename to add column family to
	//  - Column: column family descriptor of column family to be added
	AddColumnFamily(ctx context.Context, tableName *TableName, column *ColumnFamilyDescriptor) (err error)
	// Delete a column family from a table. Synchronous operation.
	//
	//
	// Parameters:
	//  - TableName: the tablename to delete column family from
	//  - Column: name of column family to be deleted
	DeleteColumnFamily(ctx context.Context, tableName *TableName, column []byte) (err error)
	// Modify an existing column family on a table. Synchronous operation.
	//
	//
	// Parameters:
	//  - TableName: the tablename to modify column family
	//  - Column: column family descriptor of column family to be modified
	ModifyColumnFamily(ctx context.Context, tableName *TableName, column *ColumnFamilyDescriptor) (err error)
	// Modify an existing table
	//
	//
	// Parameters:
	//  - Desc: the descriptor of the table to modify
	ModifyTable(ctx context.Context, desc *TableDescriptor) (err error)
	// Create a new namespace. Blocks until namespace has been successfully created or an exception is
	// thrown
	//
	//
	// Parameters:
	//  - NamespaceDesc: descriptor which describes the new namespace
	CreateNamespace(ctx context.Context, namespaceDesc *NamespaceDescriptor) (err error)
	// Modify an existing namespace.  Blocks until namespace has been successfully modified or an
	// exception is thrown
	//
	//
	// Parameters:
	//  - NamespaceDesc: descriptor which describes the new namespace
	ModifyNamespace(ctx context.Context, namespaceDesc *NamespaceDescriptor) (err error)
	// Delete an existing namespace. Only empty namespaces (no tables) can be removed.
	// Blocks until namespace has been successfully deleted or an
	// exception is thrown.
	//
	//
	// Parameters:
	//  - Name: namespace name
	DeleteNamespace(ctx context.Context, name string) (err error)
	// Get a namespace descriptor by name.
	// @retrun the descriptor
	//
	//
	// Parameters:
	//  - Name: name of namespace descriptor
	GetNamespaceDescriptor(ctx context.Context, name string) (r *NamespaceDescriptor, err error)
	// @return all namespaces
	//
	ListNamespaceDescriptors(ctx context.Context) (r []*NamespaceDescriptor, err error)
}

type client struct {
	sync.Mutex
}
