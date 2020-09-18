package hbase

type ColumnFamily struct {
	Name          []byte
	Attributes    map[string][]byte
	Configuration map[string]string
	BlockSize     *int32
	//BloomnFilterType    *TBloomFilterType
	//CompressionType     *TCompressionAlgorithm
	DfsReplication *int16
	//DataBlockEncoding   *TDataBlockEncoding
	//KeepDeletedCells    *TKeepDeletedCells
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
