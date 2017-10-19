// Code generated by optionalize
// on: 2017-10-19 08:58:19.995167699 &#43;0000 UTC
// from: github.com/Gurpartap/safer-go/optional/template/optional.go.tmpl

package optional

import (
	"fmt"
	"time"
)

type Time interface {
	// Take takes the value out of the option, leaving a None in its place.
	Take()

	// From performs the conversion.
	From(value time.Time)

	// Unwrap moves the value out of the optional, if it is Some(value).
	// This function returns multiple values, and if that's undesirable,
	// consider using Some and None functions.
	Unwrap() (time.Time, bool)

	// UnwrapOr returns the contained value or a default.
	UnwrapOr(def time.Time) time.Time

	// UnwrapOrElse returns the contained value or computes it from a closure.
	UnwrapOrElse(fn func() time.Time) time.Time

	// UnwrapOrPanic returns the contained value or panics.
	UnwrapOrPanic() time.Time

	// And returns None if the optional is None, otherwise returns optb.
	And(optb Time) Time

	// Or returns the optional if it contains a value, otherwise returns optb.
	Or(optb Time) Time

	// OrElse returns the optional if it contains a value, otherwise calls fn
	// and returns the result.
	OrElse(fn func() Time) Time

	// IsSome returns true if the optional is a Some value.
	IsSome() bool

	// IsNone returns true if the optional is a None value.
	IsNone() bool

	// Some executes the given closure if there is a Some value.
	Some(fn func(value time.Time)) someTimeHandler

	// None executes the given closure if there is a None value.
	None(fn func())
}

type someTimeHandler interface {
	None(fn func())
}

type _Time struct {
	hasValue    bool
	unsafeValue time.Time
}

func NewTime() Time {
	opt := &_Time{}
	opt.Take()
	return opt
}

func TimeFrom(value time.Time, hasValue bool) Time {
	opt := &_Time{}
	if hasValue {
		opt.From(value)
	}
	return opt
}

// Take takes the value out of the optional, leaving a None in its place.
func (opt *_Time) Take() {
	var value time.Time
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *_Time) From(value time.Time) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *_Time) Unwrap() (time.Time, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *_Time) UnwrapOr(def time.Time) time.Time {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *_Time) UnwrapOrElse(fn func() time.Time) time.Time {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *_Time) UnwrapOrPanic() time.Time {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap Time")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *_Time) And(optb Time) Time {
	if !opt.getHasValue() {
		return opt
	}
	return optb
}

// Or returns the optional if it contains a value, otherwise returns optb.
func (opt *_Time) Or(optb Time) Time {
	if opt.getHasValue() {
		return opt
	}
	return optb
}

// OrElse returns the optional if it contains a value, otherwise calls fn and
// returns the result.
func (opt *_Time) OrElse(fn func() Time) Time {
	if opt.getHasValue() {
		return opt
	}
	return fn()
}

// IsSome returns true if the optional is a Some value.
func (opt *_Time) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the optional is a None value.
func (opt *_Time) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *_Time) Some(fn func(value time.Time)) someTimeHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return _someTimeHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *_Time) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type _someTimeHandler struct {
	opt *_Time
}

func (some _someTimeHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *_Time) getHasValue() bool {
	return opt.hasValue
}

func (opt *_Time) getUnsafeValue() time.Time {
	return opt.unsafeValue
}

// String conforms to fmt.Stringer interface.
func (opt *_Time) String() string {
	if value, ok := opt.Unwrap(); ok {
		return fmt.Sprintf("Some(%v)", value)
	}
	return "None"
}
