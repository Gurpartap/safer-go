// Code generated by optionalize
// on: 2017-10-19 10:36:00.044009381 &#43;0000 UTC
// from: github.com/Gurpartap/safer-go/optional/template/optional.go.tmpl

package optional

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type BytesSlice interface {
	// Take takes the value out of the option, leaving a None in its place.
	Take()

	// From performs the conversion.
	From(value [][]byte)

	// Unwrap moves the value out of the optional, if it is Some(value).
	// This function returns multiple values, and if that's undesirable,
	// consider using Some and None functions.
	Unwrap() ([][]byte, bool)

	// UnwrapOr returns the contained value or a default.
	UnwrapOr(def [][]byte) [][]byte

	// UnwrapOrElse returns the contained value or computes it from a closure.
	UnwrapOrElse(fn func() [][]byte) [][]byte

	// UnwrapOrPanic returns the contained value or panics.
	UnwrapOrPanic() [][]byte

	// And returns None if the optional is None, otherwise returns optb.
	And(optb BytesSlice) BytesSlice

	// Or returns the optional if it contains a value, otherwise returns optb.
	Or(optb BytesSlice) BytesSlice

	// OrElse returns the optional if it contains a value, otherwise calls fn
	// and returns the result.
	OrElse(fn func() BytesSlice) BytesSlice

	// IsSome returns true if the optional is a Some value.
	IsSome() bool

	// IsNone returns true if the optional is a None value.
	IsNone() bool

	// Some executes the given closure if there is a Some value.
	Some(fn func(value [][]byte)) someBytesSliceHandler

	// None executes the given closure if there is a None value.
	None(fn func())
}

type someBytesSliceHandler interface {
	None(fn func())
}

type _BytesSlice struct {
	hasValue    bool
	unsafeValue [][]byte
}

func NewBytesSlice() BytesSlice {
	opt := &_BytesSlice{}
	opt.Take()
	return opt
}

func BytesSliceFrom(value [][]byte, hasValue bool) BytesSlice {
	opt := &_BytesSlice{}
	if hasValue {
		opt.From(value)
	}
	return opt
}

// Take takes the value out of the optional, leaving a None in its place.
func (opt *_BytesSlice) Take() {
	var value [][]byte
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *_BytesSlice) From(value [][]byte) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *_BytesSlice) Unwrap() ([][]byte, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *_BytesSlice) UnwrapOr(def [][]byte) [][]byte {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *_BytesSlice) UnwrapOrElse(fn func() [][]byte) [][]byte {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *_BytesSlice) UnwrapOrPanic() [][]byte {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap BytesSlice")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *_BytesSlice) And(optb BytesSlice) BytesSlice {
	if !opt.getHasValue() {
		return opt
	}
	return optb
}

// Or returns the optional if it contains a value, otherwise returns optb.
func (opt *_BytesSlice) Or(optb BytesSlice) BytesSlice {
	if opt.getHasValue() {
		return opt
	}
	return optb
}

// OrElse returns the optional if it contains a value, otherwise calls fn and
// returns the result.
func (opt *_BytesSlice) OrElse(fn func() BytesSlice) BytesSlice {
	if opt.getHasValue() {
		return opt
	}
	return fn()
}

// IsSome returns true if the optional is a Some value.
func (opt *_BytesSlice) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the optional is a None value.
func (opt *_BytesSlice) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *_BytesSlice) Some(fn func(value [][]byte)) someBytesSliceHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return _someBytesSliceHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *_BytesSlice) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type _someBytesSliceHandler struct {
	opt *_BytesSlice
}

func (some _someBytesSliceHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *_BytesSlice) getHasValue() bool {
	return opt.hasValue
}

func (opt *_BytesSlice) getUnsafeValue() [][]byte {
	return opt.unsafeValue
}

// String conforms to fmt Stringer interface.
func (opt *_BytesSlice) String() string {
	if value, ok := opt.Unwrap(); ok {
		return fmt.Sprintf("Some(%v)", value)
	}
	return "None"
}

// MarshalJSON implements the json Marshaler interface.
func (opt _BytesSlice) MarshalJSON() ([]byte, error) {
	if !opt.getHasValue() {
		return []byte("null"), nil
	}
	return json.Marshal(opt.getUnsafeValue())
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (opt *_BytesSlice) UnmarshalJSON(data []byte) error {
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
