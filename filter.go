package hbase

import "fmt"

type Filter interface {
	FilterExpress() string
}

type ComparatorType string

// Compare Operator
//1. LESS (<)
//2. LESS_OR_EQUAL (⇐)
//3. EQUAL (=)
//4. NOT_EQUAL (!=)
//5. GREATER_OR_EQUAL (>=)
//6. GREATER (>)
//7. NO_OP (no operation)
type CompareOp string

const (
	LESS             = CompareOp("<")
	LESS_OR_EQUAL    = CompareOp("<=")
	EQUAL            = CompareOp("=")
	NOT_EQUAL        = CompareOp("!=")
	GREATER_OR_EQUAL = CompareOp(">=")
	GREATER          = CompareOp(">")
	NO_OP            = CompareOp("")
)

type FilterType string

const (
	TypeKeyOnlyFilter                  = "KeyOnlyFilter"
	TypeFirstKeyOnlyFilter             = "FirstKeyOnlyFilter"
	TypePrefixFilter                   = "PrefixFilter"
	TypeColumnPrefixFilter             = "ColumnPrefixFilter"
	TypeMultipleColumnPrefixFilter     = "MultipleColumnPrefixFilter"
	TypeColumnCountGetFilter           = "ColumnCountGetFilter"
	TypePageFilter                     = "PageFilter"
	TypeColumnPaginationFilter         = "ColumnPaginationFilter"
	TypeInclusiveStopFilter            = "InclusiveStopFilter"
	TypeTimeStampsFilter               = "TimeStampsFilter"
	TypeRowFilter                      = "RowFilter"
	TypeFamilyFilter                   = "familyFilter"
	TypeQualifierFilter                = "QualifierFilter"
	TypeValueFilter                    = "ValueFilter"
	TypeDependentColumnFilter          = "DependentColumnFilter"
	TypeSingleColumnValueFilter        = "SingleColumnValueFilter"
	TypeSingleColumnValueExcludeFilter = "SingleColumnValueExcludeFilter"
	TypeColumnRangeFilter              = "ColumnRangeFilter"
)

type stringFilter string

func (f stringFilter) FilterExpress() string {
	return string(f)
}

//keyOnlyFilter: This filter doesn’t take any arguments. It returns only the key component of each key-value
type keyOnlyFilter struct{}

func (f keyOnlyFilter) FilterExpress() string {
	return fmt.Sprintf("%s ()", TypeKeyOnlyFilter)
}

func KeyOnlyFilter() Filter {
	return keyOnlyFilter{}
}

//firstKeyOnlyFilter: This filter doesn’t take any arguments. It returns only the first key-value from each row.
type firstKeyOnlyFilter struct{}

func (f firstKeyOnlyFilter) FilterExpress() string {
	return fmt.Sprintf("%s ()", TypeFirstKeyOnlyFilter)
}

func FirstKeyOnlyFilter() Filter {
	return firstKeyOnlyFilter{}
}

//prefixFilter: This filter takes one argument – a prefix of a row key. It returns only those key-values present in
//a row that starts with the specified row prefix
type prefixFilter struct {
	prefix string
}

func (f prefixFilter) FilterExpress() string {
	return fmt.Sprintf("%s (%s)", TypePrefixFilter, f.prefix)
}

func PrefixFilter(prefix string) Filter {
	return prefixFilter{prefix: prefix}
}

//columnPrefixFilter: This filter takes one argument – a column prefix. It returns only those key-values present in a
//column that starts with the specified column prefix. The column prefix must be of the form:
//“qualifier”.
type columnPrefixFilter struct {
	prefix string
}

func (f columnPrefixFilter) FilterExpress() string {
	return fmt.Sprintf("%s (%s)", TypeColumnPrefixFilter, f.prefix)
}

func ColumnPrefixFilter(prefix string) Filter {
	return columnPrefixFilter{prefix: prefix}
}

//multipleColumnPrefixFilter: This filter takes a list of column prefixes. It returns key-values that are present in a column that
//starts with any of the specified column prefixes. Each of the column prefixes must be of the
//form: “qualifier”.
type multipleColumnPrefixFilter struct {
	prefixes []string
}

func (f multipleColumnPrefixFilter) FilterExpress() string {
	return fmt.Sprintf("%s (%s)", TypeMultipleColumnPrefixFilter, stringSliceToString(f.prefixes))
}

func MultipleColumnPrefixFilter(prefixes []string) Filter {
	return multipleColumnPrefixFilter{prefixes: prefixes}
}

//columnCountGetFilter: This filter takes one argument – a limit. It returns the first limit number of columns in the table.
type columnCountGetFilter struct {
	limit uint64
}

