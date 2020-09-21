package hbase

import "fmt"

const (
	Binary       = ComparatorType("binary")
	BinaryPrefix = ComparatorType("binaryprefix")
	RegexString  = ComparatorType("regexstring")
	SubString    = ComparatorType("substring")
)

type Comparator interface {
	Value() string
}

//BinaryComparator: This lexicographically compares against the specified byte array using
//Bytes.compareTo(byte[], byte[])
type BinaryComparator struct {
	specified string
}

func NewBinaryComparator(specified string) *BinaryComparator {
	return &BinaryComparator{specified: specified}
}

func (c *BinaryComparator) Value() string {
	return fmt.Sprintf("binary:%s", c.specified)
}

//BinaryPrefixComparator: This lexicographically compares against a specified byte array. It only
//compares up to the length of this byte array
type BinaryPrefixComparator struct {
	prefix string
}

func NewBinaryPrefixComparator(prefix string) *BinaryPrefixComparator {
	return &BinaryPrefixComparator{prefix: prefix}
}

//RegexStringComparator: This compares against the specified byte array using the given regular
//expression. Only EQUAL and NOT_EQUAL comparisons are valid with this comparator
type RegexStringComparator struct {
	regex string
}

func NewRegexStringComparator(regex string) *RegexStringComparator {
	return &RegexStringComparator{regex: regex}
}

//SubStringComparator: This tests if the given substring appears in a specified byte array. The
//comparison is case insensitive. Only EQUAL and NOT_EQUAL comparisons are valid with this
//comparator
type SubStringComparator struct {
	sub string
}

func NewSubStringComparator(sub string) *SubStringComparator {
	return &SubStringComparator{sub: sub}
}
