// Code generated by optionalize
// on: 2017-10-19 08:41:07.984111732 &#43;0000 UTC
// from: github.com/Gurpartap/safer-go/optional/template/optional.go.tmpl

package optional

import "fmt"

type Complex128 interface {
	// Take takes the value out of the option, leaving a None in its place.
	Take()

	// From performs the conversion.
	From(value complex128)

	// Unwrap moves the value out of the optional, if it is Some(value).
	// This function returns multiple values, and if that's undesirable,
	// consider using Some and None functions.
	Unwrap() (complex128, bool)

	// UnwrapOr returns the contained value or a default.
	UnwrapOr(def complex128) complex128

	// UnwrapOrElse returns the contained value or computes it from a closure.
	UnwrapOrElse(fn func() complex128) complex128

	// UnwrapOrPanic returns the contained value or panics.
	UnwrapOrPanic() complex128

	// And returns None if the optional is None, otherwise returns optb.
	And(optb Complex128) Complex128

	// Or returns the option if it contains a value, otherwise returns optb.
	Or(optb Complex128) Complex128

	// OrElse returns the option if it contains a value, otherwise calls fn and
	// returns the result.
	OrElse(fn func() Complex128) Complex128

	// IsSome returns true if the option is a Some value.
	IsSome() bool

	// IsNone returns true if the option is a None value.
	IsNone() bool

	// Some executes the given closure if there is a Some value.
	Some(fn func(value complex128)) someComplex128Handler

	// None executes the given closure if there is a None value.
	None(fn func())
}

type someComplex128Handler interface {
	None(fn func())
}

type _Complex128 struct {
	hasValue    bool
	unsafeValue complex128
}

// String conforms to fmt.Stringer interface.
func (opt *_Complex128) String() string {
	if value, ok := opt.Unwrap(); ok {
		// some
		return fmt.Sprintf("Optional(%v)", value)
	}
	// none
	return "nil"
}

func NewComplex128() Complex128 {
	opt := &_Complex128{}
	opt.Take()
	return opt
}

func Complex128From(value complex128, hasValue bool) Complex128 {
	opt := &_Complex128{}
	if hasValue {
		opt.From(value)
	}
	return opt
}

// Take takes the value out of the option, leaving a None in its place.
func (opt *_Complex128) Take() {
	var value complex128
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *_Complex128) From(value complex128) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *_Complex128) Unwrap() (complex128, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *_Complex128) UnwrapOr(def complex128) complex128 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *_Complex128) UnwrapOrElse(fn func() complex128) complex128 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *_Complex128) UnwrapOrPanic() complex128 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap Complex128")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *_Complex128) And(optb Complex128) Complex128 {
	if !opt.getHasValue() {
		return opt
	}
	return optb
}

// Or returns the option if it contains a value, otherwise returns optb.
func (opt *_Complex128) Or(optb Complex128) Complex128 {
	if opt.getHasValue() {
		return opt
	}
	return optb
}

// OrElse returns the option if it contains a value, otherwise calls fn and
// returns the result.
func (opt *_Complex128) OrElse(fn func() Complex128) Complex128 {
	if opt.getHasValue() {
		return opt
	}
	return fn()
}

// IsSome returns true if the option is a Some value.
func (opt *_Complex128) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the option is a None value.
func (opt *_Complex128) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *_Complex128) Some(fn func(value complex128)) someComplex128Handler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return _someComplex128Handler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *_Complex128) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type _someComplex128Handler struct {
	opt *_Complex128
}

func (some _someComplex128Handler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *_Complex128) getHasValue() bool {
	return opt.hasValue
}

func (opt *_Complex128) getUnsafeValue() complex128 {
	return opt.unsafeValue
}