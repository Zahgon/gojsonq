package gojsonq

import (
	"encoding/json"
	"io"
)

// New returns a new instance of JSONQ
func New(options ...OptionFunc) *JSONQ { _ = "STUB: not implemented"; return nil }

// empty represents an empty result
var empty interface{}

const defaultSeparator = "."

// query describes a query
type query struct {
	key, operator string
	value         interface{}
}

// JSONQ describes a JSONQ type which contains all the state
type JSONQ struct {
	option           option               // contains options for JSONQ
	queryMap         map[string]QueryFunc // contains query functions
	node             string               // contains node name
	raw              json.RawMessage      // raw message from source (reader, string or file)
	rootJSONContent  interface{}          // original decoded json data
	jsonContent      interface{}          // copy of original decoded json data for further processing
	queryIndex       int                  // contains number of orWhere query call
	queries          [][]query            // nested queries
	attributes       []string             // select attributes that will be available in final resuls
	offsetRecords    int                  // number of records that will be skipped in final result
	limitRecords     int                  // number of records that will be available in final result
	distinctProperty string               // contain the distinct attribute name
	errors           []error              // contains all the errors when processing
}

// String satisfies stringer interface
func (j *JSONQ) String() string { _ = "STUB: not implemented"; return "" }

// decode decodes the raw message to Go data structure
func (j *JSONQ) decode() *JSONQ { _ = "STUB: not implemented"; return nil }

// Copy returns a new fresh instance of JSONQ with the original copy of data so that you can do
// concurrent operation on the same data without being decoded again
func (j *JSONQ) Copy() *JSONQ { _ = "STUB: not implemented"; return nil }

// File read the json content from physical file
func (j *JSONQ) File(filename string) *JSONQ { _ = "STUB: not implemented"; return nil }

// handle error

// JSONString reads the json content from valid json string
// Deprecated: this method will remove in next major release
func (j *JSONQ) JSONString(json string) *JSONQ { _ = "STUB: not implemented"; return nil }

// FromString reads the content from valid json/xml/csv/yml string
func (j *JSONQ) FromString(str string) *JSONQ { _ = "STUB: not implemented"; return nil }

// handle error

// Reader reads the json content from io reader
func (j *JSONQ) Reader(r io.Reader) *JSONQ { _ = "STUB: not implemented"; return nil }

// reset the buffer

// Error returns first occurred error
func (j *JSONQ) Error() error { _ = "STUB: not implemented"; return nil }

// Errors returns list of all errors
func (j *JSONQ) Errors() []error {
	_ = "STUB: not implemented"

	// addError adds error to error list
	return nil
}

func (j *JSONQ) addError(err error) *JSONQ { _ = "STUB: not implemented"; return nil }

// Macro adds a new query func to the JSONQ
func (j *JSONQ) Macro(operator string, fn QueryFunc) *JSONQ { _ = "STUB: not implemented"; return nil }

// From seeks the json content to provided node. e.g: "users.[0]"  or "users.[0].name"
func (j *JSONQ) From(node string) *JSONQ { _ = "STUB: not implemented"; return nil }

// FromInterface reads the content from valid map[string]interface{}
func (j *JSONQ) FromInterface(v interface{}) *JSONQ { _ = "STUB: not implemented"; return nil }

// Select use for selection of the properties from query result
func (j *JSONQ) Select(properties ...string) *JSONQ { _ = "STUB: not implemented"; return nil }

// Offset skips the number of records in result
func (j *JSONQ) Offset(offset int) *JSONQ { _ = "STUB: not implemented"; return nil }

// offset skips the records from result
func (j *JSONQ) offset() *JSONQ { _ = "STUB: not implemented"; return nil }

// Limit limits the number of records in result
func (j *JSONQ) Limit(limit int) *JSONQ { _ = "STUB: not implemented"; return nil }

// limit return the number of records in result set depending on the limit value
func (j *JSONQ) limit() *JSONQ { _ = "STUB: not implemented"; return nil }

// Where builds a where clause. e.g: Where("name", "contains", "doe")
func (j *JSONQ) Where(key, cond string, val interface{}) *JSONQ {
	_ = "STUB: not implemented"
	return nil
}

