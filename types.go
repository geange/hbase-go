package hbase

import thrift2 "github.com/geange/hbase-go/thrift/v2"

type Get struct {
	Row             []byte            `thrift:"row,1,required" db:"row" json:"row"`
	Columns         []*Column         `thrift:"columns,2" db:"columns" json:"columns,omitempty"`
	Timestamp       *int64            `thrift:"timestamp,3" db:"timestamp" json:"timestamp,omitempty"`
	TimeRange       *TimeRange        `thrift:"timeRange,4" db:"timeRange" json:"timeRange,omitempty"`
	MaxVersions     *int32            `thrift:"maxVersions,5" db:"maxVersions" json:"maxVersions,omitempty"`
	FilterString    []byte            `thrift:"filterString,6" db:"filterString" json:"filterString,omitempty"`
	Attributes      map[string][]byte `thrift:"attributes,7" db:"attributes" json:"attributes,omitempty"`
	Authorizations  *Authorization    `thrift:"authorizations,8" db:"authorizations" json:"authorizations,omitempty"`
	Consistency     *TConsistency     `thrift:"consistency,9" db:"consistency" json:"consistency,omitempty"`
	TargetReplicaId *int32            `thrift:"targetReplicaId,10" db:"targetReplicaId" json:"targetReplicaId,omitempty"`
	CacheBlocks     *bool             `thrift:"cacheBlocks,11" db:"cacheBlocks" json:"cacheBlocks,omitempty"`
	StoreLimit      *int32            `thrift:"storeLimit,12" db:"storeLimit" json:"storeLimit,omitempty"`
	StoreOffset     *int32            `thrift:"storeOffset,13" db:"storeOffset" json:"storeOffset,omitempty"`
	ExistenceOnly   *bool             `thrift:"existence_only,14" db:"existence_only" json:"existence_only,omitempty"`
	FilterBytes     []byte            `thrift:"filterBytes,15" db:"filterBytes" json:"filterBytes,omitempty"`
}

func toThrift2TGet(ptr *Get) *thrift2.TGet {
	return &thrift2.TGet{
		Row:             ptr.Row,
		Columns:         toThrift2ColumnList(ptr.Columns),
		Timestamp:       ptr.Timestamp,
		TimeRange:       toThrift2TTimeRange(ptr.TimeRange),
		MaxVersions:     ptr.MaxVersions,
		FilterString:    ptr.FilterString,
		Attributes:      ptr.Attributes,
		Authorizations:  toThrift2TAuthorization(ptr.Authorizations),
		Consistency:     toThrift2TConsistency(ptr.Consistency),
		TargetReplicaId: ptr.TargetReplicaId,
		CacheBlocks:     ptr.CacheBlocks,
		StoreLimit:      ptr.StoreLimit,
		StoreOffset:     ptr.StoreOffset,
		ExistenceOnly:   ptr.ExistenceOnly,
		FilterBytes:     ptr.FilterBytes,
	}
}

type Column struct {
	Family    []byte
	Qualifier []byte
	Timestamp *int64
}

func toThrift2TColumn(ptr *Column) *thrift2.TColumn {
	if ptr == nil {
		return nil
	}
	return &thrift2.TColumn{
		Family:    ptr.Family,
		Qualifier: ptr.Qualifier,
		Timestamp: ptr.Timestamp,
	}
}

func fromThrift2Column(ptr *thrift2.TColumn) *Column {
	if ptr == nil {
		return nil
	}
	return &Column{
		Family:    ptr.Family,
		Qualifier: ptr.Qualifier,
		Timestamp: ptr.Timestamp,
	}
}

type TimeRange struct {
	MinStamp int64
	MaxStamp int64
}

func toThrift2TTimeRange(ptr *TimeRange) *thrift2.TTimeRange {
	if ptr == nil {
		return nil
	}
	return &thrift2.TTimeRange{
		MinStamp: ptr.MinStamp,
		MaxStamp: ptr.MaxStamp,
	}
}

type Authorization struct {
	Labels []string
}

func toThrift2TAuthorization(ptr *Authorization) *thrift2.TAuthorization {
	if ptr == nil {
		return nil
	}
	return &thrift2.TAuthorization{Labels: ptr.Labels}
}

type TConsistency int64

func toThrift2TConsistency(ptr *TConsistency) *thrift2.TConsistency {
	if ptr == nil {
		return nil
	}
	return (*thrift2.TConsistency)(ptr)
}