func (f columnCountGetFilter) FilterExpress() string {
	return fmt.Sprintf("%s(%d)", TypeColumnCountGetFilter, f.limit)
}

func ColumnCountGetFilter(limit int) Filter {
	return columnCountGetFilter{limit: uint64(limit)}
}

//pageFilter: This filter takes one argument – a page size. It returns page size number of rows from the table.
type pageFilter struct {
	page uint64
}

func (f pageFilter) FilterExpress() string {
	return fmt.Sprintf("%s(%d)", TypePageFilter, f.page)
}

func PageFilter(page int) Filter {
	return pageFilter{page: uint64(page)}
}

//columnPaginationFilter: This filter takes two arguments – a limit and offset. It returns limit number of columns after
//offset number of columns. It does this for all the rows.
type columnPaginationFilter struct {
	limit  int
	offset int
}

func (f columnPaginationFilter) FilterExpress() string {
	return fmt.Sprintf("%s(%d, %d)", TypeColumnPaginationFilter, f.limit, f.offset)
}

//inclusiveStopFilter: This filter takes one argument – a row key on which to stop scanning. It returns all key-values
//present in rows up to and including the specified row.
type inclusiveStopFilter struct {
	rowKey string
}

func (f inclusiveStopFilter) FilterExpress() string {
	return fmt.Sprintf("%s('%s')", TypeInclusiveStopFilter, f.rowKey)
}

func InclusiveStopFilter(rowKey string) Filter {
	return inclusiveStopFilter{rowKey: rowKey}
}

//timeStampsFilter: This filter takes a list of timestamps. It returns those key-values whose timestamps matches any
//of the specified timestamps.
type timeStampsFilter struct {
	timeStamps []int64
}

func (f timeStampsFilter) FilterExpress() string {
	return fmt.Sprintf("%s(%s)", TypeTimeStampsFilter, int64SliceTostring(f.timeStamps))
}

func TimeStampsFilter(timeStamps []int64) Filter {
	return timeStampsFilter{timeStamps: timeStamps}
}

//rowFilter: This filter takes a compare operator and a comparator. It compares each row key with the
//comparator using the compare operator and if the comparison returns true, it returns all the
//key-values in that row.
type rowFilter struct {
	compareOp       CompareOp
	comparatorType  ComparatorType
	comparatorValue string
}

func (f rowFilter) FilterExpress() string {
	return fmt.Sprintf("%s(%s, '%s:%s')", TypeRowFilter, f.compareOp, f.comparatorType, f.comparatorValue)
}

func RowFilter(compareOp CompareOp, comparatorType ComparatorType, comparatorValue string) Filter {
	return rowFilter{
		compareOp:       compareOp,
		comparatorType:  comparatorType,
		comparatorValue: comparatorValue,
	}
}

//familyFilter: This filter takes a compare operator and a comparator. It compares each column family name
//with the comparator using the compare operator and if the comparison returns true, it returns
//all the Cells in that column family.
type familyFilter struct {
	compareOp       CompareOp
	comparatorType  ComparatorType
	comparatorValue string
}

func (f familyFilter) FilterExpress() string {
	return fmt.Sprintf("%s(%s, '%s:%s')", TypeFamilyFilter, f.compareOp, f.comparatorType, f.comparatorValue)
}

func FamilyFilter(compareOp CompareOp, comparatorType ComparatorType, comparatorValue string) Filter {
	return familyFilter{
		compareOp:       compareOp,
		comparatorType:  comparatorType,
		comparatorValue: comparatorValue,
	}
}

//qualifierFilter: This filter takes a compare operator and a comparator. It compares each qualifier name with the
//comparator using the compare operator and if the comparison returns true, it returns all the
//key-values in that column.
type qualifierFilter struct {
	compareOp       CompareOp
	comparatorType  ComparatorType
	comparatorValue string
}

func (f qualifierFilter) FilterExpress() string {
	return fmt.Sprintf("%s(%s, '%s:%s')", TypeQualifierFilter, f.compareOp, f.comparatorType, f.comparatorValue)
}

func QualifierFilter(compareOp CompareOp, comparatorType ComparatorType, comparatorValue string) Filter {
	return qualifierFilter{
		compareOp:       compareOp,
		comparatorType:  comparatorType,
		comparatorValue: comparatorValue,
	}
}

//valueFilter: This filter takes a compare operator and a comparator. It compares each value with the
//comparator using the compare operator and if the comparison returns true, it returns that keyvalue.
type valueFilter struct {
	compareOp       CompareOp
	comparatorType  ComparatorType
	comparatorValue string
}

