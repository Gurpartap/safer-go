// Code generated by optionalize
// on: 2017-10-19 08:41:09.870180282 &#43;0000 UTC
// from: github.com/Gurpartap/safer-go/optional/template/optional.go.tmpl

package optional

import "fmt"

type Error interface {
	// Take takes the value out of the option, leaving a None in its place.
	Take()

	// From performs the conversion.
	From(value error)

	// Unwrap moves the value out of the optional, if it is Some(value).
	// This function returns multiple values, and if that's undesirable,
	// consider using Some and None functions.
	Unwrap() (error, bool)

	// UnwrapOr returns the contained value or a default.
	UnwrapOr(def error) error

	// UnwrapOrElse returns the contained value or computes it from a closure.
	UnwrapOrElse(fn func() error) error

	// UnwrapOrPanic returns the contained value or panics.
	UnwrapOrPanic() error

	// And returns None if the optional is None, otherwise returns optb.
	And(optb Error) Error

	// Or returns the option if it contains a value, otherwise returns optb.
	Or(optb Error) Error

	// OrElse returns the option if it contains a value, otherwise calls fn and
	// returns the result.
	OrElse(fn func() Error) Error

	// IsSome returns true if the option is a Some value.
	IsSome() bool

	// IsNone returns true if the option is a None value.
	IsNone() bool

	// Some executes the given closure if there is a Some value.
	Some(fn func(value error)) someErrorHandler

	// None executes the given closure if there is a None value.
	None(fn func())
}

type someErrorHandler interface {
	None(fn func())
}

type _Error struct {
	hasValue    bool
	unsafeValue error
}

// String conforms to fmt.Stringer interface.
func (opt *_Error) String() string {
	if value, ok := opt.Unwrap(); ok {
		// some
		return fmt.Sprintf("Optional(%v)", value)
	}
	// none
	return "nil"
}

func NewError() Error {
	opt := &_Error{}
	opt.Take()
	return opt
}

func ErrorFrom(value error, hasValue bool) Error {
	opt := &_Error{}
	if hasValue {
		opt.From(value)
	}
	return opt
}

// Take takes the value out of the option, leaving a None in its place.
func (opt *_Error) Take() {
	var value error
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *_Error) From(value error) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *_Error) Unwrap() (error, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *_Error) UnwrapOr(def error) error {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *_Error) UnwrapOrElse(fn func() error) error {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *_Error) UnwrapOrPanic() error {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap Error")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *_Error) And(optb Error) Error {
	if !opt.getHasValue() {
		return opt
	}
	return optb
}

// Or returns the option if it contains a value, otherwise returns optb.
func (opt *_Error) Or(optb Error) Error {
	if opt.getHasValue() {
		return opt
	}
	return optb
}

// OrElse returns the option if it contains a value, otherwise calls fn and
// returns the result.
func (opt *_Error) OrElse(fn func() Error) Error {
	if opt.getHasValue() {
		return opt
	}
	return fn()
}

// IsSome returns true if the option is a Some value.
func (opt *_Error) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the option is a None value.
func (opt *_Error) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *_Error) Some(fn func(value error)) someErrorHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return _someErrorHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *_Error) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type _someErrorHandler struct {
	opt *_Error
}

func (some _someErrorHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *_Error) getHasValue() bool {
	return opt.hasValue
}

func (opt *_Error) getUnsafeValue() error {
	return opt.unsafeValue
}