// WhereEqual is an alias of Where("key", "=", val)
func (j *JSONQ) WhereEqual(key string, val interface{}) *JSONQ {
	_ = "STUB: not implemented"
	return nil
}

// WhereNotEqual is an alias of Where("key", "!=", val)
func (j *JSONQ) WhereNotEqual(key string, val interface{}) *JSONQ {
	_ = "STUB: not implemented"
	return nil
}

// WhereNil is an alias of Where("key", "=", nil)
func (j *JSONQ) WhereNil(key string) *JSONQ { _ = "STUB: not implemented"; return nil }

// WhereNotNil is an alias of Where("key", "!=", nil)
func (j *JSONQ) WhereNotNil(key string) *JSONQ { _ = "STUB: not implemented"; return nil }

// WhereIn is an alias for where("key", "in", []string{"a", "b"})
func (j *JSONQ) WhereIn(key string, val interface{}) *JSONQ { _ = "STUB: not implemented"; return nil }

// WhereNotIn is an alias for where("key", "notIn", []string{"a", "b"})
func (j *JSONQ) WhereNotIn(key string, val interface{}) *JSONQ {
	_ = "STUB: not implemented"
	return nil
}

// OrWhere builds an OrWhere clause, basically it's a group of AND clauses
func (j *JSONQ) OrWhere(key, cond string, val interface{}) *JSONQ {
	_ = "STUB: not implemented"
	return nil
}

// WhereStartsWith satisfies Where clause which starts with provided value(string)
func (j *JSONQ) WhereStartsWith(key string, val interface{}) *JSONQ {
	_ = "STUB: not implemented"
	return nil
}

// WhereEndsWith satisfies Where clause which ends with provided value(string)
func (j *JSONQ) WhereEndsWith(key string, val interface{}) *JSONQ {
	_ = "STUB: not implemented"
	return nil
}

// WhereContains satisfies Where clause which contains provided value(string)
func (j *JSONQ) WhereContains(key string, val interface{}) *JSONQ {
	_ = "STUB: not implemented"
	return nil
}

// WhereStrictContains satisfies Where clause which contains provided value(string).
// This is case sensitive
func (j *JSONQ) WhereStrictContains(key string, val interface{}) *JSONQ {
	_ = "STUB: not implemented"
	return nil
}

// WhereLenEqual is an alias of Where("key", "leneq", val)
func (j *JSONQ) WhereLenEqual(key string, val interface{}) *JSONQ {
	_ = "STUB: not implemented"
	return nil
}

// WhereLenNotEqual is an alias of Where("key", "lenneq", val)
func (j *JSONQ) WhereLenNotEqual(key string, val interface{}) *JSONQ {
	_ = "STUB: not implemented"
	return nil
}

// findInArray traverses through a list and returns the value list.
// This helps to process Where/OrWhere queries
func (j *JSONQ) findInArray(aa []interface{}) []interface{} { _ = "STUB: not implemented"; return nil }

// findInMap traverses through a map and returns the matched value list.
// This helps to process Where/OrWhere queries
func (j *JSONQ) findInMap(vm map[string]interface{}) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

// processQuery makes the result
func (j *JSONQ) processQuery() *JSONQ { _ = "STUB: not implemented"; return nil }

// prepare builds the queries
func (j *JSONQ) prepare() *JSONQ { _ = "STUB: not implemented"; return nil }

// GroupBy builds a chunk of exact matched data in a group list using provided attribute/column/property
func (j *JSONQ) GroupBy(property string) *JSONQ { _ = "STUB: not implemented"; return nil }

// replace the new result with the previous result

// Sort sorts an array
// default ascending order, pass "desc" for descending order
func (j *JSONQ) Sort(order ...string) *JSONQ { _ = "STUB: not implemented"; return nil }

// SortBy sorts an array
// default ascending order, pass "desc" for descending order
func (j *JSONQ) SortBy(order ...string) *JSONQ { _ = "STUB: not implemented"; return nil }

// Distinct builds distinct value using provided attribute/column/property
func (j *JSONQ) Distinct(property string) *JSONQ { _ = "STUB: not implemented"; return nil }

// distinct builds distinct value using provided attribute/column/property
func (j *JSONQ) distinct() *JSONQ { _ = "STUB: not implemented"; return nil }

// replace the new result with the previous result

