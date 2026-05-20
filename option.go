package gojsonq

// option describes type for providing configuration options to JSONQ
type option struct {
	decoder   Decoder
	separator string
}

// OptionFunc represents a contract for option func, it basically set options to jsonq instance options
type OptionFunc func(*JSONQ) error

// SetDecoder take a custom decoder to decode JSON
// Deprecated - use WithDecoder
func SetDecoder(u Decoder) OptionFunc {
	_ = "STUB: not implemented"
	return *

	// SetSeparator set custom separator for traversing child node, default separator is DOT (.)
	// Deprecated - use WithSeparator
	new(OptionFunc)
}

func SetSeparator(s string) OptionFunc {
	_ = "STUB: not implemented"
	return *

	// WithDecoder take a custom decoder to decode JSON
	new(OptionFunc)
}

func WithDecoder(u Decoder) OptionFunc { _ = "STUB: not implemented"; return *new(OptionFunc) }

// WithSeparator set custom separator for traversing child node, default separator is DOT (.)
func WithSeparator(s string) OptionFunc { _ = "STUB: not implemented"; return *new(OptionFunc) }
