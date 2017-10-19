// Code generated by optionalize
// on: 2017-10-19 10:05:31.799710448 &#43;0000 UTC
// from: github.com/Gurpartap/safer-go/optional/template/optional.go.tmpl

package optional

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type OptionalErrorSlice interface {
	// Take takes the value out of the option, leaving a None in its place.
	Take()

	// From performs the conversion.
	From(value Error)

	// Unwrap moves the value out of the optional, if it is Some(value).
	// This function returns multiple values, and if that's undesirable,
	// consider using Some and None functions.
	Unwrap() (Error, bool)

	// UnwrapOr returns the contained value or a default.
	UnwrapOr(def Error) Error

	// UnwrapOrElse returns the contained value or computes it from a closure.
	UnwrapOrElse(fn func() Error) Error

	// UnwrapOrPanic returns the contained value or panics.
	UnwrapOrPanic() Error

	// And returns None if the optional is None, otherwise returns optb.
	And(optb OptionalErrorSlice) OptionalErrorSlice

	// Or returns the optional if it contains a value, otherwise returns optb.
	Or(optb OptionalErrorSlice) OptionalErrorSlice

	// OrElse returns the optional if it contains a value, otherwise calls fn
	// and returns the result.
	OrElse(fn func() OptionalErrorSlice) OptionalErrorSlice

	// IsSome returns true if the optional is a Some value.
	IsSome() bool

	// IsNone returns true if the optional is a None value.
	IsNone() bool

	// Some executes the given closure if there is a Some value.
	Some(fn func(value Error)) someOptionalErrorSliceHandler

	// None executes the given closure if there is a None value.
	None(fn func())
}

type someOptionalErrorSliceHandler interface {
	None(fn func())
}

type _OptionalErrorSlice struct {
	hasValue    bool
	unsafeValue Error
}

func NewOptionalErrorSlice() OptionalErrorSlice {
	opt := &_OptionalErrorSlice{}
	opt.Take()
	return opt
}

func OptionalErrorSliceFrom(value Error, hasValue bool) OptionalErrorSlice {
	opt := &_OptionalErrorSlice{}
	if hasValue {
		opt.From(value)
	}
	return opt
}

// Take takes the value out of the optional, leaving a None in its place.
func (opt *_OptionalErrorSlice) Take() {
	var value Error
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *_OptionalErrorSlice) From(value Error) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *_OptionalErrorSlice) Unwrap() (Error, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *_OptionalErrorSlice) UnwrapOr(def Error) Error {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *_OptionalErrorSlice) UnwrapOrElse(fn func() Error) Error {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *_OptionalErrorSlice) UnwrapOrPanic() Error {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap OptionalErrorSlice")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *_OptionalErrorSlice) And(optb OptionalErrorSlice) OptionalErrorSlice {
	if !opt.getHasValue() {
		return opt
	}
	return optb
}

// Or returns the optional if it contains a value, otherwise returns optb.
func (opt *_OptionalErrorSlice) Or(optb OptionalErrorSlice) OptionalErrorSlice {
	if opt.getHasValue() {
		return opt
	}
	return optb
}

// OrElse returns the optional if it contains a value, otherwise calls fn and
// returns the result.
func (opt *_OptionalErrorSlice) OrElse(fn func() OptionalErrorSlice) OptionalErrorSlice {
	if opt.getHasValue() {
		return opt
	}
	return fn()
}

// IsSome returns true if the optional is a Some value.
func (opt *_OptionalErrorSlice) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the optional is a None value.
func (opt *_OptionalErrorSlice) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *_OptionalErrorSlice) Some(fn func(value Error)) someOptionalErrorSliceHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return _someOptionalErrorSliceHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *_OptionalErrorSlice) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type _someOptionalErrorSliceHandler struct {
	opt *_OptionalErrorSlice
}

func (some _someOptionalErrorSliceHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *_OptionalErrorSlice) getHasValue() bool {
	return opt.hasValue
}

func (opt *_OptionalErrorSlice) getUnsafeValue() Error {
	return opt.unsafeValue
}

// String conforms to fmt.Stringer interface.
func (opt *_OptionalErrorSlice) String() string {
	if value, ok := opt.Unwrap(); ok {
		return fmt.Sprintf("Some(%v)", value)
	}
	return "None"
}

// MarshalJSON implements the json.Marshaler interface.
func (opt _OptionalErrorSlice) MarshalJSON() ([]byte, error) {
	if !opt.getHasValue() {
		return []byte("null"), nil
	}
	return json.Marshal(opt.getUnsafeValue())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (opt *_OptionalErrorSlice) UnmarshalJSON(data []byte) error {
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
