// Code generated by optionalize
// on: 2017-10-19 10:35:57.159581847 &#43;0000 UTC
// from: github.com/Gurpartap/safer-go/optional/template/optional.go.tmpl

package optional

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type Int32Slice interface {
	// Take takes the value out of the option, leaving a None in its place.
	Take()

	// From performs the conversion.
	From(value []int32)

	// Unwrap moves the value out of the optional, if it is Some(value).
	// This function returns multiple values, and if that's undesirable,
	// consider using Some and None functions.
	Unwrap() ([]int32, bool)

	// UnwrapOr returns the contained value or a default.
	UnwrapOr(def []int32) []int32

	// UnwrapOrElse returns the contained value or computes it from a closure.
	UnwrapOrElse(fn func() []int32) []int32

	// UnwrapOrPanic returns the contained value or panics.
	UnwrapOrPanic() []int32

	// And returns None if the optional is None, otherwise returns optb.
	And(optb Int32Slice) Int32Slice

	// Or returns the optional if it contains a value, otherwise returns optb.
	Or(optb Int32Slice) Int32Slice

	// OrElse returns the optional if it contains a value, otherwise calls fn
	// and returns the result.
	OrElse(fn func() Int32Slice) Int32Slice

	// IsSome returns true if the optional is a Some value.
	IsSome() bool

	// IsNone returns true if the optional is a None value.
	IsNone() bool

	// Some executes the given closure if there is a Some value.
	Some(fn func(value []int32)) someInt32SliceHandler

	// None executes the given closure if there is a None value.
	None(fn func())
}

type someInt32SliceHandler interface {
	None(fn func())
}

type _Int32Slice struct {
	hasValue    bool
	unsafeValue []int32
}

func NewInt32Slice() Int32Slice {
	opt := &_Int32Slice{}
	opt.Take()
	return opt
}

func Int32SliceFrom(value []int32, hasValue bool) Int32Slice {
	opt := &_Int32Slice{}
	if hasValue {
		opt.From(value)
	}
	return opt
}

// Take takes the value out of the optional, leaving a None in its place.
func (opt *_Int32Slice) Take() {
	var value []int32
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *_Int32Slice) From(value []int32) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *_Int32Slice) Unwrap() ([]int32, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *_Int32Slice) UnwrapOr(def []int32) []int32 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *_Int32Slice) UnwrapOrElse(fn func() []int32) []int32 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *_Int32Slice) UnwrapOrPanic() []int32 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap Int32Slice")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *_Int32Slice) And(optb Int32Slice) Int32Slice {
	if !opt.getHasValue() {
		return opt
	}
	return optb
}

// Or returns the optional if it contains a value, otherwise returns optb.
func (opt *_Int32Slice) Or(optb Int32Slice) Int32Slice {
	if opt.getHasValue() {
		return opt
	}
	return optb
}

// OrElse returns the optional if it contains a value, otherwise calls fn and
// returns the result.
func (opt *_Int32Slice) OrElse(fn func() Int32Slice) Int32Slice {
	if opt.getHasValue() {
		return opt
	}
	return fn()
}

// IsSome returns true if the optional is a Some value.
func (opt *_Int32Slice) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the optional is a None value.
func (opt *_Int32Slice) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *_Int32Slice) Some(fn func(value []int32)) someInt32SliceHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return _someInt32SliceHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *_Int32Slice) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type _someInt32SliceHandler struct {
	opt *_Int32Slice
}

func (some _someInt32SliceHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *_Int32Slice) getHasValue() bool {
	return opt.hasValue
}

func (opt *_Int32Slice) getUnsafeValue() []int32 {
	return opt.unsafeValue
}

// String conforms to fmt Stringer interface.
func (opt *_Int32Slice) String() string {
	if value, ok := opt.Unwrap(); ok {
		return fmt.Sprintf("Some(%v)", value)
	}
	return "None"
}

// MarshalJSON implements the json Marshaler interface.
func (opt _Int32Slice) MarshalJSON() ([]byte, error) {
	if !opt.getHasValue() {
		return []byte("null"), nil
	}
	return json.Marshal(opt.getUnsafeValue())
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (opt *_Int32Slice) UnmarshalJSON(data []byte) error {
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
