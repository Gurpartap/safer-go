// Code generated by optionalize
// on: 2017-10-19 10:35:51.025508118 &#43;0000 UTC
// from: github.com/Gurpartap/safer-go/optional/template/optional.go.tmpl

package optional

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type Int64 interface {
	// Take takes the value out of the option, leaving a None in its place.
	Take()

	// From performs the conversion.
	From(value int64)

	// Unwrap moves the value out of the optional, if it is Some(value).
	// This function returns multiple values, and if that's undesirable,
	// consider using Some and None functions.
	Unwrap() (int64, bool)

	// UnwrapOr returns the contained value or a default.
	UnwrapOr(def int64) int64

	// UnwrapOrElse returns the contained value or computes it from a closure.
	UnwrapOrElse(fn func() int64) int64

	// UnwrapOrPanic returns the contained value or panics.
	UnwrapOrPanic() int64

	// And returns None if the optional is None, otherwise returns optb.
	And(optb Int64) Int64

	// Or returns the optional if it contains a value, otherwise returns optb.
	Or(optb Int64) Int64

	// OrElse returns the optional if it contains a value, otherwise calls fn
	// and returns the result.
	OrElse(fn func() Int64) Int64

	// IsSome returns true if the optional is a Some value.
	IsSome() bool

	// IsNone returns true if the optional is a None value.
	IsNone() bool

	// Some executes the given closure if there is a Some value.
	Some(fn func(value int64)) someInt64Handler

	// None executes the given closure if there is a None value.
	None(fn func())
}

type someInt64Handler interface {
	None(fn func())
}

type _Int64 struct {
	hasValue    bool
	unsafeValue int64
}

func NewInt64() Int64 {
	opt := &_Int64{}
	opt.Take()
	return opt
}

func Int64From(value int64, hasValue bool) Int64 {
	opt := &_Int64{}
	if hasValue {
		opt.From(value)
	}
	return opt
}

// Take takes the value out of the optional, leaving a None in its place.
func (opt *_Int64) Take() {
	var value int64
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *_Int64) From(value int64) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *_Int64) Unwrap() (int64, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *_Int64) UnwrapOr(def int64) int64 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *_Int64) UnwrapOrElse(fn func() int64) int64 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *_Int64) UnwrapOrPanic() int64 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap Int64")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *_Int64) And(optb Int64) Int64 {
	if !opt.getHasValue() {
		return opt
	}
	return optb
}

// Or returns the optional if it contains a value, otherwise returns optb.
func (opt *_Int64) Or(optb Int64) Int64 {
	if opt.getHasValue() {
		return opt
	}
	return optb
}

// OrElse returns the optional if it contains a value, otherwise calls fn and
// returns the result.
func (opt *_Int64) OrElse(fn func() Int64) Int64 {
	if opt.getHasValue() {
		return opt
	}
	return fn()
}

// IsSome returns true if the optional is a Some value.
func (opt *_Int64) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the optional is a None value.
func (opt *_Int64) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *_Int64) Some(fn func(value int64)) someInt64Handler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return _someInt64Handler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *_Int64) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type _someInt64Handler struct {
	opt *_Int64
}

func (some _someInt64Handler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *_Int64) getHasValue() bool {
	return opt.hasValue
}

func (opt *_Int64) getUnsafeValue() int64 {
	return opt.unsafeValue
}

// String conforms to fmt Stringer interface.
func (opt *_Int64) String() string {
	if value, ok := opt.Unwrap(); ok {
		return fmt.Sprintf("Some(%v)", value)
	}
	return "None"
}

// MarshalJSON implements the json Marshaler interface.
func (opt _Int64) MarshalJSON() ([]byte, error) {
	if !opt.getHasValue() {
		return []byte("null"), nil
	}
	return json.Marshal(opt.getUnsafeValue())
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (opt *_Int64) UnmarshalJSON(data []byte) error {
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
