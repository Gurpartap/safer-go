// Code generated by optionalize
// on: 2017-10-19 12:53:35.806038532 &#43;0000 UTC
// from: github.com/Gurpartap/safer-go/optional/template/optional.go.tmpl

package optional

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type OptionalTimeSlice struct {
	hasValue    bool
	unsafeValue Time
}

func NewOptionalTimeSlice() OptionalTimeSlice {
	opt := &OptionalTimeSlice{}
	opt.Take()
	return *opt
}

func OptionalTimeSliceFrom(value Time, hasValue bool) OptionalTimeSlice {
	opt := &OptionalTimeSlice{}
	if hasValue {
		opt.From(value)
	}
	return *opt
}

// Take takes the value out of the optional, leaving a None in its place.
func (opt *OptionalTimeSlice) Take() {
	var value Time
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *OptionalTimeSlice) From(value Time) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *OptionalTimeSlice) Unwrap() (Time, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *OptionalTimeSlice) UnwrapOr(def Time) Time {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *OptionalTimeSlice) UnwrapOrElse(fn func() Time) Time {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *OptionalTimeSlice) UnwrapOrPanic() Time {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap OptionalTimeSlice")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *OptionalTimeSlice) And(optb OptionalTimeSlice) OptionalTimeSlice {
	if !opt.getHasValue() {
		return *opt
	}
	return optb
}

// Or returns the optional if it contains a value, otherwise returns optb.
func (opt *OptionalTimeSlice) Or(optb OptionalTimeSlice) OptionalTimeSlice {
	if opt.getHasValue() {
		return *opt
	}
	return optb
}

// OrElse returns the optional if it contains a value, otherwise calls fn and
// returns the result.
func (opt *OptionalTimeSlice) OrElse(fn func() OptionalTimeSlice) OptionalTimeSlice {
	if opt.getHasValue() {
		return *opt
	}
	return fn()
}

// IsSome returns true if the optional is a Some value.
func (opt *OptionalTimeSlice) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the optional is a None value.
func (opt *OptionalTimeSlice) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *OptionalTimeSlice) Some(fn func(value Time)) someOptionalTimeSliceHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return someOptionalTimeSliceHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *OptionalTimeSlice) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type someOptionalTimeSliceHandler struct {
	opt *OptionalTimeSlice
}

func (some someOptionalTimeSliceHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *OptionalTimeSlice) getHasValue() bool {
	return opt.hasValue
}

func (opt *OptionalTimeSlice) getUnsafeValue() Time {
	return opt.unsafeValue
}

// String conforms to fmt Stringer interface.
func (opt *OptionalTimeSlice) String() string {
	if value, ok := opt.Unwrap(); ok {
		return fmt.Sprintf("Some(%v)", value)
	}
	return "None"
}

// MarshalJSON implements the json Marshaler interface.
func (opt OptionalTimeSlice) MarshalJSON() ([]byte, error) {
	if !opt.getHasValue() {
		return []byte("null"), nil
	}
	return json.Marshal(opt.getUnsafeValue())
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (opt *OptionalTimeSlice) UnmarshalJSON(data []byte) error {
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
