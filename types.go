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

func (t *Get) Trans() *thrift2.TGet {
	return &thrift2.TGet{
		Row:             nil,
		Columns:         nil,
		Timestamp:       nil,
		TimeRange:       nil,
		MaxVersions:     nil,
		FilterString:    nil,
		Attributes:      nil,
		Authorizations:  nil,
		Consistency:     nil,
		TargetReplicaId: nil,
		CacheBlocks:     nil,
		StoreLimit:      nil,
		StoreOffset:     nil,
		ExistenceOnly:   nil,
		FilterBytes:     nil,
	}
}

type Column struct {
	Family    []byte
	Qualifier []byte
	Timestamp *int64
}

func (t *Column) Trans() *thrift2.TColumn {
	return &thrift2.TColumn{
		Family:    t.Family,
		Qualifier: t.Qualifier,
		Timestamp: t.Timestamp,
	}
}

type TimeRange struct {
	MinStamp int64
	MaxStamp int64
}

func (t *TimeRange) TransThrift2() *thrift2.TTimeRange {
	return &thrift2.TTimeRange{
		MinStamp: t.MinStamp,
		MaxStamp: t.MaxStamp,
	}
}

type Authorization struct {
	Labels []string
}

func (t *Authorization) TransThrift2() *thrift2.TAuthorization {
	return &thrift2.TAuthorization{Labels: t.Labels}
}

type TConsistency int64

type Result struct {
	Row          []byte
	ColumnValues []*ColumnValue
	Stale        bool
	Partial      bool
}

