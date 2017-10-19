// Code generated by optionalize
// on: 2017-10-19 08:41:09.564974223 &#43;0000 UTC
// from: github.com/Gurpartap/safer-go/optional/template/optional.go.tmpl

package optional

import "fmt"

type Rune interface {
	// Take takes the value out of the option, leaving a None in its place.
	Take()

	// From performs the conversion.
	From(value rune)

	// Unwrap moves the value out of the optional, if it is Some(value).
	// This function returns multiple values, and if that's undesirable,
	// consider using Some and None functions.
	Unwrap() (rune, bool)

	// UnwrapOr returns the contained value or a default.
	UnwrapOr(def rune) rune

	// UnwrapOrElse returns the contained value or computes it from a closure.
	UnwrapOrElse(fn func() rune) rune

	// UnwrapOrPanic returns the contained value or panics.
	UnwrapOrPanic() rune

	// And returns None if the optional is None, otherwise returns optb.
	And(optb Rune) Rune

	// Or returns the option if it contains a value, otherwise returns optb.
	Or(optb Rune) Rune

	// OrElse returns the option if it contains a value, otherwise calls fn and
	// returns the result.
	OrElse(fn func() Rune) Rune

	// IsSome returns true if the option is a Some value.
	IsSome() bool

	// IsNone returns true if the option is a None value.
	IsNone() bool

	// Some executes the given closure if there is a Some value.
	Some(fn func(value rune)) someRuneHandler

	// None executes the given closure if there is a None value.
	None(fn func())
}

type someRuneHandler interface {
	None(fn func())
}

type _Rune struct {
	hasValue    bool
	unsafeValue rune
}

// String conforms to fmt.Stringer interface.
func (opt *_Rune) String() string {
	if value, ok := opt.Unwrap(); ok {
		// some
		return fmt.Sprintf("Optional(%v)", value)
	}
	// none
	return "nil"
}

func NewRune() Rune {
	opt := &_Rune{}
	opt.Take()
	return opt
}

func RuneFrom(value rune, hasValue bool) Rune {
	opt := &_Rune{}
	if hasValue {
		opt.From(value)
	}
	return opt
}

// Take takes the value out of the option, leaving a None in its place.
func (opt *_Rune) Take() {
	var value rune
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *_Rune) From(value rune) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *_Rune) Unwrap() (rune, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *_Rune) UnwrapOr(def rune) rune {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *_Rune) UnwrapOrElse(fn func() rune) rune {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *_Rune) UnwrapOrPanic() rune {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap Rune")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *_Rune) And(optb Rune) Rune {
	if !opt.getHasValue() {
		return opt
	}
	return optb
}

// Or returns the option if it contains a value, otherwise returns optb.
func (opt *_Rune) Or(optb Rune) Rune {
	if opt.getHasValue() {
		return opt
	}
	return optb
}

// OrElse returns the option if it contains a value, otherwise calls fn and
// returns the result.
func (opt *_Rune) OrElse(fn func() Rune) Rune {
	if opt.getHasValue() {
		return opt
	}
	return fn()
}

// IsSome returns true if the option is a Some value.
func (opt *_Rune) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the option is a None value.
func (opt *_Rune) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *_Rune) Some(fn func(value rune)) someRuneHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return _someRuneHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *_Rune) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type _someRuneHandler struct {
	opt *_Rune
}

func (some _someRuneHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *_Rune) getHasValue() bool {
	return opt.hasValue
}

func (opt *_Rune) getUnsafeValue() rune {
	return opt.unsafeValue
}
