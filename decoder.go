package gojsonq

// Decoder provide contract to decode JSON using custom decoder
type Decoder interface {
	Decode(data []byte, v interface{}) error
}

// DefaultDecoder use json.Unmarshal to decode JSON
type DefaultDecoder struct{}

// Decode decodes using json.Unmarshal
func (u *DefaultDecoder) Decode(data []byte, v interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
