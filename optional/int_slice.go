// Code generated by optionalize
// on: 2017-10-19 10:35:59.431746403 &#43;0000 UTC
// from: github.com/Gurpartap/safer-go/optional/template/optional.go.tmpl

package optional

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type IntSlice interface {
	// Take takes the value out of the option, leaving a None in its place.
	Take()

	// From performs the conversion.
	From(value []int)

	// Unwrap moves the value out of the optional, if it is Some(value).
	// This function returns multiple values, and if that's undesirable,
	// consider using Some and None functions.
	Unwrap() ([]int, bool)

	// UnwrapOr returns the contained value or a default.
	UnwrapOr(def []int) []int

	// UnwrapOrElse returns the contained value or computes it from a closure.
	UnwrapOrElse(fn func() []int) []int

	// UnwrapOrPanic returns the contained value or panics.
	UnwrapOrPanic() []int

	// And returns None if the optional is None, otherwise returns optb.
	And(optb IntSlice) IntSlice

	// Or returns the optional if it contains a value, otherwise returns optb.
	Or(optb IntSlice) IntSlice

	// OrElse returns the optional if it contains a value, otherwise calls fn
	// and returns the result.
	OrElse(fn func() IntSlice) IntSlice

	// IsSome returns true if the optional is a Some value.
	IsSome() bool

	// IsNone returns true if the optional is a None value.
	IsNone() bool

	// Some executes the given closure if there is a Some value.
	Some(fn func(value []int)) someIntSliceHandler

	// None executes the given closure if there is a None value.
	None(fn func())
}

type someIntSliceHandler interface {
	None(fn func())
}

type _IntSlice struct {
	hasValue    bool
	unsafeValue []int
}

func NewIntSlice() IntSlice {
	opt := &_IntSlice{}
	opt.Take()
	return opt
}

func IntSliceFrom(value []int, hasValue bool) IntSlice {
	opt := &_IntSlice{}
	if hasValue {
		opt.From(value)
	}
	return opt
}

// Take takes the value out of the optional, leaving a None in its place.
func (opt *_IntSlice) Take() {
	var value []int
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *_IntSlice) From(value []int) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *_IntSlice) Unwrap() ([]int, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *_IntSlice) UnwrapOr(def []int) []int {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *_IntSlice) UnwrapOrElse(fn func() []int) []int {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *_IntSlice) UnwrapOrPanic() []int {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap IntSlice")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *_IntSlice) And(optb IntSlice) IntSlice {
	if !opt.getHasValue() {
		return opt
	}
	return optb
}

// Or returns the optional if it contains a value, otherwise returns optb.
func (opt *_IntSlice) Or(optb IntSlice) IntSlice {
	if opt.getHasValue() {
		return opt
	}
	return optb
}

// OrElse returns the optional if it contains a value, otherwise calls fn and
// returns the result.
func (opt *_IntSlice) OrElse(fn func() IntSlice) IntSlice {
	if opt.getHasValue() {
		return opt
	}
	return fn()
}

// IsSome returns true if the optional is a Some value.
func (opt *_IntSlice) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the optional is a None value.
func (opt *_IntSlice) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *_IntSlice) Some(fn func(value []int)) someIntSliceHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return _someIntSliceHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *_IntSlice) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type _someIntSliceHandler struct {
	opt *_IntSlice
}

func (some _someIntSliceHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *_IntSlice) getHasValue() bool {
	return opt.hasValue
}

func (opt *_IntSlice) getUnsafeValue() []int {
	return opt.unsafeValue
}

// String conforms to fmt Stringer interface.
func (opt *_IntSlice) String() string {
	if value, ok := opt.Unwrap(); ok {
		return fmt.Sprintf("Some(%v)", value)
	}
	return "None"
}

// MarshalJSON implements the json Marshaler interface.
func (opt _IntSlice) MarshalJSON() ([]byte, error) {
	if !opt.getHasValue() {
		return []byte("null"), nil
	}
	return json.Marshal(opt.getUnsafeValue())
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (opt *_IntSlice) UnmarshalJSON(data []byte) error {
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