func (f valueFilter) FilterExpress() string {
	return fmt.Sprintf("%s(%s, '%s:%s')", TypeValueFilter, f.compareOp, f.comparatorType, f.comparatorValue)
}

func ValueFilter(compareOp CompareOp, comparatorType ComparatorType, comparatorValue string) Filter {
	return valueFilter{
		compareOp:       compareOp,
		comparatorType:  comparatorType,
		comparatorValue: comparatorValue,
	}
}

//dependentColumnFilter: This filter takes two arguments – a family and a qualifier. It tries to locate this column in each
//row and returns all key-values in that row that have the same timestamp. If the row doesn’t
//contain the specified column – none of the key-values in that row will be returned.
type dependentColumnFilter struct {
	family    string
	qualifier string
}

func (f dependentColumnFilter) FilterExpress() string {
	return fmt.Sprintf("%s('%s', '%s')", TypeDependentColumnFilter, f.family, f.qualifier)
}

func DependentColumnFilter(family, qualifier string) Filter {
	return dependentColumnFilter{
		family:    family,
		qualifier: qualifier,
	}
}

//singleColumnValueFilter: This filter takes a column family, a qualifier, a compare operator and a comparator. If the
//specified column is not found – all the columns of that row will be emitted. If the column is
//found and the comparison with the comparator returns true, all the columns of the row will be
//emitted. If the condition fails, the row will not be emitted.
type singleColumnValueFilter struct {
	family          string
	qualifier       string
	compareOp       CompareOp
	comparatorType  ComparatorType
	comparatorValue string
}

func (f singleColumnValueFilter) FilterExpress() string {
	return fmt.Sprintf("%s('%s', '%s', %s, '%s:%s')", TypeSingleColumnValueFilter,
		f.family, f.qualifier, f.compareOp, f.comparatorType, f.comparatorValue)
}

func SingleColumnValueFilter(family, qualifier string,
	compareOp CompareOp, comparatorType ComparatorType, comparatorValue string) Filter {
	return singleColumnValueFilter{
		family:          family,
		qualifier:       qualifier,
		compareOp:       compareOp,
		comparatorType:  comparatorType,
		comparatorValue: comparatorValue,
	}
}

//singleColumnValueExcludeFilter: This filter takes the same arguments and behaves same as SingleColumnValueFilter –
//however, if the column is found and the condition passes, all the columns of the row will be emitted except
//for the tested column value.
type singleColumnValueExcludeFilter struct {
	family          string
	qualifier       string
	compareOp       CompareOp
	comparatorType  ComparatorType
	comparatorValue string
}

func (f singleColumnValueExcludeFilter) FilterExpress() string {
	return fmt.Sprintf("%s('%s', '%s', %s, '%s:%s')", TypeSingleColumnValueExcludeFilter,
		f.family, f.qualifier, f.compareOp, f.comparatorType, f.comparatorValue)
}

func SingleColumnValueExcludeFilter(family, qualifier string,
	compareOp CompareOp, comparatorType ComparatorType, comparatorValue string) Filter {
	return singleColumnValueExcludeFilter{
		family:          family,
		qualifier:       qualifier,
		compareOp:       compareOp,
		comparatorType:  comparatorType,
		comparatorValue: comparatorValue,
	}
}

//columnRangeFilter: This filter is used for selecting only those keys with columns that are between minColumn and
//maxColumn. It also takes two boolean variables to indicate whether to include the minColumn
//and maxColumn or not.
type columnRangeFilter struct {
	minColumn string
	minFlag   bool
	maxColumn string
	maxFlag   bool
}

func (f columnRangeFilter) FilterExpress() string {
	return fmt.Sprintf("%s('%s', %v, '%s', %v)", TypeColumnRangeFilter, f.minColumn, f.minFlag, f.maxColumn, f.maxFlag)
}

func ColumnRangeFilter(minColumn string, minFlag bool, maxColumn string, maxFlag bool) Filter {
	return columnRangeFilter{
		minColumn: minColumn,
		minFlag:   minFlag,
		maxColumn: maxColumn,
		maxFlag:   maxFlag,
	}
}

func ColumnRangeFilterV1(minColumn, maxColumn string) Filter {
	minFlag, maxFlag := false, false

	if minColumn != "" {
		minFlag = true
	}

	if maxColumn != "" {
		maxFlag = true
	}

	return columnRangeFilter{
		minColumn: minColumn,
		minFlag:   minFlag,
		maxColumn: maxColumn,
		maxFlag:   maxFlag,
	}
}
