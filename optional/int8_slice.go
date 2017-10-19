// Code generated by optionalize
// on: 2017-10-19 12:53:25.151038124 &#43;0000 UTC
// from: github.com/Gurpartap/safer-go/optional/template/optional.go.tmpl

package optional

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type Int8Slice struct {
	hasValue    bool
	unsafeValue []int8
}

func NewInt8Slice() Int8Slice {
	opt := &Int8Slice{}
	opt.Take()
	return *opt
}

func Int8SliceFrom(value []int8, hasValue bool) Int8Slice {
	opt := &Int8Slice{}
	if hasValue {
		opt.From(value)
	}
	return *opt
}

// Take takes the value out of the optional, leaving a None in its place.
func (opt *Int8Slice) Take() {
	var value []int8
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *Int8Slice) From(value []int8) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *Int8Slice) Unwrap() ([]int8, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *Int8Slice) UnwrapOr(def []int8) []int8 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *Int8Slice) UnwrapOrElse(fn func() []int8) []int8 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *Int8Slice) UnwrapOrPanic() []int8 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap Int8Slice")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *Int8Slice) And(optb Int8Slice) Int8Slice {
	if !opt.getHasValue() {
		return *opt
	}
	return optb
}

// Or returns the optional if it contains a value, otherwise returns optb.
func (opt *Int8Slice) Or(optb Int8Slice) Int8Slice {
	if opt.getHasValue() {
		return *opt
	}
	return optb
}

// OrElse returns the optional if it contains a value, otherwise calls fn and
// returns the result.
func (opt *Int8Slice) OrElse(fn func() Int8Slice) Int8Slice {
	if opt.getHasValue() {
		return *opt
	}
	return fn()
}

// IsSome returns true if the optional is a Some value.
func (opt *Int8Slice) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the optional is a None value.
func (opt *Int8Slice) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *Int8Slice) Some(fn func(value []int8)) someInt8SliceHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return someInt8SliceHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *Int8Slice) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type someInt8SliceHandler struct {
	opt *Int8Slice
}

func (some someInt8SliceHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *Int8Slice) getHasValue() bool {
	return opt.hasValue
}

func (opt *Int8Slice) getUnsafeValue() []int8 {
	return opt.unsafeValue
}

// String conforms to fmt Stringer interface.
func (opt *Int8Slice) String() string {
	if value, ok := opt.Unwrap(); ok {
		return fmt.Sprintf("Some(%v)", value)
	}
	return "None"
}

// MarshalJSON implements the json Marshaler interface.
func (opt Int8Slice) MarshalJSON() ([]byte, error) {
	if !opt.getHasValue() {
		return []byte("null"), nil
	}
	return json.Marshal(opt.getUnsafeValue())
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (opt *Int8Slice) UnmarshalJSON(data []byte) error {
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