type Result struct {
	Row          []byte
	ColumnValues []*ColumnValue
	Stale        bool
	Partial      bool
}

func toThrift2TResult(ptr *Result) *thrift2.TResult {
	if ptr == nil {
		return nil
	}
	return &thrift2.TResult{
		Row:          ptr.Row,
		ColumnValues: toThrift2ColumnValueList(ptr.ColumnValues),
		Stale:        ptr.Stale,
		Partial:      ptr.Partial,
	}
}

func fromThrift2Result(ptr *thrift2.TResult) *Result {
	if ptr == nil {
		return nil
	}
	return &Result{
		Row:          ptr.Row,
		ColumnValues: fromThrift2ColumnValueList(ptr.ColumnValues),
		Stale:        ptr.Stale,
		Partial:      ptr.Partial,
	}
}

type ColumnValue struct {
	Family    []byte
	Qualifier []byte
	Value     []byte
	Timestamp *int64
	Tags      []byte
	Type      *int8
}

func toThrift2TColumnValue(ptr *ColumnValue) *thrift2.TColumnValue {
	return &thrift2.TColumnValue{
		Family:    ptr.Family,
		Qualifier: ptr.Qualifier,
		Value:     ptr.Value,
		Timestamp: ptr.Timestamp,
		Tags:      ptr.Tags,
		Type:      ptr.Type,
	}
}

func fromThrift2ColumnValue(ptr *thrift2.TColumnValue) *ColumnValue {
	if ptr == nil {
		return nil
	}
	return &ColumnValue{
		Family:    ptr.Family,
		Qualifier: ptr.Qualifier,
		Value:     ptr.Value,
		Timestamp: ptr.Timestamp,
		Tags:      ptr.Tags,
		Type:      ptr.Type,
	}
}

type Put struct {
	Row          []byte
	ColumnValues []*ColumnValue
	Timestamp    *int64
	// unused field # 4
	Attributes     map[string][]byte
	Durability     *TDurability
	CellVisibility *CellVisibility
}

func toThrift2TPut(ptr *Put) *thrift2.TPut {
	if ptr == nil {
		return nil
	}
	return &thrift2.TPut{
		Row:            ptr.Row,
		ColumnValues:   toThrift2ColumnValueList(ptr.ColumnValues),
		Timestamp:      ptr.Timestamp,
		Attributes:     ptr.Attributes,
		Durability:     toThrift2TDurability(ptr.Durability),
		CellVisibility: toThrift2TCellVisibility(ptr.CellVisibility),
	}
}

type TDurability int64

func toThrift2TDurability(ptr *TDurability) *thrift2.TDurability {
	if ptr == nil {
		return nil
	}
	r := thrift2.TDurability(*ptr)
	return &r
}

const (
	UseDefault TDurability = 0
	SkipWAL    TDurability = 1
	AsyncWAL   TDurability = 2
	SyncWAL    TDurability = 3
	FsyncWAL   TDurability = 4
)

type CellVisibility struct {
	Expression *string
}

func toThrift2TCellVisibility(ptr *CellVisibility) *thrift2.TCellVisibility {
	if ptr == nil {
		return nil
	}
	return &thrift2.TCellVisibility{
		Expression: ptr.Expression,
	}
}

type Delete struct {
	Row        []byte
	Columns    []*Column
	Timestamp  *int64
	DeleteType TDeleteType
	Attributes map[string][]byte
	Durability *TDurability
}

func toThrift2TDelete(ptr *Delete) *thrift2.TDelete {
	if ptr == nil {
		return nil
	}
	return &thrift2.TDelete{
		Row:        ptr.Row,
		Columns:    toThrift2ColumnList(ptr.Columns),
		Timestamp:  ptr.Timestamp,
		DeleteType: thrift2.TDeleteType(ptr.DeleteType),
		Attributes: ptr.Attributes,
		Durability: toThrift2TDurability(ptr.Durability),
	}
}

func fromThrift2Delete(ptr *thrift2.TDelete) *Delete {
	if ptr == nil {
		return nil
	}
	return &Delete{
		Row:        ptr.Row,
		Columns:    fromThrift2ColumnList(ptr.Columns),
		Timestamp:  ptr.Timestamp,
		DeleteType: TDeleteType(ptr.DeleteType),
		Attributes: ptr.Attributes,
		Durability: (*TDurability)(ptr.Durability),
	}
}

type TDeleteType int64

const (
	DeleteColumn        TDeleteType = 0
	DeleteColumns       TDeleteType = 1
	DeleteFamily        TDeleteType = 2
	DeleteFamilyVersion TDeleteType = 3
)

