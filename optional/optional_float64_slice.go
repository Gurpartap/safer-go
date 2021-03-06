// Code generated by optionalize
// on: 2017-10-19 12:56:04.581101591 &#43;0000 UTC
// from: github.com/Gurpartap/safer-go/optional/template/optional.go.tmpl

package optional

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/Gurpartap/safer-go/optional/internal/convert"
	"github.com/pkg/errors"
)

type OptionalFloat64Slice struct {
	hasValue    bool
	unsafeValue Float64
}

func NewOptionalFloat64Slice() OptionalFloat64Slice {
	opt := &OptionalFloat64Slice{}
	opt.Take()
	return *opt
}

func OptionalFloat64SliceFrom(value Float64, hasValue bool) OptionalFloat64Slice {
	opt := &OptionalFloat64Slice{}
	if hasValue {
		opt.From(value)
	}
	return *opt
}

// Take takes the value out of the optional, leaving a None in its place.
func (opt *OptionalFloat64Slice) Take() {
	var value Float64
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *OptionalFloat64Slice) From(value Float64) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *OptionalFloat64Slice) Unwrap() (Float64, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *OptionalFloat64Slice) UnwrapOr(def Float64) Float64 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *OptionalFloat64Slice) UnwrapOrElse(fn func() Float64) Float64 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *OptionalFloat64Slice) UnwrapOrPanic() Float64 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap OptionalFloat64Slice")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *OptionalFloat64Slice) And(optb OptionalFloat64Slice) OptionalFloat64Slice {
	if !opt.getHasValue() {
		return *opt
	}
	return optb
}

// Or returns the optional if it contains a value, otherwise returns optb.
func (opt *OptionalFloat64Slice) Or(optb OptionalFloat64Slice) OptionalFloat64Slice {
	if opt.getHasValue() {
		return *opt
	}
	return optb
}

// OrElse returns the optional if it contains a value, otherwise calls fn and
// returns the result.
func (opt *OptionalFloat64Slice) OrElse(fn func() OptionalFloat64Slice) OptionalFloat64Slice {
	if opt.getHasValue() {
		return *opt
	}
	return fn()
}

// IsSome returns true if the optional is a Some value.
func (opt *OptionalFloat64Slice) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the optional is a None value.
func (opt *OptionalFloat64Slice) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *OptionalFloat64Slice) Some(fn func(value Float64)) someOptionalFloat64SliceHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return someOptionalFloat64SliceHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *OptionalFloat64Slice) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type someOptionalFloat64SliceHandler struct {
	opt *OptionalFloat64Slice
}

func (some someOptionalFloat64SliceHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *OptionalFloat64Slice) getHasValue() bool {
	return opt.hasValue
}

func (opt *OptionalFloat64Slice) getUnsafeValue() Float64 {
	return opt.unsafeValue
}

// String conforms to fmt Stringer interface.
func (opt *OptionalFloat64Slice) String() string {
	if value, ok := opt.Unwrap(); ok {
		return fmt.Sprintf("Some(%v)", value)
	}
	return "None"
}

// MarshalJSON implements the json Marshaler interface.
func (opt OptionalFloat64Slice) MarshalJSON() ([]byte, error) {
	if !opt.getHasValue() {
		return []byte("null"), nil
	}
	return json.Marshal(opt.getUnsafeValue())
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (opt *OptionalFloat64Slice) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) || data == nil {
		opt.Take()
		return nil
	}

	err := json.Unmarshal(data, &opt.unsafeValue)
	if err != nil {
		opt.hasValue = false
		return errors.WithStack(err)
	}
	opt.hasValue = true

	return nil
}

// Scan implements the sql Scanner interface.
func (opt *OptionalFloat64Slice) Scan(value interface{}) error {
	if value == nil {
		opt.Take()
		return nil
	}

	var unsafeValue Float64
	err := convert.ConvertAssign(&unsafeValue, value)
	if err != nil {
		return err
	}
	opt.From(unsafeValue)

	return nil
}

// Value implements the driver Valuer interface.
func (opt OptionalFloat64Slice) Value() (driver.Value, error) {
	if !opt.getHasValue() {
		return nil, nil
	}
	return opt.getUnsafeValue(), nil
}
