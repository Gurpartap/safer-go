// Code generated by optionalize
// on: 2017-10-19 12:55:55.679183582 &#43;0000 UTC
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

type UInt8Slice struct {
	hasValue    bool
	unsafeValue []uint8
}

func NewUInt8Slice() UInt8Slice {
	opt := &UInt8Slice{}
	opt.Take()
	return *opt
}

func UInt8SliceFrom(value []uint8, hasValue bool) UInt8Slice {
	opt := &UInt8Slice{}
	if hasValue {
		opt.From(value)
	}
	return *opt
}

// Take takes the value out of the optional, leaving a None in its place.
func (opt *UInt8Slice) Take() {
	var value []uint8
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *UInt8Slice) From(value []uint8) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *UInt8Slice) Unwrap() ([]uint8, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *UInt8Slice) UnwrapOr(def []uint8) []uint8 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *UInt8Slice) UnwrapOrElse(fn func() []uint8) []uint8 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *UInt8Slice) UnwrapOrPanic() []uint8 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap UInt8Slice")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *UInt8Slice) And(optb UInt8Slice) UInt8Slice {
	if !opt.getHasValue() {
		return *opt
	}
	return optb
}

// Or returns the optional if it contains a value, otherwise returns optb.
func (opt *UInt8Slice) Or(optb UInt8Slice) UInt8Slice {
	if opt.getHasValue() {
		return *opt
	}
	return optb
}

// OrElse returns the optional if it contains a value, otherwise calls fn and
// returns the result.
func (opt *UInt8Slice) OrElse(fn func() UInt8Slice) UInt8Slice {
	if opt.getHasValue() {
		return *opt
	}
	return fn()
}

// IsSome returns true if the optional is a Some value.
func (opt *UInt8Slice) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the optional is a None value.
func (opt *UInt8Slice) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *UInt8Slice) Some(fn func(value []uint8)) someUInt8SliceHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return someUInt8SliceHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *UInt8Slice) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type someUInt8SliceHandler struct {
	opt *UInt8Slice
}

func (some someUInt8SliceHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *UInt8Slice) getHasValue() bool {
	return opt.hasValue
}

func (opt *UInt8Slice) getUnsafeValue() []uint8 {
	return opt.unsafeValue
}

// String conforms to fmt Stringer interface.
func (opt *UInt8Slice) String() string {
	if value, ok := opt.Unwrap(); ok {
		return fmt.Sprintf("Some(%v)", value)
	}
	return "None"
}

// MarshalJSON implements the json Marshaler interface.
func (opt UInt8Slice) MarshalJSON() ([]byte, error) {
	if !opt.getHasValue() {
		return []byte("null"), nil
	}
	return json.Marshal(opt.getUnsafeValue())
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (opt *UInt8Slice) UnmarshalJSON(data []byte) error {
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
func (opt *UInt8Slice) Scan(value interface{}) error {
	if value == nil {
		opt.Take()
		return nil
	}

	var unsafeValue []uint8
	err := convert.ConvertAssign(&unsafeValue, value)
	if err != nil {
		return err
	}
	opt.From(unsafeValue)

	return nil
}

// Value implements the driver Valuer interface.
func (opt UInt8Slice) Value() (driver.Value, error) {
	if !opt.getHasValue() {
		return nil, nil
	}
	return opt.getUnsafeValue(), nil
}