type Increment struct {
	Row            []byte             `thrift:"row,1,required" db:"row" json:"row"`
	Columns        []*ColumnIncrement `thrift:"columns,2,required" db:"columns" json:"columns"`
	Attributes     map[string][]byte  `thrift:"attributes,4" db:"attributes" json:"attributes,omitempty"`
	Durability     *TDurability       `thrift:"durability,5" db:"durability" json:"durability,omitempty"`
	CellVisibility *CellVisibility    `thrift:"cellVisibility,6" db:"cellVisibility" json:"cellVisibility,omitempty"`
	ReturnResults  *bool              `thrift:"returnResults,7" db:"returnResults" json:"returnResults,omitempty"`
}

func toThrift2TIncrement(ptr *Increment) *thrift2.TIncrement {
	if ptr == nil {
		return nil
	}
	return &thrift2.TIncrement{
		Row:            ptr.Row,
		Columns:        toThrift2TColumnIncrementList(ptr.Columns),
		Attributes:     ptr.Attributes,
		Durability:     (*thrift2.TDurability)(ptr.Durability),
		CellVisibility: toThrift2TCellVisibility(ptr.CellVisibility),
		ReturnResults:  ptr.ReturnResults,
	}
}

type ColumnIncrement struct {
	Family    []byte `thrift:"family,1,required" db:"family" json:"family"`
	Qualifier []byte `thrift:"qualifier,2,required" db:"qualifier" json:"qualifier"`
	Amount    int64  `thrift:"amount,3" db:"amount" json:"amount"`
}

func toThrift2TColumnIncrement(ptr *ColumnIncrement) *thrift2.TColumnIncrement {
	if ptr == nil {
		return nil
	}
	return &thrift2.TColumnIncrement{
		Family:    ptr.Family,
		Qualifier: ptr.Qualifier,
		Amount:    ptr.Amount,
	}
}

type Append struct {
	Row            []byte
	Columns        []*ColumnValue
	Attributes     map[string][]byte
	Durability     *TDurability
	CellVisibility *CellVisibility
	ReturnResults  *bool
}

func toThrift2TAppend(ptr *Append) *thrift2.TAppend {
	if ptr == nil {
		return nil
	}
	return &thrift2.TAppend{
		Row:            ptr.Row,
		Columns:        toThrift2ColumnValueList(ptr.Columns),
		Attributes:     ptr.Attributes,
		Durability:     (*thrift2.TDurability)(ptr.Durability),
		CellVisibility: toThrift2TCellVisibility(ptr.CellVisibility),
		ReturnResults:  ptr.ReturnResults,
	}
}

type Scan struct {
	StartRow           []byte                `thrift:"startRow,1" db:"startRow" json:"startRow,omitempty"`
	StopRow            []byte                `thrift:"stopRow,2" db:"stopRow" json:"stopRow,omitempty"`
	Columns            []*Column             `thrift:"columns,3" db:"columns" json:"columns,omitempty"`
	Caching            *int32                `thrift:"caching,4" db:"caching" json:"caching,omitempty"`
	MaxVersions        int32                 `thrift:"maxVersions,5" db:"maxVersions" json:"maxVersions"`
	TimeRange          *TimeRange            `thrift:"timeRange,6" db:"timeRange" json:"timeRange,omitempty"`
	FilterString       []byte                `thrift:"filterString,7" db:"filterString" json:"filterString,omitempty"`
	BatchSize          *int32                `thrift:"batchSize,8" db:"batchSize" json:"batchSize,omitempty"`
	Attributes         map[string][]byte     `thrift:"attributes,9" db:"attributes" json:"attributes,omitempty"`
	Authorizations     *Authorization        `thrift:"authorizations,10" db:"authorizations" json:"authorizations,omitempty"`
	Reversed           *bool                 `thrift:"reversed,11" db:"reversed" json:"reversed,omitempty"`
	CacheBlocks        *bool                 `thrift:"cacheBlocks,12" db:"cacheBlocks" json:"cacheBlocks,omitempty"`
	ColFamTimeRangeMap map[string]*TimeRange `thrift:"colFamTimeRangeMap,13" db:"colFamTimeRangeMap" json:"colFamTimeRangeMap,omitempty"`
	ReadType           *ReadType             `thrift:"readType,14" db:"readType" json:"readType,omitempty"`
	Limit              *int32                `thrift:"limit,15" db:"limit" json:"limit,omitempty"`
	Consistency        *TConsistency         `thrift:"consistency,16" db:"consistency" json:"consistency,omitempty"`
	TargetReplicaId    *int32                `thrift:"targetReplicaId,17" db:"targetReplicaId" json:"targetReplicaId,omitempty"`
	FilterBytes        []byte                `thrift:"filterBytes,18" db:"filterBytes" json:"filterBytes,omitempty"`
}

