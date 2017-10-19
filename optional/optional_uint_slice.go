// Code generated by optionalize
// on: 2017-10-19 10:36:07.23211406 &#43;0000 UTC
// from: github.com/Gurpartap/safer-go/optional/template/optional.go.tmpl

package optional

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type OptionalUIntSlice interface {
	// Take takes the value out of the option, leaving a None in its place.
	Take()

	// From performs the conversion.
	From(value UInt)

	// Unwrap moves the value out of the optional, if it is Some(value).
	// This function returns multiple values, and if that's undesirable,
	// consider using Some and None functions.
	Unwrap() (UInt, bool)

	// UnwrapOr returns the contained value or a default.
	UnwrapOr(def UInt) UInt

	// UnwrapOrElse returns the contained value or computes it from a closure.
	UnwrapOrElse(fn func() UInt) UInt

	// UnwrapOrPanic returns the contained value or panics.
	UnwrapOrPanic() UInt

	// And returns None if the optional is None, otherwise returns optb.
	And(optb OptionalUIntSlice) OptionalUIntSlice

	// Or returns the optional if it contains a value, otherwise returns optb.
	Or(optb OptionalUIntSlice) OptionalUIntSlice

	// OrElse returns the optional if it contains a value, otherwise calls fn
	// and returns the result.
	OrElse(fn func() OptionalUIntSlice) OptionalUIntSlice

	// IsSome returns true if the optional is a Some value.
	IsSome() bool

	// IsNone returns true if the optional is a None value.
	IsNone() bool

	// Some executes the given closure if there is a Some value.
	Some(fn func(value UInt)) someOptionalUIntSliceHandler

	// None executes the given closure if there is a None value.
	None(fn func())
}

type someOptionalUIntSliceHandler interface {
	None(fn func())
}

type _OptionalUIntSlice struct {
	hasValue    bool
	unsafeValue UInt
}

func NewOptionalUIntSlice() OptionalUIntSlice {
	opt := &_OptionalUIntSlice{}
	opt.Take()
	return opt
}

func OptionalUIntSliceFrom(value UInt, hasValue bool) OptionalUIntSlice {
	opt := &_OptionalUIntSlice{}
	if hasValue {
		opt.From(value)
	}
	return opt
}

// Take takes the value out of the optional, leaving a None in its place.
func (opt *_OptionalUIntSlice) Take() {
	var value UInt
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *_OptionalUIntSlice) From(value UInt) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *_OptionalUIntSlice) Unwrap() (UInt, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *_OptionalUIntSlice) UnwrapOr(def UInt) UInt {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *_OptionalUIntSlice) UnwrapOrElse(fn func() UInt) UInt {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *_OptionalUIntSlice) UnwrapOrPanic() UInt {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap OptionalUIntSlice")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *_OptionalUIntSlice) And(optb OptionalUIntSlice) OptionalUIntSlice {
	if !opt.getHasValue() {
		return opt
	}
	return optb
}

// Or returns the optional if it contains a value, otherwise returns optb.
func (opt *_OptionalUIntSlice) Or(optb OptionalUIntSlice) OptionalUIntSlice {
	if opt.getHasValue() {
		return opt
	}
	return optb
}

// OrElse returns the optional if it contains a value, otherwise calls fn and
// returns the result.
func (opt *_OptionalUIntSlice) OrElse(fn func() OptionalUIntSlice) OptionalUIntSlice {
	if opt.getHasValue() {
		return opt
	}
	return fn()
}

// IsSome returns true if the optional is a Some value.
func (opt *_OptionalUIntSlice) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the optional is a None value.
func (opt *_OptionalUIntSlice) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *_OptionalUIntSlice) Some(fn func(value UInt)) someOptionalUIntSliceHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return _someOptionalUIntSliceHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *_OptionalUIntSlice) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type _someOptionalUIntSliceHandler struct {
	opt *_OptionalUIntSlice
}

func (some _someOptionalUIntSliceHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *_OptionalUIntSlice) getHasValue() bool {
	return opt.hasValue
}

func (opt *_OptionalUIntSlice) getUnsafeValue() UInt {
	return opt.unsafeValue
}

// String conforms to fmt Stringer interface.
func (opt *_OptionalUIntSlice) String() string {
	if value, ok := opt.Unwrap(); ok {
		return fmt.Sprintf("Some(%v)", value)
	}
	return "None"
}

// MarshalJSON implements the json Marshaler interface.
func (opt _OptionalUIntSlice) MarshalJSON() ([]byte, error) {
	if !opt.getHasValue() {
		return []byte("null"), nil
	}
	return json.Marshal(opt.getUnsafeValue())
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (opt *_OptionalUIntSlice) UnmarshalJSON(data []byte) error {
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
