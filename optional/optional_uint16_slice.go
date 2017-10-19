// Code generated by optionalize
// on: 2017-10-19 12:53:30.737123111 &#43;0000 UTC
// from: github.com/Gurpartap/safer-go/optional/template/optional.go.tmpl

package optional

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type OptionalUInt16Slice struct {
	hasValue    bool
	unsafeValue UInt16
}

func NewOptionalUInt16Slice() OptionalUInt16Slice {
	opt := &OptionalUInt16Slice{}
	opt.Take()
	return *opt
}

func OptionalUInt16SliceFrom(value UInt16, hasValue bool) OptionalUInt16Slice {
	opt := &OptionalUInt16Slice{}
	if hasValue {
		opt.From(value)
	}
	return *opt
}

// Take takes the value out of the optional, leaving a None in its place.
func (opt *OptionalUInt16Slice) Take() {
	var value UInt16
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *OptionalUInt16Slice) From(value UInt16) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *OptionalUInt16Slice) Unwrap() (UInt16, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *OptionalUInt16Slice) UnwrapOr(def UInt16) UInt16 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *OptionalUInt16Slice) UnwrapOrElse(fn func() UInt16) UInt16 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *OptionalUInt16Slice) UnwrapOrPanic() UInt16 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap OptionalUInt16Slice")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *OptionalUInt16Slice) And(optb OptionalUInt16Slice) OptionalUInt16Slice {
	if !opt.getHasValue() {
		return *opt
	}
	return optb
}

// Or returns the optional if it contains a value, otherwise returns optb.
func (opt *OptionalUInt16Slice) Or(optb OptionalUInt16Slice) OptionalUInt16Slice {
	if opt.getHasValue() {
		return *opt
	}
	return optb
}

// OrElse returns the optional if it contains a value, otherwise calls fn and
// returns the result.
func (opt *OptionalUInt16Slice) OrElse(fn func() OptionalUInt16Slice) OptionalUInt16Slice {
	if opt.getHasValue() {
		return *opt
	}
	return fn()
}

// IsSome returns true if the optional is a Some value.
func (opt *OptionalUInt16Slice) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the optional is a None value.
func (opt *OptionalUInt16Slice) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *OptionalUInt16Slice) Some(fn func(value UInt16)) someOptionalUInt16SliceHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return someOptionalUInt16SliceHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *OptionalUInt16Slice) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type someOptionalUInt16SliceHandler struct {
	opt *OptionalUInt16Slice
}

func (some someOptionalUInt16SliceHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *OptionalUInt16Slice) getHasValue() bool {
	return opt.hasValue
}

func (opt *OptionalUInt16Slice) getUnsafeValue() UInt16 {
	return opt.unsafeValue
}

// String conforms to fmt Stringer interface.
func (opt *OptionalUInt16Slice) String() string {
	if value, ok := opt.Unwrap(); ok {
		return fmt.Sprintf("Some(%v)", value)
	}
	return "None"
}

// MarshalJSON implements the json Marshaler interface.
func (opt OptionalUInt16Slice) MarshalJSON() ([]byte, error) {
	if !opt.getHasValue() {
		return []byte("null"), nil
	}
	return json.Marshal(opt.getUnsafeValue())
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (opt *OptionalUInt16Slice) UnmarshalJSON(data []byte) error {
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