type ColumnValue struct {
	Family    []byte
	Qualifier []byte
	Value     []byte
	Timestamp *int64
	Tags      []byte
	Type      *int8
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

type TDurability int64

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

type Delete struct {
	Row        []byte     `thrift:"row,1,required" db:"row" json:"row"`
	Columns    []*Column  `thrift:"columns,2" db:"columns" json:"columns,omitempty"`
	Timestamp  *int64     `thrift:"timestamp,3" db:"timestamp" json:"timestamp,omitempty"`
	DeleteType DeleteType `thrift:"deleteType,4" db:"deleteType" json:"deleteType"`
	// unused field # 5
	Attributes map[string][]byte `thrift:"attributes,6" db:"attributes" json:"attributes,omitempty"`
	Durability *TDurability      `thrift:"durability,7" db:"durability" json:"durability,omitempty"`
}

type DeleteType int64

const (
	DeleteColumn        DeleteType = 0
	DeleteColumns       DeleteType = 1
	DeleteFamily        DeleteType = 2
	DeleteFamilyVersion DeleteType = 3
)

type Increment struct {
	Row            []byte             `thrift:"row,1,required" db:"row" json:"row"`
	Columns        []*ColumnIncrement `thrift:"columns,2,required" db:"columns" json:"columns"`
	Attributes     map[string][]byte  `thrift:"attributes,4" db:"attributes" json:"attributes,omitempty"`
	Durability     *TDurability       `thrift:"durability,5" db:"durability" json:"durability,omitempty"`
	CellVisibility *CellVisibility    `thrift:"cellVisibility,6" db:"cellVisibility" json:"cellVisibility,omitempty"`
	ReturnResults  *bool              `thrift:"returnResults,7" db:"returnResults" json:"returnResults,omitempty"`
}

type ColumnIncrement struct {
	Family    []byte `thrift:"family,1,required" db:"family" json:"family"`
	Qualifier []byte `thrift:"qualifier,2,required" db:"qualifier" json:"qualifier"`
	Amount    int64  `thrift:"amount,3" db:"amount" json:"amount"`
}

type TAppend struct {
	Row            []byte
	Columns        []*ColumnValue
	Attributes     map[string][]byte
	Durability     *TDurability
	CellVisibility *CellVisibility
	ReturnResults  *bool
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

// todo
func (t *Scan) TransThrift2() *thrift2.TScan {
	return &thrift2.TScan{
		StartRow:           t.StartRow,
		StopRow:            t.StopRow,
		Columns:            nil,
		Caching:            nil,
		MaxVersions:        0,
		TimeRange:          nil,
		FilterString:       nil,
		BatchSize:          nil,
		Attributes:         nil,
		Authorizations:     nil,
		Reversed:           nil,
		CacheBlocks:        nil,
		ColFamTimeRangeMap: nil,
		ReadType:           nil,
		Limit:              nil,
		Consistency:        nil,
		TargetReplicaId:    nil,
		FilterBytes:        nil,
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

// todo
func (t *RowMutations) TransThrift2() *thrift2.TRowMutations {
	return &thrift2.TRowMutations{
		Row:       t.Row,
		Mutations: nil,
	}
}

type Mutation struct {
	Put          *Put    `thrift:"put,1" db:"put" json:"put,omitempty"`
	DeleteSingle *Delete `thrift:"deleteSingle,2" db:"deleteSingle" json:"deleteSingle,omitempty"`
}

// todo
func (t *Mutation) TransThrift2() *thrift2.TMutation {
	return &thrift2.TMutation{
		Put:          nil,
		DeleteSingle: nil,
	}
}

type THRegionLocation struct {
	ServerName *ServerName  `thrift:"serverName,1,required" db:"serverName" json:"serverName"`
	RegionInfo *HRegionInfo `thrift:"regionInfo,2,required" db:"regionInfo" json:"regionInfo"`
}

type ServerName struct {
	HostName  string `thrift:"hostName,1,required" db:"hostName" json:"hostName"`
	Port      *int32 `thrift:"port,2" db:"port" json:"port,omitempty"`
	StartCode *int64 `thrift:"startCode,3" db:"startCode" json:"startCode,omitempty"`
}

type HRegionInfo struct {
	RegionId  int64  `thrift:"regionId,1,required" db:"regionId" json:"regionId"`
	TableName []byte `thrift:"tableName,2,required" db:"tableName" json:"tableName"`
	StartKey  []byte `thrift:"startKey,3" db:"startKey" json:"startKey,omitempty"`
	EndKey    []byte `thrift:"endKey,4" db:"endKey" json:"endKey,omitempty"`
	Offline   *bool  `thrift:"offline,5" db:"offline" json:"offline,omitempty"`
	Split     *bool  `thrift:"split,6" db:"split" json:"split,omitempty"`
	ReplicaId *int32 `thrift:"replicaId,7" db:"replicaId" json:"replicaId,omitempty"`
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

func (t *TableName) TransThrift2() *thrift2.TTableName {
	return &thrift2.TTableName{
		Ns:        t.Ns,
		Qualifier: t.Qualifier,
	}
}

type TableDescriptor struct {
	TableName  *TableName                `thrift:"tableName,1,required" db:"tableName" json:"tableName"`
	Columns    []*ColumnFamilyDescriptor `thrift:"columns,2" db:"columns" json:"columns,omitempty"`
	Attributes map[string][]byte         `thrift:"attributes,3" db:"attributes" json:"attributes,omitempty"`
	Durability *TDurability              `thrift:"durability,4" db:"durability" json:"durability,omitempty"`
}

//todo
func (t *TableDescriptor) TransThrift2() *thrift2.TTableDescriptor {
	return &thrift2.TTableDescriptor{
		TableName:  t.TableName.TransThrift2(),
		Columns:    nil,
		Attributes: t.Attributes,
		Durability: nil,
	}
}

type ColumnFamilyDescriptor struct {
	Name                []byte                 `thrift:"name,1,required" db:"name" json:"name"`
	Attributes          map[string][]byte      `thrift:"attributes,2" db:"attributes" json:"attributes,omitempty"`
	Configuration       map[string]string      `thrift:"configuration,3" db:"configuration" json:"configuration,omitempty"`
	BlockSize           *int32                 `thrift:"blockSize,4" db:"blockSize" json:"blockSize,omitempty"`
	BloomnFilterType    *TBloomFilterType      `thrift:"bloomnFilterType,5" db:"bloomnFilterType" json:"bloomnFilterType,omitempty"`
	CompressionType     *TCompressionAlgorithm `thrift:"compressionType,6" db:"compressionType" json:"compressionType,omitempty"`
	DfsReplication      *int16                 `thrift:"dfsReplication,7" db:"dfsReplication" json:"dfsReplication,omitempty"`
	DataBlockEncoding   *TDataBlockEncoding    `thrift:"dataBlockEncoding,8" db:"dataBlockEncoding" json:"dataBlockEncoding,omitempty"`
	KeepDeletedCells    *TKeepDeletedCells     `thrift:"keepDeletedCells,9" db:"keepDeletedCells" json:"keepDeletedCells,omitempty"`
	MaxVersions         *int32                 `thrift:"maxVersions,10" db:"maxVersions" json:"maxVersions,omitempty"`
	MinVersions         *int32                 `thrift:"minVersions,11" db:"minVersions" json:"minVersions,omitempty"`
	Scope               *int32                 `thrift:"scope,12" db:"scope" json:"scope,omitempty"`
	TimeToLive          *int32                 `thrift:"timeToLive,13" db:"timeToLive" json:"timeToLive,omitempty"`
	BlockCacheEnabled   *bool                  `thrift:"blockCacheEnabled,14" db:"blockCacheEnabled" json:"blockCacheEnabled,omitempty"`
	CacheBloomsOnWrite  *bool                  `thrift:"cacheBloomsOnWrite,15" db:"cacheBloomsOnWrite" json:"cacheBloomsOnWrite,omitempty"`
	CacheDataOnWrite    *bool                  `thrift:"cacheDataOnWrite,16" db:"cacheDataOnWrite" json:"cacheDataOnWrite,omitempty"`
	CacheIndexesOnWrite *bool                  `thrift:"cacheIndexesOnWrite,17" db:"cacheIndexesOnWrite" json:"cacheIndexesOnWrite,omitempty"`
	CompressTags        *bool                  `thrift:"compressTags,18" db:"compressTags" json:"compressTags,omitempty"`
	EvictBlocksOnClose  *bool                  `thrift:"evictBlocksOnClose,19" db:"evictBlocksOnClose" json:"evictBlocksOnClose,omitempty"`
	InMemory            *bool                  `thrift:"inMemory,20" db:"inMemory" json:"inMemory,omitempty"`
}

//todo
func (t *ColumnFamilyDescriptor) TransThrift2() *thrift2.TColumnFamilyDescriptor {
	return &thrift2.TColumnFamilyDescriptor{
		Name:                t.Name,
		Attributes:          t.Attributes,
		Configuration:       t.Configuration,
		BlockSize:           t.BlockSize,
		BloomnFilterType:    nil,
		CompressionType:     nil,
		DfsReplication:      t.DfsReplication,
		DataBlockEncoding:   nil,
		KeepDeletedCells:    nil,
		MaxVersions:         t.MaxVersions,
		MinVersions:         t.MinVersions,
		Scope:               t.Scope,
		TimeToLive:          t.TimeToLive,
		BlockCacheEnabled:   t.BlockCacheEnabled,
		CacheBloomsOnWrite:  t.CacheBloomsOnWrite,
		CacheDataOnWrite:    t.CacheDataOnWrite,
		CacheIndexesOnWrite: t.CacheIndexesOnWrite,
		CompressTags:        t.CompressTags,
		EvictBlocksOnClose:  t.EvictBlocksOnClose,
		InMemory:            t.InMemory,
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

const (
	TDataBlockEncoding_NONE         TDataBlockEncoding = 0
	TDataBlockEncoding_PREFIX       TDataBlockEncoding = 2
	TDataBlockEncoding_DIFF         TDataBlockEncoding = 3
	TDataBlockEncoding_FAST_DIFF    TDataBlockEncoding = 4
	TDataBlockEncoding_ROW_INDEX_V1 TDataBlockEncoding = 7
)

type TKeepDeletedCells int64

const (
	TKeepDeletedCells_FALSE TKeepDeletedCells = 0
	TKeepDeletedCells_TRUE  TKeepDeletedCells = 1
	TKeepDeletedCells_TTL   TKeepDeletedCells = 2
)

type NamespaceDescriptor struct {
	Name          string            `thrift:"name,1,required" db:"name" json:"name"`
	Configuration map[string]string `thrift:"configuration,2" db:"configuration" json:"configuration,omitempty"`
}

func (t *NamespaceDescriptor) TransThrift2() *thrift2.TNamespaceDescriptor {
	return &thrift2.TNamespaceDescriptor{
		Name:          t.Name,
		Configuration: t.Configuration,
	}
}