func toThrift2TScan(ptr *Scan) *thrift2.TScan {
	return &thrift2.TScan{
		StartRow:           ptr.StartRow,
		StopRow:            ptr.StopRow,
		Columns:            toThrift2ColumnList(ptr.Columns),
		Caching:            ptr.Caching,
		MaxVersions:        ptr.MaxVersions,
		TimeRange:          toThrift2TTimeRange(ptr.TimeRange),
		FilterString:       ptr.FilterString,
		BatchSize:          ptr.BatchSize,
		Attributes:         ptr.Attributes,
		Authorizations:     toThrift2TAuthorization(ptr.Authorizations),
		Reversed:           ptr.Reversed,
		CacheBlocks:        ptr.CacheBlocks,
		ColFamTimeRangeMap: toThriftTTimeRangeMap(ptr.ColFamTimeRangeMap),
		ReadType:           (*thrift2.TReadType)(ptr.ReadType),
		Limit:              ptr.Limit,
		Consistency:        (*thrift2.TConsistency)(ptr.Consistency),
		TargetReplicaId:    ptr.TargetReplicaId,
		FilterBytes:        ptr.FilterBytes,
	}
}

type ReadType int64

const (
	DEFAULT ReadType = 1
	Stream  ReadType = 2
	PRead   ReadType = 3
)

type RowMutations struct {
	Row       []byte      `thrift:"row,1,required" db:"row" json:"row"`
	Mutations []*Mutation `thrift:"mutations,2,required" db:"mutations" json:"mutations"`
}

func toThrift2TRowMutations(ptr *RowMutations) *thrift2.TRowMutations {
	return &thrift2.TRowMutations{
		Row:       ptr.Row,
		Mutations: toThrift2MutationList(ptr.Mutations),
	}
}

type Mutation struct {
	Put          *Put
	DeleteSingle *Delete
}

func toThrift2TMutation(ptr *Mutation) *thrift2.TMutation {
	return &thrift2.TMutation{
		Put:          toThrift2TPut(ptr.Put),
		DeleteSingle: toThrift2TDelete(ptr.DeleteSingle),
	}
}

type HRegionLocation struct {
	ServerName *ServerName
	RegionInfo *HRegionInfo
}

func fromThrift2RegionLocation(ptr *thrift2.THRegionLocation) *HRegionLocation {
	if ptr == nil {
		return nil
	}
	return &HRegionLocation{
		ServerName: fromThrift2TServerName(ptr.ServerName),
		RegionInfo: fromThrift2HRegionInfo(ptr.RegionInfo),
	}
}

type ServerName struct {
	HostName  string
	Port      *int32
	StartCode *int64
}

func fromThrift2TServerName(ptr *thrift2.TServerName) *ServerName {
	if ptr == nil {
		return nil
	}
	return &ServerName{
		HostName:  ptr.HostName,
		Port:      ptr.Port,
		StartCode: ptr.StartCode,
	}
}

type HRegionInfo struct {
	RegionId  int64
	TableName []byte
	StartKey  []byte
	EndKey    []byte
	Offline   *bool
	Split     *bool
	ReplicaId *int32
}

func fromThrift2HRegionInfo(ptr *thrift2.THRegionInfo) *HRegionInfo {
	if ptr == nil {
		return nil
	}

	return &HRegionInfo{
		RegionId:  ptr.RegionId,
		TableName: ptr.TableName,
		StartKey:  ptr.StartKey,
		EndKey:    ptr.EndKey,
		Offline:   ptr.Offline,
		Split:     ptr.Split,
		ReplicaId: ptr.ReplicaId,
	}
}

type TCompareOp int64

const (
	CompareOpLESS           TCompareOp = 0
	CompareOpLESSOREQUAL    TCompareOp = 1
	CompareOpEQUAL          TCompareOp = 2
	CompareOpNOTEQUAL       TCompareOp = 3
	CompareOpGREATEROREQUAL TCompareOp = 4
	CompareOpGREATER        TCompareOp = 5
	CompareOpNOOP           TCompareOp = 6
)

