package nullable

// inspired by
// https://github.com/romanyx/nullable/blob/master/bool.go
//
// The Validate and ContextValidate methods below are needed by the
// code generated by go-swagger and so we create our own class

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/strfmt"
)

// Bool represents a boolean that may be null or not
// present in json at all.
type Bool struct {
	Present bool // Present is true if key is present in json
	Value   *bool
}

// NewBool creates a new nullable Bool
func NewBool(b bool) Bool {
	return Bool{
		Present: true,
		Value:   &b,
	}
}

// NewNullBool creates a null Bool
func NewNullBool() Bool {
	return Bool{
		Present: true,
		Value:   nil,
	}
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (s *Bool) UnmarshalJSON(data []byte) error {
	s.Present = true

	if bytes.Equal(data, null) {
		return nil
	}

	if err := json.Unmarshal(data, &s.Value); err != nil {
		return err
	}

	return nil
}

// Validate always validates ok
// this is called from the go swagger generated code
func (s *Bool) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate always validates ok
// this is called from the go swagger generated code
func (s *Bool) ContextValidate(context context.Context, formats strfmt.Registry) error {
	// always validate ok
	return nil
}