// sortBy sorts list of map
func (j *JSONQ) sortBy(property string, asc bool) *JSONQ { _ = "STUB: not implemented"; return nil }

// replace the new result with the previous result

// only return selected properties in result
func (j *JSONQ) only(properties ...string) interface{} { _ = "STUB: not implemented"; return nil }

// Only collects the properties from a list of object
func (j *JSONQ) Only(properties ...string) interface{} { _ = "STUB: not implemented"; return nil }

// OnlyR collects the properties from a list of object and return as Result instance
func (j *JSONQ) OnlyR(properties ...string) (*Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Pluck build an array of values form a property of a list of objects
func (j *JSONQ) Pluck(property string) interface{} { _ = "STUB: not implemented"; return nil }

// PluckR build an array of values form a property of a list of objects and return as Result instance
func (j *JSONQ) PluckR(property string) (*Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// reset resets the current state of JSONQ instance
func (j *JSONQ) reset() *JSONQ { _ = "STUB: not implemented"; return nil }

// Reset resets the current state of JSON instance and make a fresh object with the original json content
func (j *JSONQ) Reset() *JSONQ {
	_ = "STUB: not implemented"

	// Get return the result
	return nil
}

func (j *JSONQ) Get() interface{} { _ = "STUB: not implemented"; return nil }

// GetR return the query results as Result instance
func (j *JSONQ) GetR() (*Result, error) { _ = "STUB: not implemented"; return nil, nil }

// First returns the first element of a list
func (j *JSONQ) First() interface{} { _ = "STUB: not implemented"; return nil }

// FirstR returns the first element of a list as Result instance
func (j *JSONQ) FirstR() (*Result, error) { _ = "STUB: not implemented"; return nil, nil }

// Last returns the last element of a list
func (j *JSONQ) Last() interface{} { _ = "STUB: not implemented"; return nil }

// LastR returns the last element of a list as Result instance
func (j *JSONQ) LastR() (*Result, error) { _ = "STUB: not implemented"; return nil, nil }

// Nth returns the nth element of a list
func (j *JSONQ) Nth(index int) interface{} { _ = "STUB: not implemented"; return nil }

// NthR returns the nth element of a list as Result instance
func (j *JSONQ) NthR(index int) (*Result, error) { _ = "STUB: not implemented"; return nil, nil }

// Find returns the result of a exact matching path
func (j *JSONQ) Find(path string) interface{} { _ = "STUB: not implemented"; return nil }

// FindR returns the result as Result instance from the exact matching path
func (j *JSONQ) FindR(path string) (*Result, error) { _ = "STUB: not implemented"; return nil, nil }

// Count returns the number of total items.
// This could be a length of list/array/map
func (j *JSONQ) Count() int { _ = "STUB: not implemented"; return 0 }

// list of items

// return map len // TODO: need to think about map

// group data items

// Out write the queried data to defined custom type
func (j *JSONQ) Out(v interface{}) { _ = "STUB: not implemented"; return }

// Writer write the queried data to a io.Writer
func (j *JSONQ) Writer(w io.Writer) { _ = "STUB: not implemented"; return }

// More provides the functionality to query over the resultant data. See https://github.com/thedevsaddam/gojsonq/wiki/Queries#More
func (j *JSONQ) More() *JSONQ { _ = "STUB: not implemented"; return nil }

// getFloatValFromArray returns a list of float64 values from array/map for aggregation
func (j *JSONQ) getFloatValFromArray(arr []interface{}, property ...string) []float64 {
	_ = "STUB: not implemented"
	return nil
}

// getAggregationValues returns a list of float64 values for aggregation
func (j *JSONQ) getAggregationValues(property ...string) []float64 {
	_ = "STUB: not implemented"
	return nil
}

// Sum returns sum of values from array or from map using property
func (j *JSONQ) Sum(property ...string) float64 { _ = "STUB: not implemented"; return 0 }

// Avg returns average of values from array or from map using property
func (j *JSONQ) Avg(property ...string) float64 { _ = "STUB: not implemented"; return 0 }

// Min returns minimum value from array or from map using property
func (j *JSONQ) Min(property ...string) float64 { _ = "STUB: not implemented"; return 0 }

// Max returns maximum value from array or from map using property
func (j *JSONQ) Max(property ...string) float64 { _ = "STUB: not implemented"; return 0 }