type TableName struct {
	Ns        []byte `thrift:"ns,1" db:"ns" json:"ns,omitempty"`
	Qualifier []byte `thrift:"qualifier,2,required" db:"qualifier" json:"qualifier"`
}

func toThrift2TTableName(ptr *TableName) *thrift2.TTableName {
	if ptr == nil {
		return nil
	}
	return &thrift2.TTableName{
		Ns:        ptr.Ns,
		Qualifier: ptr.Qualifier,
	}
}

func fromThrift2TableName(ptr *thrift2.TTableName) *TableName {
	if ptr == nil {
		return nil
	}
	return &TableName{
		Ns:        ptr.Ns,
		Qualifier: ptr.Qualifier,
	}
}

type TableDescriptor struct {
	TableName  *TableName                `thrift:"tableName,1,required" db:"tableName" json:"tableName"`
	Columns    []*ColumnFamilyDescriptor `thrift:"columns,2" db:"columns" json:"columns,omitempty"`
	Attributes map[string][]byte         `thrift:"attributes,3" db:"attributes" json:"attributes,omitempty"`
	Durability *TDurability              `thrift:"durability,4" db:"durability" json:"durability,omitempty"`
}

func toThrift2TTableDescriptor(ptr *TableDescriptor) *thrift2.TTableDescriptor {
	if ptr == nil {
		return nil
	}
	return &thrift2.TTableDescriptor{
		TableName:  toThrift2TTableName(ptr.TableName),
		Columns:    toThrift2ColumnFamilyDescriptorList(ptr.Columns),
		Attributes: ptr.Attributes,
		Durability: (*thrift2.TDurability)(ptr.Durability),
	}
}

func fromThrift2TableDescriptor(ptr *thrift2.TTableDescriptor) *TableDescriptor {
	if ptr == nil {
		return nil
	}
	return &TableDescriptor{
		TableName:  fromThrift2TableName(ptr.TableName),
		Columns:    fromThrift2ColumnFamilyDescriptorList(ptr.Columns),
		Attributes: nil,
		Durability: nil,
	}
}

type ColumnFamilyDescriptor struct {
	Name                []byte
	Attributes          map[string][]byte
	Configuration       map[string]string
	BlockSize           *int32
	BloomnFilterType    *TBloomFilterType
	CompressionType     *TCompressionAlgorithm
	DfsReplication      *int16
	DataBlockEncoding   *TDataBlockEncoding
	KeepDeletedCells    *TKeepDeletedCells
	MaxVersions         *int32
	MinVersions         *int32
	Scope               *int32
	TimeToLive          *int32
	BlockCacheEnabled   *bool
	CacheBloomsOnWrite  *bool
	CacheDataOnWrite    *bool
	CacheIndexesOnWrite *bool
	CompressTags        *bool
	EvictBlocksOnClose  *bool
	InMemory            *bool
}

func toThrift2TColumnFamilyDescriptor(ptr *ColumnFamilyDescriptor) *thrift2.TColumnFamilyDescriptor {
	if ptr == nil {
		return nil
	}
	return &thrift2.TColumnFamilyDescriptor{
		Name:                ptr.Name,
		Attributes:          ptr.Attributes,
		Configuration:       ptr.Configuration,
		BlockSize:           ptr.BlockSize,
		BloomnFilterType:    (*thrift2.TBloomFilterType)(ptr.BloomnFilterType),
		CompressionType:     (*thrift2.TCompressionAlgorithm)(ptr.CompressionType),
		DfsReplication:      ptr.DfsReplication,
		DataBlockEncoding:   (*thrift2.TDataBlockEncoding)(ptr.DataBlockEncoding),
		KeepDeletedCells:    (*thrift2.TKeepDeletedCells)(ptr.KeepDeletedCells),
		MaxVersions:         ptr.MaxVersions,
		MinVersions:         ptr.MinVersions,
		Scope:               ptr.Scope,
		TimeToLive:          ptr.TimeToLive,
		BlockCacheEnabled:   ptr.BlockCacheEnabled,
		CacheBloomsOnWrite:  ptr.CacheBloomsOnWrite,
		CacheDataOnWrite:    ptr.CacheDataOnWrite,
		CacheIndexesOnWrite: ptr.CacheIndexesOnWrite,
		CompressTags:        ptr.CompressTags,
		EvictBlocksOnClose:  ptr.EvictBlocksOnClose,
		InMemory:            ptr.InMemory,
	}
}

