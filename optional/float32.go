// Code generated by optionalize
// on: 2017-10-19 10:35:51.339412506 &#43;0000 UTC
// from: github.com/Gurpartap/safer-go/optional/template/optional.go.tmpl

package optional

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type Float32 interface {
	// Take takes the value out of the option, leaving a None in its place.
	Take()

	// From performs the conversion.
	From(value float32)

	// Unwrap moves the value out of the optional, if it is Some(value).
	// This function returns multiple values, and if that's undesirable,
	// consider using Some and None functions.
	Unwrap() (float32, bool)

	// UnwrapOr returns the contained value or a default.
	UnwrapOr(def float32) float32

	// UnwrapOrElse returns the contained value or computes it from a closure.
	UnwrapOrElse(fn func() float32) float32

	// UnwrapOrPanic returns the contained value or panics.
	UnwrapOrPanic() float32

	// And returns None if the optional is None, otherwise returns optb.
	And(optb Float32) Float32

	// Or returns the optional if it contains a value, otherwise returns optb.
	Or(optb Float32) Float32

	// OrElse returns the optional if it contains a value, otherwise calls fn
	// and returns the result.
	OrElse(fn func() Float32) Float32

	// IsSome returns true if the optional is a Some value.
	IsSome() bool

	// IsNone returns true if the optional is a None value.
	IsNone() bool

	// Some executes the given closure if there is a Some value.
	Some(fn func(value float32)) someFloat32Handler

	// None executes the given closure if there is a None value.
	None(fn func())
}

type someFloat32Handler interface {
	None(fn func())
}

type _Float32 struct {
	hasValue    bool
	unsafeValue float32
}

func NewFloat32() Float32 {
	opt := &_Float32{}
	opt.Take()
	return opt
}

func Float32From(value float32, hasValue bool) Float32 {
	opt := &_Float32{}
	if hasValue {
		opt.From(value)
	}
	return opt
}

// Take takes the value out of the optional, leaving a None in its place.
func (opt *_Float32) Take() {
	var value float32
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *_Float32) From(value float32) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *_Float32) Unwrap() (float32, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *_Float32) UnwrapOr(def float32) float32 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *_Float32) UnwrapOrElse(fn func() float32) float32 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *_Float32) UnwrapOrPanic() float32 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap Float32")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *_Float32) And(optb Float32) Float32 {
	if !opt.getHasValue() {
		return opt
	}
	return optb
}

// Or returns the optional if it contains a value, otherwise returns optb.
func (opt *_Float32) Or(optb Float32) Float32 {
	if opt.getHasValue() {
		return opt
	}
	return optb
}

// OrElse returns the optional if it contains a value, otherwise calls fn and
// returns the result.
func (opt *_Float32) OrElse(fn func() Float32) Float32 {
	if opt.getHasValue() {
		return opt
	}
	return fn()
}

// IsSome returns true if the optional is a Some value.
func (opt *_Float32) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the optional is a None value.
func (opt *_Float32) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *_Float32) Some(fn func(value float32)) someFloat32Handler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return _someFloat32Handler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *_Float32) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type _someFloat32Handler struct {
	opt *_Float32
}

func (some _someFloat32Handler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *_Float32) getHasValue() bool {
	return opt.hasValue
}

func (opt *_Float32) getUnsafeValue() float32 {
	return opt.unsafeValue
}

// String conforms to fmt Stringer interface.
func (opt *_Float32) String() string {
	if value, ok := opt.Unwrap(); ok {
		return fmt.Sprintf("Some(%v)", value)
	}
	return "None"
}

// MarshalJSON implements the json Marshaler interface.
func (opt _Float32) MarshalJSON() ([]byte, error) {
	if !opt.getHasValue() {
		return []byte("null"), nil
	}
	return json.Marshal(opt.getUnsafeValue())
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (opt *_Float32) UnmarshalJSON(data []byte) error {
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
