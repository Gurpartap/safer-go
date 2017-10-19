// Code generated by optionalize
// on: 2017-10-19 08:41:13.96099068 &#43;0000 UTC
// from: github.com/Gurpartap/safer-go/optional/template/optional.go.tmpl

package optional

import "fmt"

type Complex64Slice interface {
	// Take takes the value out of the option, leaving a None in its place.
	Take()

	// From performs the conversion.
	From(value []complex64)

	// Unwrap moves the value out of the optional, if it is Some(value).
	// This function returns multiple values, and if that's undesirable,
	// consider using Some and None functions.
	Unwrap() ([]complex64, bool)

	// UnwrapOr returns the contained value or a default.
	UnwrapOr(def []complex64) []complex64

	// UnwrapOrElse returns the contained value or computes it from a closure.
	UnwrapOrElse(fn func() []complex64) []complex64

	// UnwrapOrPanic returns the contained value or panics.
	UnwrapOrPanic() []complex64

	// And returns None if the optional is None, otherwise returns optb.
	And(optb Complex64Slice) Complex64Slice

	// Or returns the option if it contains a value, otherwise returns optb.
	Or(optb Complex64Slice) Complex64Slice

	// OrElse returns the option if it contains a value, otherwise calls fn and
	// returns the result.
	OrElse(fn func() Complex64Slice) Complex64Slice

	// IsSome returns true if the option is a Some value.
	IsSome() bool

	// IsNone returns true if the option is a None value.
	IsNone() bool

	// Some executes the given closure if there is a Some value.
	Some(fn func(value []complex64)) someComplex64SliceHandler

	// None executes the given closure if there is a None value.
	None(fn func())
}

type someComplex64SliceHandler interface {
	None(fn func())
}

type _Complex64Slice struct {
	hasValue    bool
	unsafeValue []complex64
}

// String conforms to fmt.Stringer interface.
func (opt *_Complex64Slice) String() string {
	if value, ok := opt.Unwrap(); ok {
		// some
		return fmt.Sprintf("Optional(%v)", value)
	}
	// none
	return "nil"
}

func NewComplex64Slice() Complex64Slice {
	opt := &_Complex64Slice{}
	opt.Take()
	return opt
}

func Complex64SliceFrom(value []complex64, hasValue bool) Complex64Slice {
	opt := &_Complex64Slice{}
	if hasValue {
		opt.From(value)
	}
	return opt
}

// Take takes the value out of the option, leaving a None in its place.
func (opt *_Complex64Slice) Take() {
	var value []complex64
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *_Complex64Slice) From(value []complex64) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *_Complex64Slice) Unwrap() ([]complex64, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *_Complex64Slice) UnwrapOr(def []complex64) []complex64 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *_Complex64Slice) UnwrapOrElse(fn func() []complex64) []complex64 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *_Complex64Slice) UnwrapOrPanic() []complex64 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap Complex64Slice")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *_Complex64Slice) And(optb Complex64Slice) Complex64Slice {
	if !opt.getHasValue() {
		return opt
	}
	return optb
}

// Or returns the option if it contains a value, otherwise returns optb.
func (opt *_Complex64Slice) Or(optb Complex64Slice) Complex64Slice {
	if opt.getHasValue() {
		return opt
	}
	return optb
}

// OrElse returns the option if it contains a value, otherwise calls fn and
// returns the result.
func (opt *_Complex64Slice) OrElse(fn func() Complex64Slice) Complex64Slice {
	if opt.getHasValue() {
		return opt
	}
	return fn()
}

// IsSome returns true if the option is a Some value.
func (opt *_Complex64Slice) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the option is a None value.
func (opt *_Complex64Slice) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *_Complex64Slice) Some(fn func(value []complex64)) someComplex64SliceHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return _someComplex64SliceHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *_Complex64Slice) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type _someComplex64SliceHandler struct {
	opt *_Complex64Slice
}

func (some _someComplex64SliceHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *_Complex64Slice) getHasValue() bool {
	return opt.hasValue
}

func (opt *_Complex64Slice) getUnsafeValue() []complex64 {
	return opt.unsafeValue
}