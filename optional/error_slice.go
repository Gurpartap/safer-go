// Code generated by optionalize
// on: 2017-10-19 10:05:25.519802359 &#43;0000 UTC
// from: github.com/Gurpartap/safer-go/optional/template/optional.go.tmpl

package optional

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type ErrorSlice interface {
	// Take takes the value out of the option, leaving a None in its place.
	Take()

	// From performs the conversion.
	From(value []error)

	// Unwrap moves the value out of the optional, if it is Some(value).
	// This function returns multiple values, and if that's undesirable,
	// consider using Some and None functions.
	Unwrap() ([]error, bool)

	// UnwrapOr returns the contained value or a default.
	UnwrapOr(def []error) []error

	// UnwrapOrElse returns the contained value or computes it from a closure.
	UnwrapOrElse(fn func() []error) []error

	// UnwrapOrPanic returns the contained value or panics.
	UnwrapOrPanic() []error

	// And returns None if the optional is None, otherwise returns optb.
	And(optb ErrorSlice) ErrorSlice

	// Or returns the optional if it contains a value, otherwise returns optb.
	Or(optb ErrorSlice) ErrorSlice

	// OrElse returns the optional if it contains a value, otherwise calls fn
	// and returns the result.
	OrElse(fn func() ErrorSlice) ErrorSlice

	// IsSome returns true if the optional is a Some value.
	IsSome() bool

	// IsNone returns true if the optional is a None value.
	IsNone() bool

	// Some executes the given closure if there is a Some value.
	Some(fn func(value []error)) someErrorSliceHandler

	// None executes the given closure if there is a None value.
	None(fn func())
}

type someErrorSliceHandler interface {
	None(fn func())
}

type _ErrorSlice struct {
	hasValue    bool
	unsafeValue []error
}

func NewErrorSlice() ErrorSlice {
	opt := &_ErrorSlice{}
	opt.Take()
	return opt
}

func ErrorSliceFrom(value []error, hasValue bool) ErrorSlice {
	opt := &_ErrorSlice{}
	if hasValue {
		opt.From(value)
	}
	return opt
}

// Take takes the value out of the optional, leaving a None in its place.
func (opt *_ErrorSlice) Take() {
	var value []error
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *_ErrorSlice) From(value []error) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *_ErrorSlice) Unwrap() ([]error, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *_ErrorSlice) UnwrapOr(def []error) []error {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *_ErrorSlice) UnwrapOrElse(fn func() []error) []error {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *_ErrorSlice) UnwrapOrPanic() []error {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap ErrorSlice")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *_ErrorSlice) And(optb ErrorSlice) ErrorSlice {
	if !opt.getHasValue() {
		return opt
	}
	return optb
}

// Or returns the optional if it contains a value, otherwise returns optb.
func (opt *_ErrorSlice) Or(optb ErrorSlice) ErrorSlice {
	if opt.getHasValue() {
		return opt
	}
	return optb
}

// OrElse returns the optional if it contains a value, otherwise calls fn and
// returns the result.
func (opt *_ErrorSlice) OrElse(fn func() ErrorSlice) ErrorSlice {
	if opt.getHasValue() {
		return opt
	}
	return fn()
}

// IsSome returns true if the optional is a Some value.
func (opt *_ErrorSlice) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the optional is a None value.
func (opt *_ErrorSlice) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *_ErrorSlice) Some(fn func(value []error)) someErrorSliceHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return _someErrorSliceHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *_ErrorSlice) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type _someErrorSliceHandler struct {
	opt *_ErrorSlice
}

func (some _someErrorSliceHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *_ErrorSlice) getHasValue() bool {
	return opt.hasValue
}

func (opt *_ErrorSlice) getUnsafeValue() []error {
	return opt.unsafeValue
}

// String conforms to fmt.Stringer interface.
func (opt *_ErrorSlice) String() string {
	if value, ok := opt.Unwrap(); ok {
		return fmt.Sprintf("Some(%v)", value)
	}
	return "None"
}

// MarshalJSON implements the json.Marshaler interface.
func (opt _ErrorSlice) MarshalJSON() ([]byte, error) {
	if !opt.getHasValue() {
		return []byte("null"), nil
	}
	return json.Marshal(opt.getUnsafeValue())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (opt *_ErrorSlice) UnmarshalJSON(data []byte) error {
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
