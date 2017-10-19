// Code generated by optionalize
// on: 2017-10-19 08:58:23.571688207 &#43;0000 UTC
// from: github.com/Gurpartap/safer-go/optional/template/optional.go.tmpl

package optional

import "fmt"

type Float64Slice interface {
	// Take takes the value out of the option, leaving a None in its place.
	Take()

	// From performs the conversion.
	From(value []float64)

	// Unwrap moves the value out of the optional, if it is Some(value).
	// This function returns multiple values, and if that's undesirable,
	// consider using Some and None functions.
	Unwrap() ([]float64, bool)

	// UnwrapOr returns the contained value or a default.
	UnwrapOr(def []float64) []float64

	// UnwrapOrElse returns the contained value or computes it from a closure.
	UnwrapOrElse(fn func() []float64) []float64

	// UnwrapOrPanic returns the contained value or panics.
	UnwrapOrPanic() []float64

	// And returns None if the optional is None, otherwise returns optb.
	And(optb Float64Slice) Float64Slice

	// Or returns the optional if it contains a value, otherwise returns optb.
	Or(optb Float64Slice) Float64Slice

	// OrElse returns the optional if it contains a value, otherwise calls fn
	// and returns the result.
	OrElse(fn func() Float64Slice) Float64Slice

	// IsSome returns true if the optional is a Some value.
	IsSome() bool

	// IsNone returns true if the optional is a None value.
	IsNone() bool

	// Some executes the given closure if there is a Some value.
	Some(fn func(value []float64)) someFloat64SliceHandler

	// None executes the given closure if there is a None value.
	None(fn func())
}

type someFloat64SliceHandler interface {
	None(fn func())
}

type _Float64Slice struct {
	hasValue    bool
	unsafeValue []float64
}

func NewFloat64Slice() Float64Slice {
	opt := &_Float64Slice{}
	opt.Take()
	return opt
}

func Float64SliceFrom(value []float64, hasValue bool) Float64Slice {
	opt := &_Float64Slice{}
	if hasValue {
		opt.From(value)
	}
	return opt
}

// Take takes the value out of the optional, leaving a None in its place.
func (opt *_Float64Slice) Take() {
	var value []float64
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *_Float64Slice) From(value []float64) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *_Float64Slice) Unwrap() ([]float64, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *_Float64Slice) UnwrapOr(def []float64) []float64 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *_Float64Slice) UnwrapOrElse(fn func() []float64) []float64 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *_Float64Slice) UnwrapOrPanic() []float64 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap Float64Slice")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *_Float64Slice) And(optb Float64Slice) Float64Slice {
	if !opt.getHasValue() {
		return opt
	}
	return optb
}

// Or returns the optional if it contains a value, otherwise returns optb.
func (opt *_Float64Slice) Or(optb Float64Slice) Float64Slice {
	if opt.getHasValue() {
		return opt
	}
	return optb
}

// OrElse returns the optional if it contains a value, otherwise calls fn and
// returns the result.
func (opt *_Float64Slice) OrElse(fn func() Float64Slice) Float64Slice {
	if opt.getHasValue() {
		return opt
	}
	return fn()
}

// IsSome returns true if the optional is a Some value.
func (opt *_Float64Slice) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the optional is a None value.
func (opt *_Float64Slice) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *_Float64Slice) Some(fn func(value []float64)) someFloat64SliceHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return _someFloat64SliceHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *_Float64Slice) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type _someFloat64SliceHandler struct {
	opt *_Float64Slice
}

func (some _someFloat64SliceHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *_Float64Slice) getHasValue() bool {
	return opt.hasValue
}

func (opt *_Float64Slice) getUnsafeValue() []float64 {
	return opt.unsafeValue
}

// String conforms to fmt.Stringer interface.
func (opt *_Float64Slice) String() string {
	if value, ok := opt.Unwrap(); ok {
		return fmt.Sprintf("Some(%v)", value)
	}
	return "None"
}
