package json

import (
	"io"

	jsoniter "github.com/json-iterator/go"
)

func getJSONUtil() jsoniter.API {
	return jsoniter.Config{
		EscapeHTML:              true,
		SortMapKeys:             true,
		ValidateJsonRawMessage:  true,
		MarshalFloatWith6Digits: true,
	}.Froze()
}

// Marshal marshal function wrapper for jsoniter
func Marshal(v any) ([]byte, error) {
	return getJSONUtil().Marshal(v)
}

// MarshalIndent marshal indent function wrapper
func MarshalIndent(v any) ([]byte, error) {
	return getJSONUtil().MarshalIndent(v, "", " ")
}

// Unmarshal unmarshal function wrapper
func Unmarshal(data []byte, v any) error {
	return getJSONUtil().Unmarshal(data, v)
}

// Decode decode wrapper
func Decode(data io.Reader, v any) error {
	return jsoniter.NewDecoder(data).Decode(v)
}