func fromThrift2ColumnFamilyDescriptor(ptr *thrift2.TColumnFamilyDescriptor) *ColumnFamilyDescriptor {
	if ptr == nil {
		return nil
	}
	return &ColumnFamilyDescriptor{
		Name:                ptr.Name,
		Attributes:          ptr.Attributes,
		Configuration:       ptr.Configuration,
		BlockSize:           ptr.BlockSize,
		BloomnFilterType:    (*TBloomFilterType)(ptr.BloomnFilterType),
		CompressionType:     (*TCompressionAlgorithm)(ptr.CompressionType),
		DfsReplication:      ptr.DfsReplication,
		DataBlockEncoding:   (*TDataBlockEncoding)(ptr.DataBlockEncoding),
		KeepDeletedCells:    (*TKeepDeletedCells)(ptr.KeepDeletedCells),
		MaxVersions:         ptr.MaxVersions,
		MinVersions:         ptr.MinVersions,
		Scope:               ptr.Scope,
		TimeToLive:          ptr.TimeToLive,
		BlockCacheEnabled:   ptr.BlockCacheEnabled,
		CacheBloomsOnWrite:  ptr.CacheBloomsOnWrite,
		CacheDataOnWrite:    ptr.CacheDataOnWrite,
		CacheIndexesOnWrite: ptr.CacheIndexesOnWrite,
		CompressTags:        ptr.CompressTags,
		EvictBlocksOnClose:  ptr.EvictBlocksOnClose,
		InMemory:            ptr.InMemory,
	}
}

type TBloomFilterType int64

const (
	TBloomFilterType_NONE                   TBloomFilterType = 0
	TBloomFilterType_ROW                    TBloomFilterType = 1
	TBloomFilterType_ROWCOL                 TBloomFilterType = 2
	TBloomFilterType_ROWPREFIX_FIXED_LENGTH TBloomFilterType = 3
)

type TCompressionAlgorithm int64

const (
	TCompressionAlgorithm_LZO    TCompressionAlgorithm = 0
	TCompressionAlgorithm_GZ     TCompressionAlgorithm = 1
	TCompressionAlgorithm_NONE   TCompressionAlgorithm = 2
	TCompressionAlgorithm_SNAPPY TCompressionAlgorithm = 3
	TCompressionAlgorithm_LZ4    TCompressionAlgorithm = 4
	TCompressionAlgorithm_BZIP2  TCompressionAlgorithm = 5
	TCompressionAlgorithm_ZSTD   TCompressionAlgorithm = 6
)

type TDataBlockEncoding int64

func Thrift2TDataBlockEncoding(v *TDataBlockEncoding) *thrift2.TDataBlockEncoding {
	if v == nil {
		return nil
	}
	r := thrift2.TDataBlockEncoding(*v)
	return &r
}

const (
	TDataBlockEncoding_NONE         TDataBlockEncoding = 0
	TDataBlockEncoding_PREFIX       TDataBlockEncoding = 2
	TDataBlockEncoding_DIFF         TDataBlockEncoding = 3
	TDataBlockEncoding_FAST_DIFF    TDataBlockEncoding = 4
	TDataBlockEncoding_ROW_INDEX_V1 TDataBlockEncoding = 7
)

type TKeepDeletedCells int64

func Thrift2TKeepDeletedCells(v *TKeepDeletedCells) *thrift2.TKeepDeletedCells {
	if v == nil {
		return nil
	}
	r := thrift2.TKeepDeletedCells(*v)
	return &r
}

const (
	TKeepDeletedCells_FALSE TKeepDeletedCells = 0
	TKeepDeletedCells_TRUE  TKeepDeletedCells = 1
	TKeepDeletedCells_TTL   TKeepDeletedCells = 2
)

type NamespaceDescriptor struct {
	Name          string            `thrift:"name,1,required" db:"name" json:"name"`
	Configuration map[string]string `thrift:"configuration,2" db:"configuration" json:"configuration,omitempty"`
}

func toThrift2TNamespaceDescriptor(ptr *NamespaceDescriptor) *thrift2.TNamespaceDescriptor {
	if ptr == nil {
		return nil
	}
	return &thrift2.TNamespaceDescriptor{
		Name:          ptr.Name,
		Configuration: ptr.Configuration,
	}
}

func fromThrift2TNamespaceDescriptor(ptr *thrift2.TNamespaceDescriptor) *NamespaceDescriptor {
	if ptr == nil {
		return nil
	}
	return &NamespaceDescriptor{
		Name:          ptr.Name,
		Configuration: ptr.Configuration,
	}
}
