// Code generated by optionalize
// on: 2017-10-19 12:55:59.999143993 &#43;0000 UTC
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

type UIntSlice struct {
	hasValue    bool
	unsafeValue []uint
}

func NewUIntSlice() UIntSlice {
	opt := &UIntSlice{}
	opt.Take()
	return *opt
}

func UIntSliceFrom(value []uint, hasValue bool) UIntSlice {
	opt := &UIntSlice{}
	if hasValue {
		opt.From(value)
	}
	return *opt
}

// Take takes the value out of the optional, leaving a None in its place.
func (opt *UIntSlice) Take() {
	var value []uint
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *UIntSlice) From(value []uint) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *UIntSlice) Unwrap() ([]uint, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *UIntSlice) UnwrapOr(def []uint) []uint {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *UIntSlice) UnwrapOrElse(fn func() []uint) []uint {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *UIntSlice) UnwrapOrPanic() []uint {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap UIntSlice")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *UIntSlice) And(optb UIntSlice) UIntSlice {
	if !opt.getHasValue() {
		return *opt
	}
	return optb
}

// Or returns the optional if it contains a value, otherwise returns optb.
func (opt *UIntSlice) Or(optb UIntSlice) UIntSlice {
	if opt.getHasValue() {
		return *opt
	}
	return optb
}

// OrElse returns the optional if it contains a value, otherwise calls fn and
// returns the result.
func (opt *UIntSlice) OrElse(fn func() UIntSlice) UIntSlice {
	if opt.getHasValue() {
		return *opt
	}
	return fn()
}

// IsSome returns true if the optional is a Some value.
func (opt *UIntSlice) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the optional is a None value.
func (opt *UIntSlice) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *UIntSlice) Some(fn func(value []uint)) someUIntSliceHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return someUIntSliceHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *UIntSlice) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type someUIntSliceHandler struct {
	opt *UIntSlice
}

func (some someUIntSliceHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *UIntSlice) getHasValue() bool {
	return opt.hasValue
}

func (opt *UIntSlice) getUnsafeValue() []uint {
	return opt.unsafeValue
}

// String conforms to fmt Stringer interface.
func (opt *UIntSlice) String() string {
	if value, ok := opt.Unwrap(); ok {
		return fmt.Sprintf("Some(%v)", value)
	}
	return "None"
}

// MarshalJSON implements the json Marshaler interface.
func (opt UIntSlice) MarshalJSON() ([]byte, error) {
	if !opt.getHasValue() {
		return []byte("null"), nil
	}
	return json.Marshal(opt.getUnsafeValue())
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (opt *UIntSlice) UnmarshalJSON(data []byte) error {
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
func (opt *UIntSlice) Scan(value interface{}) error {
	if value == nil {
		opt.Take()
		return nil
	}

	var unsafeValue []uint
	err := convert.ConvertAssign(&unsafeValue, value)
	if err != nil {
		return err
	}
	opt.From(unsafeValue)

	return nil
}

// Value implements the driver Valuer interface.
func (opt UIntSlice) Value() (driver.Value, error) {
	if !opt.getHasValue() {
		return nil, nil
	}
	return opt.getUnsafeValue(), nil
}
