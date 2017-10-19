// Code generated by optionalize
// on: 2017-10-19 12:55:57.834638158 &#43;0000 UTC
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

type Int64Slice struct {
	hasValue    bool
	unsafeValue []int64
}

func NewInt64Slice() Int64Slice {
	opt := &Int64Slice{}
	opt.Take()
	return *opt
}

func Int64SliceFrom(value []int64, hasValue bool) Int64Slice {
	opt := &Int64Slice{}
	if hasValue {
		opt.From(value)
	}
	return *opt
}

// Take takes the value out of the optional, leaving a None in its place.
func (opt *Int64Slice) Take() {
	var value []int64
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *Int64Slice) From(value []int64) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *Int64Slice) Unwrap() ([]int64, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *Int64Slice) UnwrapOr(def []int64) []int64 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *Int64Slice) UnwrapOrElse(fn func() []int64) []int64 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *Int64Slice) UnwrapOrPanic() []int64 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap Int64Slice")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *Int64Slice) And(optb Int64Slice) Int64Slice {
	if !opt.getHasValue() {
		return *opt
	}
	return optb
}

// Or returns the optional if it contains a value, otherwise returns optb.
func (opt *Int64Slice) Or(optb Int64Slice) Int64Slice {
	if opt.getHasValue() {
		return *opt
	}
	return optb
}

// OrElse returns the optional if it contains a value, otherwise calls fn and
// returns the result.
func (opt *Int64Slice) OrElse(fn func() Int64Slice) Int64Slice {
	if opt.getHasValue() {
		return *opt
	}
	return fn()
}

// IsSome returns true if the optional is a Some value.
func (opt *Int64Slice) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the optional is a None value.
func (opt *Int64Slice) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *Int64Slice) Some(fn func(value []int64)) someInt64SliceHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return someInt64SliceHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *Int64Slice) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type someInt64SliceHandler struct {
	opt *Int64Slice
}

func (some someInt64SliceHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *Int64Slice) getHasValue() bool {
	return opt.hasValue
}

func (opt *Int64Slice) getUnsafeValue() []int64 {
	return opt.unsafeValue
}

// String conforms to fmt Stringer interface.
func (opt *Int64Slice) String() string {
	if value, ok := opt.Unwrap(); ok {
		return fmt.Sprintf("Some(%v)", value)
	}
	return "None"
}

// MarshalJSON implements the json Marshaler interface.
func (opt Int64Slice) MarshalJSON() ([]byte, error) {
	if !opt.getHasValue() {
		return []byte("null"), nil
	}
	return json.Marshal(opt.getUnsafeValue())
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (opt *Int64Slice) UnmarshalJSON(data []byte) error {
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
func (opt *Int64Slice) Scan(value interface{}) error {
	if value == nil {
		opt.Take()
		return nil
	}

	var unsafeValue []int64
	err := convert.ConvertAssign(&unsafeValue, value)
	if err != nil {
		return err
	}
	opt.From(unsafeValue)

	return nil
}

// Value implements the driver Valuer interface.
func (opt Int64Slice) Value() (driver.Value, error) {
	if !opt.getHasValue() {
		return nil, nil
	}
	return opt.getUnsafeValue(), nil
}
