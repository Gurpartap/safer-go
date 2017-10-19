// Code generated by optionalize
// on: 2017-10-19 12:53:21.856175435 &#43;0000 UTC
// from: github.com/Gurpartap/safer-go/optional/template/optional.go.tmpl

package optional

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type UInt struct {
	hasValue    bool
	unsafeValue uint
}

func NewUInt() UInt {
	opt := &UInt{}
	opt.Take()
	return *opt
}

func UIntFrom(value uint, hasValue bool) UInt {
	opt := &UInt{}
	if hasValue {
		opt.From(value)
	}
	return *opt
}

// Take takes the value out of the optional, leaving a None in its place.
func (opt *UInt) Take() {
	var value uint
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *UInt) From(value uint) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *UInt) Unwrap() (uint, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *UInt) UnwrapOr(def uint) uint {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *UInt) UnwrapOrElse(fn func() uint) uint {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *UInt) UnwrapOrPanic() uint {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap UInt")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *UInt) And(optb UInt) UInt {
	if !opt.getHasValue() {
		return *opt
	}
	return optb
}

// Or returns the optional if it contains a value, otherwise returns optb.
func (opt *UInt) Or(optb UInt) UInt {
	if opt.getHasValue() {
		return *opt
	}
	return optb
}

// OrElse returns the optional if it contains a value, otherwise calls fn and
// returns the result.
func (opt *UInt) OrElse(fn func() UInt) UInt {
	if opt.getHasValue() {
		return *opt
	}
	return fn()
}

// IsSome returns true if the optional is a Some value.
func (opt *UInt) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the optional is a None value.
func (opt *UInt) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *UInt) Some(fn func(value uint)) someUIntHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return someUIntHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *UInt) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type someUIntHandler struct {
	opt *UInt
}

func (some someUIntHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *UInt) getHasValue() bool {
	return opt.hasValue
}

func (opt *UInt) getUnsafeValue() uint {
	return opt.unsafeValue
}

// String conforms to fmt Stringer interface.
func (opt *UInt) String() string {
	if value, ok := opt.Unwrap(); ok {
		return fmt.Sprintf("Some(%v)", value)
	}
	return "None"
}

// MarshalJSON implements the json Marshaler interface.
func (opt UInt) MarshalJSON() ([]byte, error) {
	if !opt.getHasValue() {
		return []byte("null"), nil
	}
	return json.Marshal(opt.getUnsafeValue())
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (opt *UInt) UnmarshalJSON(data []byte) error {
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
