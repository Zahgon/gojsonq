package gojsonq

func abs(i int) int { _ = "STUB: not implemented"; return 0 }

func isIndex(in string) bool { _ = "STUB: not implemented"; return false }

func getIndex(in string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func toString(v interface{}) string { _ = "STUB: not implemented"; return "" }

// toFloat64 converts interface{} value to float64 if value is numeric else return false
func toFloat64(v interface{}) (float64, bool) {
	_ = "STUB: not implemented"
	return 0,

		// as Go convert the json Numeric value to float64
		false
}

// sortList sorts a list of interfaces
func sortList(list []interface{}, asc bool) []interface{} { _ = "STUB: not implemented"; return nil }

// sort elements for string

// sort elements for float64

type sortMap struct {
	data      interface{}
	key       string
	desc      bool
	separator string
	errs      []error
}

// Sort sorts the slice of maps
func (s *sortMap) Sort(data interface{}) { _ = "STUB: not implemented"; return }

// Len satisfies the sort.Interface
func (s *sortMap) Len() int { _ = "STUB: not implemented"; return 0 }

// Swap satisfies the sort.Interface
func (s *sortMap) Swap(i, j int) { _ = "STUB: not implemented"; return }

// TODO: need improvement
// Less satisfies the sort.Interface
// This will work for string/float64 only
func (s *sortMap) Less(i, j int) (res bool) { _ = "STUB: not implemented"; return false }

// compare nested values

// compare compare two values
func (s *sortMap) compare(x, y interface{}) (res bool) { _ = "STUB: not implemented"; return false }

// getNestedValue fetch nested value from node
func getNestedValue(input interface{}, node, separator string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// find slice/array

// find in map

// find in group data

// makeAlias provide syntactic suger. when provide Property name as "user.name as userName"
// it return userName as output and pure node name like: "user.name".
// If "user.name" does not use "as" clause then it'll return "user.name", "user.name"
func makeAlias(in, separator string) (string, string) { _ = "STUB: not implemented"; return "", "" }

// length return length of strings/array/map
func length(v interface{}) (int, error) { _ = "STUB: not implemented"; return 0, nil }
