// Code generated by optionalize
// on: 2017-10-19 08:58:13.917498445 &#43;0000 UTC
// from: github.com/Gurpartap/safer-go/optional/template/optional.go.tmpl

package optional

import "fmt"

type Bool interface {
	// Take takes the value out of the option, leaving a None in its place.
	Take()

	// From performs the conversion.
	From(value bool)

	// Unwrap moves the value out of the optional, if it is Some(value).
	// This function returns multiple values, and if that's undesirable,
	// consider using Some and None functions.
	Unwrap() (bool, bool)

	// UnwrapOr returns the contained value or a default.
	UnwrapOr(def bool) bool

	// UnwrapOrElse returns the contained value or computes it from a closure.
	UnwrapOrElse(fn func() bool) bool

	// UnwrapOrPanic returns the contained value or panics.
	UnwrapOrPanic() bool

	// And returns None if the optional is None, otherwise returns optb.
	And(optb Bool) Bool

	// Or returns the optional if it contains a value, otherwise returns optb.
	Or(optb Bool) Bool

	// OrElse returns the optional if it contains a value, otherwise calls fn
	// and returns the result.
	OrElse(fn func() Bool) Bool

	// IsSome returns true if the optional is a Some value.
	IsSome() bool

	// IsNone returns true if the optional is a None value.
	IsNone() bool

	// Some executes the given closure if there is a Some value.
	Some(fn func(value bool)) someBoolHandler

	// None executes the given closure if there is a None value.
	None(fn func())
}

type someBoolHandler interface {
	None(fn func())
}

type _Bool struct {
	hasValue    bool
	unsafeValue bool
}

func NewBool() Bool {
	opt := &_Bool{}
	opt.Take()
	return opt
}

func BoolFrom(value bool, hasValue bool) Bool {
	opt := &_Bool{}
	if hasValue {
		opt.From(value)
	}
	return opt
}

// Take takes the value out of the optional, leaving a None in its place.
func (opt *_Bool) Take() {
	var value bool
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *_Bool) From(value bool) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *_Bool) Unwrap() (bool, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *_Bool) UnwrapOr(def bool) bool {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *_Bool) UnwrapOrElse(fn func() bool) bool {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *_Bool) UnwrapOrPanic() bool {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap Bool")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *_Bool) And(optb Bool) Bool {
	if !opt.getHasValue() {
		return opt
	}
	return optb
}

// Or returns the optional if it contains a value, otherwise returns optb.
func (opt *_Bool) Or(optb Bool) Bool {
	if opt.getHasValue() {
		return opt
	}
	return optb
}

// OrElse returns the optional if it contains a value, otherwise calls fn and
// returns the result.
func (opt *_Bool) OrElse(fn func() Bool) Bool {
	if opt.getHasValue() {
		return opt
	}
	return fn()
}

// IsSome returns true if the optional is a Some value.
func (opt *_Bool) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the optional is a None value.
func (opt *_Bool) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *_Bool) Some(fn func(value bool)) someBoolHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return _someBoolHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *_Bool) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type _someBoolHandler struct {
	opt *_Bool
}

func (some _someBoolHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *_Bool) getHasValue() bool {
	return opt.hasValue
}

func (opt *_Bool) getUnsafeValue() bool {
	return opt.unsafeValue
}

// String conforms to fmt.Stringer interface.
func (opt *_Bool) String() string {
	if value, ok := opt.Unwrap(); ok {
		return fmt.Sprintf("Some(%v)", value)
	}
	return "None"
}
