package gojsonq

const (
	operatorEq             = "="
	operatorEqEng          = "eq"
	operatorNotEq          = "!="
	operatorNotEqEng       = "neq"
	operatorNotEqAnother   = "<>"
	operatorGt             = ">"
	operatorGtEng          = "gt"
	operatorLt             = "<"
	operatorLtEng          = "lt"
	operatorGtE            = ">="
	operatorGtEEng         = "gte"
	operatorLtE            = "<="
	operatorLtEEng         = "lte"
	operatorStrictContains = "strictContains"
	operatorContains       = "contains"
	operatorEndsWith       = "endsWith"
	operatorStartsWith     = "startsWith"
	operatorIn             = "in"
	operatorNotIn          = "notIn"
	operatorLenEq          = "leneq"
	operatorLenNotEq       = "lenneq"
	operatorLenGt          = "lengt"
	operatorLenGte         = "lengte"
	operatorLenLt          = "lenlt"
	operatorLenLte         = "lenlte"
)

func defaultQueries() map[string]QueryFunc { _ = "STUB: not implemented"; return nil }

// QueryFunc describes a conditional function which perform comparison
type QueryFunc func(x, y interface{}) (bool, error)

// eq checks whether x, y are deeply eq
func eq(x, y interface{}) (bool, error) {
	_ = "STUB: not implemented"
	// if the y value is numeric (int/int8-int64/float32/float64) then convert to float64
	return false, nil
}

// neq checks whether x, y are deeply not equal
func neq(x, y interface{}) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// gt checks whether x is greather than y
func gt(x, y interface{}) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// if the y value is numeric (int/int8-int64/float32/float64) then convert to float64

// lt checks whether x is less than y
func lt(x, y interface{}) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// if the y value is numeric (int/int8-int64/float32/float64) then convert to float64

// gte checks whether x is greater than or equal to y
func gte(x, y interface{}) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// if the y value is numeric (int/int8-int64/float32/float64) then convert to float64

// lte checks whether x is less than or equal to y
func lte(x, y interface{}) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// if the y value is numeric (int/int8-int64/float32/float64) then convert to float64

// strStrictContains checks if x contains y
// This is case sensitive search
func strStrictContains(x, y interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// strContains checks if x contains y
// This is case insensitive search
func strContains(x, y interface{}) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// strStartsWith checks if x starts with y
func strStartsWith(x, y interface{}) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// strEndsWith checks if x ends with y
func strEndsWith(x, y interface{}) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// in checks if x exists in y e.g: in("id", []int{1,3,5,8})
func in(x, y interface{}) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// notIn checks if x doesn't exists in y e.g: in("id", []int{1,3,5,8})
func notIn(x, y interface{}) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// lenEq checks if the string/array/list value is equal
func lenEq(x, y interface{}) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// lenNotEq checks if the string/array/list value is not equal
func lenNotEq(x, y interface{}) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// lenGt checks if the string/array/list value is greater
func lenGt(x, y interface{}) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// lenLt checks if the string/array/list value is less
func lenLt(x, y interface{}) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// lenGte checks if the string/array/list value is greater than equal
func lenGte(x, y interface{}) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// lenLte checks if the string/array/list value is less than equal
func lenLte(x, y interface{}) (bool, error) { _ = "STUB: not implemented"; return false, nil }
