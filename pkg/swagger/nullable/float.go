package nullable

// inspired by
// https://github.com/romanyx/nullable/blob/master/float.go
//
// The Validate and ContextValidate methods below are needed by the
// code generated by go-swagger and so we create our own class

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/strfmt"
)

// Float represents a float that may be null or not
// present in json at all.
type Float struct {
	Present bool // Present is true if key is present in json
	Value   *float64
}

// NewFloat creates a new nullable Float
func NewFloat(f float64) Float {
	return Float{
		Present: true,
		Value:   &f,
	}
}

// NewNullFloat creates a null Float
func NewNullFloat() Float {
	return Float{
		Present: true,
		Value:   nil,
	}
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (f *Float) UnmarshalJSON(data []byte) error {
	f.Present = true

	if bytes.Equal(data, null) {
		return nil
	}

	if err := json.Unmarshal(data, &f.Value); err != nil {
		return err
	}

	return nil
}

// Validate always validates ok
// this is called from the go swagger generated code
func (s *Float) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate always validates ok
// this is called from the go swagger generated code
func (s *Float) ContextValidate(context context.Context, formats strfmt.Registry) error {
	// always validate ok
	return nil
}