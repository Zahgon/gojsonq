package gojsonq

import (
	"fmt"
	"time"
)

const errMessage = "gojsonq: wrong method call for %v"

// Available named error values
var (
	ErrExpectsPointer = fmt.Errorf("gojsonq: failed to unmarshal, expects pointer")
	ErrImmutable      = fmt.Errorf("gojsonq: failed to unmarshal, target is not mutable")
	ErrTypeMismatch   = fmt.Errorf("gojsonq: failed to unmarshal, target type misatched")
)

// NewResult return an instance of Result
func NewResult(v interface{}) *Result { _ = "STUB: not implemented"; return nil }

// Result represent custom type
type Result struct {
	value interface{}
}

// Nil check the query has result or not
func (r *Result) Nil() bool { _ = "STUB: not implemented"; return false }

// As sets the value of Result to v; It does not support methods with argument available in Result
func (r *Result) As(v interface{}) error { _ = "STUB: not implemented"; return nil }

// Bool assert the result to boolean value
func (r *Result) Bool() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Time assert the result to time.Time
func (r *Result) Time(layout string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// Duration assert the result to time.Duration
func (r *Result) Duration() (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// String assert the result to String
func (r *Result) String() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Int assert the result to int
func (r *Result) Int() (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Int8 assert the result to int8
func (r *Result) Int8() (int8, error) { _ = "STUB: not implemented"; return 0, nil }

// Int16 assert the result to int16
func (r *Result) Int16() (int16, error) { _ = "STUB: not implemented"; return 0, nil }

// Int32 assert the result to int32
func (r *Result) Int32() (int32, error) { _ = "STUB: not implemented"; return 0, nil }

// Int64 assert the result to int64
func (r *Result) Int64() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// Uint assert the result to uint
func (r *Result) Uint() (uint, error) { _ = "STUB: not implemented"; return 0, nil }

// Uint8 assert the result to uint8
func (r *Result) Uint8() (uint8, error) { _ = "STUB: not implemented"; return 0, nil }

// Uint16 assert the result to uint16
func (r *Result) Uint16() (uint16, error) { _ = "STUB: not implemented"; return 0, nil }

// Uint32 assert the result to uint32
func (r *Result) Uint32() (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

// Uint64 assert the result to uint64
func (r *Result) Uint64() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// Float32 assert the result to float32
func (r *Result) Float32() (float32, error) { _ = "STUB: not implemented"; return 0, nil }

// Float64 assert the result to 64
func (r *Result) Float64() (float64, error) { _ = "STUB: not implemented"; return 0, nil }

// TODO: Slice related methods - ideally they should return nil instead of empty struct
// in case of any error or no result. To keep compatibility with older version refactored
// to use make, in order to create slices.

// BoolSlice assert the result to []bool
func (r *Result) BoolSlice() ([]bool, error) { _ = "STUB: not implemented"; return nil, nil }

// TimeSlice assert the result to []time.Time
func (r *Result) TimeSlice(layout string) ([]time.Time, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DurationSlice assert the result to []time.Duration
func (r *Result) DurationSlice() ([]time.Duration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StringSlice assert the result to []string
func (r *Result) StringSlice() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// IntSlice assert the result to []int
func (r *Result) IntSlice() ([]int, error) { _ = "STUB: not implemented"; return nil, nil }

// Int8Slice assert the result to []int8
func (r *Result) Int8Slice() ([]int8, error) { _ = "STUB: not implemented"; return nil, nil }

// Int16Slice assert the result to []int16
func (r *Result) Int16Slice() ([]int16, error) { _ = "STUB: not implemented"; return nil, nil }

// Int32Slice assert the result to []int32
func (r *Result) Int32Slice() ([]int32, error) { _ = "STUB: not implemented"; return nil, nil }

// Int64Slice assert the result to []int64
func (r *Result) Int64Slice() ([]int64, error) { _ = "STUB: not implemented"; return nil, nil }

// UintSlice assert the result to []uint
func (r *Result) UintSlice() ([]uint, error) { _ = "STUB: not implemented"; return nil, nil }

// Uint8Slice assert the result to []uint8
func (r *Result) Uint8Slice() ([]uint8, error) { _ = "STUB: not implemented"; return nil, nil }

// Uint16Slice assert the result to []uint16
func (r *Result) Uint16Slice() ([]uint16, error) { _ = "STUB: not implemented"; return nil, nil }

// Uint32Slice assert the result to []uint32
func (r *Result) Uint32Slice() ([]uint32, error) { _ = "STUB: not implemented"; return nil, nil }

// Uint64Slice assert the result to []uint64
func (r *Result) Uint64Slice() ([]uint64, error) { _ = "STUB: not implemented"; return nil, nil }

// Float32Slice assert the result to []float32
func (r *Result) Float32Slice() ([]float32, error) { _ = "STUB: not implemented"; return nil, nil }

// Float64Slice assert the result to []float64
func (r *Result) Float64Slice() ([]float64, error) { _ = "STUB: not implemented"; return nil, nil }
