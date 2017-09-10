// Code generated by optionalize
// on: 2017-09-10 10:30:40.953417007 &#43;0000 UTC
// from: github.com/Gurpartap/safer-go/optional/template/optional.go.tmpl

package optional

import "fmt"

type Float32Slice interface {
	// Take takes the value out of the option, leaving a None in its place.
	Take()

	// From performs the conversion.
	From(value []float32)

	// Unwrap moves the value out of the optional, if it is Some(value).
	// This function returns multiple values, and if that's undesirable,
	// consider using Some and None functions.
	Unwrap() ([]float32, bool)

	// UnwrapOr returns the contained value or a default.
	UnwrapOr(def []float32) []float32

	// UnwrapOrElse returns the contained value or computes it from a closure.
	UnwrapOrElse(fn func() []float32) []float32

	// UnwrapOrPanic returns the contained value or panics.
	UnwrapOrPanic() []float32

	// And returns None if the optional is None, otherwise returns optb.
	And(optb Float32Slice) Float32Slice

	// Or returns the option if it contains a value, otherwise returns optb.
	Or(optb Float32Slice) Float32Slice

	// OrElse returns the option if it contains a value, otherwise calls fn and
	// returns the result.
	OrElse(fn func() Float32Slice) Float32Slice

	// IsSome returns true if the option is a Some value.
	IsSome() bool

	// IsNone returns true if the option is a None value.
	IsNone() bool

	// Some executes the given closure if there is a Some value.
	Some(fn func(value []float32)) someFloat32SliceHandler

	// None executes the given closure if there is a None value.
	None(fn func())
}

type someFloat32SliceHandler interface {
	None(fn func())
}

type _Float32Slice struct {
	hasValue    bool
	unsafeValue []float32
}

// String conforms to fmt.Stringer interface.
func (opt *_Float32Slice) String() string {
	if value, ok := opt.Unwrap(); ok {
		// some
		return fmt.Sprintf("Optional(%v)", value)
	}
	// none
	return "nil"
}

func NewFloat32Slice() Float32Slice {
	opt := &_Float32Slice{}
	opt.Take()
	return opt
}

func Float32SliceFrom(value []float32) Float32Slice {
	opt := &_Float32Slice{}
	opt.From(value)
	return opt
}

// Take takes the value out of the option, leaving a None in its place.
func (opt *_Float32Slice) Take() {
	var value []float32
	opt.unsafeValue = value
	opt.hasValue = false
}

// From performs the conversion.
func (opt *_Float32Slice) From(value []float32) {
	opt.unsafeValue = value
	opt.hasValue = true
}

// Unwrap moves the value out of the optional, if it is Some(value).
// This function returns multiple values, and if that's undesirable,
// consider using Some and None functions.
func (opt *_Float32Slice) Unwrap() ([]float32, bool) {
	return opt.getUnsafeValue(), opt.getHasValue()
}

// UnwrapOr returns the contained value or a default.
func (opt *_Float32Slice) UnwrapOr(def []float32) []float32 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return def
}

// UnwrapOrElse returns the contained value or computes it from a closure.
func (opt *_Float32Slice) UnwrapOrElse(fn func() []float32) []float32 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	return fn()
}

// UnwrapOrPanic returns the contained value or panics.
func (opt *_Float32Slice) UnwrapOrPanic() []float32 {
	if opt.getHasValue() {
		return opt.getUnsafeValue()
	}
	panic("unable to unwrap Float32Slice")
}

// And returns None if the optional is None, otherwise returns optb.
func (opt *_Float32Slice) And(optb Float32Slice) Float32Slice {
	if !opt.getHasValue() {
		return opt
	}
	return optb
}

// Or returns the option if it contains a value, otherwise returns optb.
func (opt *_Float32Slice) Or(optb Float32Slice) Float32Slice {
	if opt.getHasValue() {
		return opt
	}
	return optb
}

// OrElse returns the option if it contains a value, otherwise calls fn and
// returns the result.
func (opt *_Float32Slice) OrElse(fn func() Float32Slice) Float32Slice {
	if opt.getHasValue() {
		return opt
	}
	return fn()
}

// IsSome returns true if the option is a Some value.
func (opt *_Float32Slice) IsSome() bool {
	return opt.getHasValue()
}

// IsNone returns true if the option is a None value.
func (opt *_Float32Slice) IsNone() bool {
	return !opt.getHasValue()
}

// Some executes the given closure if there is a Some value.
func (opt *_Float32Slice) Some(fn func(value []float32)) someFloat32SliceHandler {
	if opt.getHasValue() {
		fn(opt.getUnsafeValue())
	}
	return _someFloat32SliceHandler{opt: opt}
}

// None executes the given closure if there is a None value.
func (opt *_Float32Slice) None(fn func()) {
	if !opt.getHasValue() {
		fn()
	}
}

type _someFloat32SliceHandler struct {
	opt *_Float32Slice
}

func (some _someFloat32SliceHandler) None(fn func()) {
	if !some.opt.getHasValue() {
		fn()
	}
}

func (opt *_Float32Slice) getHasValue() bool {
	return opt.hasValue
}

func (opt *_Float32Slice) getUnsafeValue() []float32 {
	return opt.unsafeValue
}
