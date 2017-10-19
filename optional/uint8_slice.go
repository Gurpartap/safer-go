// Code generated by optionalize
// on: 2017-10-19 10:05:19.795000385 &#43;0000 UTC
// from: github.com/Gurpartap/safer-go/optional/template/optional.go.tmpl

package optional

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type UInt8Slice interface {
	// Take takes the value out of the option, leaving a None in its place.
	Take()

	// From performs the conversion.
	From(value []uint8)

	// Unwrap moves the value out of the optional, if it is Some(value).
	// This function returns multiple values, and if that's undesirable,
	// consider using Some and None functions.
	Unwrap() ([]uint8, bool)

	// UnwrapOr returns the contained value or a default.
	UnwrapOr(def []uint8) []uint8

	// UnwrapOrElse returns the contained value or computes it from a closure.
	UnwrapOrElse(fn func() []uint8) []uint8

	// UnwrapOrPanic returns the contained value or panics.
	UnwrapOrPanic() []uint8

	// And returns None if the optional is None, otherwise returns optb.
	And(optb UInt8Slice) UInt8Slice

	// Or returns the optional if it contains a value, otherwise returns optb.
	Or(optb UInt8Slice) UInt8Slice

	// OrElse returns the optional if it contains a value, otherwise calls fn
	// and returns the result.
	OrElse(fn func() UInt8Slice) UInt8Slice

	// IsSome returns true if the optional is a Some value.
	IsSome() bool

	// IsNone returns true if the optional is a None value.
	IsNone() bool

	// Some executes the given closure if there is a Some value.
	Some(fn func(value []uint8)) someUInt8SliceHandler

	// None executes the given closure if there is a None value.
	None(fn func())
}

type someUInt8SliceHandler interface {
	None(fn func())
}

type _UInt8Slice struct {
	hasValue    bool
	unsafeValue []uint8
}

func NewUInt8Slice() UInt8Slice {
	opt := &_UInt8Slice{}
	opt.Take()
	return opt
}

func UInt8SliceFrom(value []uint8, hasValue bool) UInt8Slice {
	opt := &_UInt8Slice{}
	if hasValue {
		opt.From(value)
	}
	return opt
}

// Take takes the value out of the optional, leaving a None in its place.
func (opt *_UInt8Slice) Take() {
	var value []uint8
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *_UInt8Slice) From(value []uint8) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *_UInt8Slice) Unwrap() ([]uint8, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *_UInt8Slice) UnwrapOr(def []uint8) []uint8 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *_UInt8Slice) UnwrapOrElse(fn func() []uint8) []uint8 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *_UInt8Slice) UnwrapOrPanic() []uint8 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap UInt8Slice")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *_UInt8Slice) And(optb UInt8Slice) UInt8Slice {
	if !opt.getHasValue() {
		return opt
	}
	return optb
}

// Or returns the optional if it contains a value, otherwise returns optb.
func (opt *_UInt8Slice) Or(optb UInt8Slice) UInt8Slice {
	if opt.getHasValue() {
		return opt
	}
	return optb
}

// OrElse returns the optional if it contains a value, otherwise calls fn and
// returns the result.
func (opt *_UInt8Slice) OrElse(fn func() UInt8Slice) UInt8Slice {
	if opt.getHasValue() {
		return opt
	}
	return fn()
}

// IsSome returns true if the optional is a Some value.
func (opt *_UInt8Slice) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the optional is a None value.
func (opt *_UInt8Slice) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *_UInt8Slice) Some(fn func(value []uint8)) someUInt8SliceHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return _someUInt8SliceHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *_UInt8Slice) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type _someUInt8SliceHandler struct {
	opt *_UInt8Slice
}

func (some _someUInt8SliceHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *_UInt8Slice) getHasValue() bool {
	return opt.hasValue
}

func (opt *_UInt8Slice) getUnsafeValue() []uint8 {
	return opt.unsafeValue
}

// String conforms to fmt.Stringer interface.
func (opt *_UInt8Slice) String() string {
	if value, ok := opt.Unwrap(); ok {
		return fmt.Sprintf("Some(%v)", value)
	}
	return "None"
}

// MarshalJSON implements the json.Marshaler interface.
func (opt _UInt8Slice) MarshalJSON() ([]byte, error) {
	if !opt.getHasValue() {
		return []byte("null"), nil
	}
	return json.Marshal(opt.getUnsafeValue())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (opt *_UInt8Slice) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) || data == nil {
		opt.Take()
		return nil
	}

	err := json.Unmarshal(data, &opt.unsafeValue)
	if err != nil {
		opt.hasValue = false
		return err
	}
	opt.hasValue = true

	return nil
}
