module github.com/geange/hbase-go

go 1.15

require (
	git.apache.org/thrift.git v0.9.0
	github.com/apache/thrift v0.13.0
)

replace git.apache.org/thrift.git v0.9.0 => github.com/apache/thrift v0.0.0-20200208223004-33dcbd08c080
